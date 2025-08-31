package env

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

const (
	prodStage = "prod"
	devStage  = "dev"
)

type serviceConfig struct {
	ServiceName                    string        `env:"SERVICE_NAME,required"`
	ServiceStage                   string        `env:"SERVICE_STAGE,required"`
	ServiceGracefulShutdownTimeout time.Duration `env:"SERVICE_GRACEFUL_SHUTDOWN_TIMEOUT,required"`
}

func NewServiceConfig() (*serviceConfig, error) {
	config := new(serviceConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	if config.ServiceStage != prodStage && config.ServiceStage != devStage {
		return nil, fmt.Errorf(`unknown stage option: %s (only %s or %s)`, config.ServiceStage, prodStage, devStage)
	}

	return config, nil
}

func (c *serviceConfig) Name() string {
	return c.ServiceName
}

func (c *serviceConfig) Stage() string {
	return c.ServiceStage
}

func (c *serviceConfig) IsProd() bool {
	return c.ServiceStage == prodStage
}

func (c *serviceConfig) GracefulShutdownTimeout() time.Duration {
	return c.ServiceGracefulShutdownTimeout
}
