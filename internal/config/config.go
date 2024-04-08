package config

import (
	"fmt"
	"os"

	"simple-queue-writer/internal"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server
}

type Server struct {
	Port string `yaml:"server_port"`
}

func ConfigLoad() (*Config, error) {
	f, err := os.Open("application.yml")
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, internal.ErrInvalidEnv)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, internal.ErrConfig)
	}

	return &cfg, nil
}
