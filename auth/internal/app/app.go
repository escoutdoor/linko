package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"buf.build/go/protovalidate"
	authv1_implementation "github.com/escoutdoor/linko/auth/internal/api/auth/v1"
	userv1_implementation "github.com/escoutdoor/linko/auth/internal/api/user/v1"
	"github.com/escoutdoor/linko/auth/internal/config"
	"github.com/escoutdoor/linko/auth/internal/interceptor"
	"github.com/escoutdoor/linko/auth/internal/metrics"
	"github.com/escoutdoor/linko/common/pkg/closer"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
	common_interceptor "github.com/escoutdoor/linko/common/pkg/interceptor"
	"github.com/escoutdoor/linko/common/pkg/logger"
	authv1 "github.com/escoutdoor/linko/common/pkg/proto/auth/v1"
	userv1 "github.com/escoutdoor/linko/common/pkg/proto/user/v1"
	"github.com/escoutdoor/linko/common/pkg/tracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	swaggerFile  = "swagger/api.swagger.json"
	swaggerUIDir = "swagger"
)

type App struct {
	di *di

	grpcServer       *grpc.Server
	httpServer       *http.Server
	prometheusServer *http.Server
}

func New(ctx context.Context) (*App, error) {
	app := &App{di: newDiContainer()}
	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return nil, errwrap.Wrap("set migrations dialect", err)
	}

	db := stdlib.OpenDBFromPool(app.di.DBClient(ctx).DB().Pool())
	if err := goose.UpContext(ctx, db, config.AppConfig().Postgres.MigrationsDir()); err != nil {
		return nil, errwrap.Wrap("migrate up", err)
	}

	if err := db.Close(); err != nil {
		return nil, errwrap.Wrap("close db after migrate up", err)
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		logger.Info(ctx, "grpc server is running")
		if err := a.runGrpcServer(); err != nil {
			logger.Fatal(ctx, "run grpc server", err)
		}
	}()

	go func() {
		logger.Info(ctx, "http server is running")
		if err := a.runHttpServer(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(ctx, "run http server", err)
		}
	}()

	go func() {
		logger.Info(ctx, "prometheus server is running")
		if err := a.runPrometheusServer(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(ctx, "run prometheus server", err)
		}
	}()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initTracing,
		a.initMetrics,
		a.initGrpcServer,
		a.initHttpServer,
		a.initPrometheusServer,
	}

	for _, d := range deps {
		if err := d(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initTracing(ctx context.Context) error {
	if err := tracing.Init(ctx, config.AppConfig().GrpcServer.Address(), config.AppConfig().Service.Name()); err != nil {
		return errwrap.Wrap("init tracing", err)
	}

	return nil
}

func (a *App) initMetrics(_ context.Context) error {
	metrics.Init(config.AppConfig().Prometheus.Namespace(), config.AppConfig().Service.Name())
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	validator, err := protovalidate.New()
	if err != nil {
		return errwrap.Wrap("new validator", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			interceptor.ErrorsUnaryServerInterceptor(),
			interceptor.MetricsUnaryServerInterceptor(),
			common_interceptor.LoggingUnaryServerInterceptor(),
			common_interceptor.ValidationUnaryServerInterceptor(validator),
			common_interceptor.RecoverUnaryServerInterceptor(),
		),
	)

	authv1Impl := authv1_implementation.NewImplementation(a.di.AuthService(ctx))
	userv1Impl := userv1_implementation.NewImplementation(a.di.UserService(ctx))

	authv1.RegisterAuthServiceServer(grpcServer, authv1Impl)
	userv1.RegisterUserServiceServer(grpcServer, userv1Impl)

	reflection.Register(grpcServer)

	a.grpcServer = grpcServer

	closer.Add(func(ctx context.Context) error {
		a.grpcServer.GracefulStop()
		return nil
	})

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	gwMux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, gwMux, config.AppConfig().GrpcServer.Address(), opts); err != nil {
		return errwrap.Wrap("register user service handler from endpoint", err)
	}

	if err := authv1.RegisterAuthServiceHandlerFromEndpoint(ctx, gwMux, config.AppConfig().GrpcServer.Address(), opts); err != nil {
		return errwrap.Wrap("register auth service handler from endpoint", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)

	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerFile)
	})
	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir(swaggerUIDir))))

	s := &http.Server{
		Addr:              config.AppConfig().HttpServer.Address(),
		Handler:           mux,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}

	a.httpServer = s

	closer.Add(func(ctx context.Context) error {
		return a.httpServer.Shutdown(ctx)
	})

	return nil
}

func (a *App) initPrometheusServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              config.AppConfig().Prometheus.Address(),
		Handler:           mux,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}

	closer.Add(func(ctx context.Context) error {
		return a.prometheusServer.Shutdown(ctx)
	})
	return nil
}

func (a *App) runGrpcServer() error {
	ln, err := net.Listen("tcp", config.AppConfig().GrpcServer.Address())
	if err != nil {
		return errwrap.Wrap("net listen", err)
	}

	if err := a.grpcServer.Serve(ln); err != nil {
		return errwrap.Wrap("grpc server serve", err)
	}

	return nil
}

func (a *App) runHttpServer() error {
	if err := a.httpServer.ListenAndServe(); err != nil {
		return errwrap.Wrap("gateway server listen and serve", err)
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	if err := a.prometheusServer.ListenAndServe(); err != nil {
		return errwrap.Wrap("prometheus server listen and serve", err)
	}

	return nil
}
