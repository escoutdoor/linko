package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type grpcServerConfig struct {
	Host string `env:"GRPC_SERVER_HOST,required"`
	Port string `env:"GRPC_SERVER_PORT,required"`
}

func NewGrpcServerConfig() (*grpcServerConfig, error) {
	config := new(grpcServerConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *grpcServerConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}
