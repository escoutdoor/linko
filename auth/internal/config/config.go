package config

import (
	"time"

	"github.com/escoutdoor/linko/auth/internal/config/env"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/joho/godotenv"
)

type config struct {
	App        AppConfig
	GrpcServer GrpcServerConfig
	HttpServer HttpServerConfig
	Prometheus PrometheusConfig
	Postgres   PostgresConfig
	Jaeger     JaegerConfig
	JwtToken   JwtTokenConfig
}

var cfg *config

func Config() *config {
	return cfg
}

type AppConfig interface {
	Name() string
	Stage() string
	IsProd() bool
	GracefulShutdownTimeout() time.Duration
}

type GrpcServerConfig interface {
	Address() string
}

type HttpServerConfig interface {
	Address() string
}

type PrometheusConfig interface {
	Address() string
	Namespace() string
}

type PostgresConfig interface {
	Dsn() string
	MigrationsDir() string
}

type JaegerConfig interface {
	Address() string
}

type JwtTokenConfig interface {
	AccessTokenSecretKey() string
	AccessTokenTTL() time.Duration
	RefreshTokenSecretKey() string
	RefreshTokenTTL() time.Duration
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

	jwtTokenConfig, err := env.NewJwtTokenConfig()
	if err != nil {
		return errwrap.Wrap("jwt token config", err)
	}

	cfg = &config{
		App:        appConfig,
		GrpcServer: grpcServerConfig,
		HttpServer: httpServerConfig,
		Prometheus: prometheusConfig,
		Postgres:   postgresConfig,
		Jaeger:     jaegerConfig,
		JwtToken:   jwtTokenConfig,
	}

	return nil
}
