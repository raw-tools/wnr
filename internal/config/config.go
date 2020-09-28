package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Store *viper.Viper

	Config struct {
	}
)

func New() *Config {
	return &Config{}
}

func FromArgs(args []string) (*Config, error) {
	return nil, errors.New("Not implemented")
}
