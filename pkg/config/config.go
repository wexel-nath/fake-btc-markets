package config

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"
)

var (
	instance Config
)

type Config struct{
	ImageTag string `env:"IMAGE_TAG"`
	LogLevel int    `env:"LOG_LEVEL,default=20"`
	Port     int    `env:"PORT,default=80"`
}

func Configure() {
	err := envconfig.Process(context.Background(), &instance)
	if err != nil {
		log.Fatal(err)
	}
}

func Get() Config {
	return instance
}

func (c Config) GetListenAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", c.Port)
}
