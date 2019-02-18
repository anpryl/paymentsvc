package config

import (
	"log"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	PostgreSQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Http struct {
		Host string `envconfig:"default=0.0.0.0"`
		Port int    `envconfig:"default=80"`
	}
}

func FromEnv() Config {
	var c Config
	if err := envconfig.Init(&c); err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	return c
}
