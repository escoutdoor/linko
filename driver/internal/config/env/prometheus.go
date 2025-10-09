package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type prometheusConfig struct {
	PrometheusNamespace string `env:"PROMETHEUS_NAMESPACE,required"`
	Host                string `env:"PROMETHEUS_SERVER_HOST,required"`
	Port                string `env:"PROMETHEUS_SERVER_PORT,required"`
}

func NewPrometheusConfig() (*prometheusConfig, error) {
	config := new(prometheusConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *prometheusConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func (c *prometheusConfig) Namespace() string {
	return c.PrometheusNamespace
}
