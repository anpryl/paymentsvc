package config

import (
	"log"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	PostgreSQL struct {
		Host     string
		Port     int
		Database struct {
			User     string
			Password string
			Name     string
		}
	}
	Http struct {
		Host string `envconfig:"default=127.0.0.1"`
		Port int    `envconfig:"default=9090"`
	}
}

func FromEnv() Config {
	var c Config
	if err := envconfig.Init(&c); err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	return c
}
