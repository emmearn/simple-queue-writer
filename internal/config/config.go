package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	AWS    AWS
}

type Server struct {
	Port string `yaml:"port"`
}

type AWS struct {
	Region string `yaml:"region"`
	SQS    SQS    `yaml:"sqs"`
}

type SQS struct {
	QueueName string `yaml:"queue_name"`
	QueueURL  string `yaml:"queue_url"`
}

var (
	ErrInvalidEnv = errors.New("invalid_env")
	ErrConfig     = errors.New("config_error")
)

func ConfigLoad() (*Config, error) {
	var cfg Config

	yamlFile, err := os.ReadFile("application.yml")
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, ErrInvalidEnv)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, ErrConfig)
	}

	return &cfg, nil
}
