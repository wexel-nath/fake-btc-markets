package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

var (
	instance Config
)

type Config struct{
	LogLevel int    `env:"LOG_LEVEL,default=20"`
	Port     int    `env:"PORT,default=80"`
	Version  string `env:"VERSION"`
}

func Configure() error {
	return envconfig.Process(context.Background(), &instance)
}

func Get() Config {
	return instance
}

func (c Config) GetListenAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", c.Port)
}
