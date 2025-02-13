package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/LLIEPJIOK/orderserver/pkg/database/postgres"
	"github.com/LLIEPJIOK/orderserver/pkg/database/redis"
)

type Config struct {
	postgres.PostgresConfig
	redis.RedisConfig

	GRPCServerPort int `env:"GRPC_SERVER_PORT" env-default:"50051"`
	RESTServerPort int `env:"REST_SERVER_PORT" env-default:"8080"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := cleanenv.ReadConfig("./.env", &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return &cfg, nil
}
