package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type jaegerConfig struct {
	Host string `env:"JAEGER_HOST,required"`
	Port string `env:"JAEGER_PORT,required"`
}

func NewJaegerServerConfig() (*jaegerConfig, error) {
	config := new(jaegerConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *jaegerConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}
