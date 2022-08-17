package config

import (
	// 3rd party
	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel      string
	InfraHttpPort int `env:"INFRA_HTTP_PORT" envDefault:"4000"`
	GRPCPort      int `env:"GRPC_PORT" envDefault:"50000"`

	DB struct {
		Host     string `env:"PG_HOST" envDefault:"localhost:5432"`
		DBName   string `env:"PG_DBNAME" envDefault:"users_db"`
		Username string `env:"PG_USERNAME" envDefault:"db_user"`
		Password string `env:"PG_PASSWORD" envDefault:"pwd123"`
	}

	Kafka struct {
		ProducerBrokers string `env:"PRODUCER_BROKERS" envDefault:"localhost:9092"`
	}
}

func New() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
