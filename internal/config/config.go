package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server
	AWS    AWS
}

type Server struct {
	Port string `yaml:"server_port"`
}

type AWS struct {
	Stage  string `yaml:"STAGE,required"`
	Region string `yaml:"REGION,required"`
	SQS    SQS
}

type SQS struct {
	QueueName string `yaml:"AWS_SQS_QUEUE_NAME,required"`
	QueueURL  string `yaml:"AWS_SQS_QUEUE_URL,required"`
}

var (
	ErrInvalidEnv = errors.New("invalid_env")
	ErrConfig     = errors.New("config_error")
)

func ConfigLoad() (*Config, error) {
	f, err := os.Open("application.yml")
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, ErrInvalidEnv)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, ErrConfig)
	}

	return &cfg, nil
}
