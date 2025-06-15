package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env string `yaml:"env"`
	Add string `yaml:"add"`
}

func ConfigInit() *Config {
	configPath := "config.yaml"
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Printf("ConfigInit err: %s", err.Error())
		cfg.Env = "local"
		cfg.Add = "localhost:8080"
	}
	return &cfg
}
