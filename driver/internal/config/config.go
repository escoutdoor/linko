package config

import (
	"time"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/config/env"
	"github.com/joho/godotenv"
)

type config struct {
	App        App
	GrpcServer GrpcServer
	HttpServer HttpServer
	Prometheus Prometheus
	Postgres   Postgres
	Jaeger     Jaeger
}

var cfg *config

func Config() *config {
	return cfg
}

type App interface {
	Name() string
	Stage() string
	IsProd() bool
	GracefulShutdownTimeout() time.Duration
}

type GrpcServer interface {
	Address() string
}

type HttpServer interface {
	Address() string
}

type Prometheus interface {
	Address() string
	Namespace() string
}

type Postgres interface {
	Dsn() string
	MigrationsDir() string
}

type Jaeger interface {
	Address() string
}

func Load(paths ...string) error {
	if len(paths) > 0 {
		if err := godotenv.Load(paths...); err != nil {
			return errwrap.Wrap("load config", err)
		}
	}

	appConfig, err := env.NewAppConfig()
	if err != nil {
		return errwrap.Wrap("app config", err)
	}

	grpcServerConfig, err := env.NewGrpcServerConfig()
	if err != nil {
		return errwrap.Wrap("grpc server config", err)
	}

	httpServerConfig, err := env.NewHttpServerConfig()
	if err != nil {
		return errwrap.Wrap("http server config", err)
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		return errwrap.Wrap("postgres config", err)
	}

	prometheusConfig, err := env.NewPrometheusConfig()
	if err != nil {
		return errwrap.Wrap("prometheus config", err)
	}

	jaegerConfig, err := env.NewJaegerServerConfig()
	if err != nil {
		return errwrap.Wrap("jaeger config", err)
	}

	appConfig = &config{
		App:        appConfig,
		GrpcServer: grpcServerConfig,
		HttpServer: httpServerConfig,
		Prometheus: prometheusConfig,
		Postgres:   postgresConfig,
		Jaeger:     jaegerConfig,
	}

	return nil
}
