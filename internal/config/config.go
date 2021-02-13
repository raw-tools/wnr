package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Store *viper.Viper

	Watcher struct {
		Shortcut string
		Details *WatcherDetails
	}

	WatcherDetails struct {
		Files []string
		Exclude []string
		Extends string
	}

	Task struct {
		Shortcut string
		Details *TaskDetails
	}

	TaskDetails struct {
		Extends string
		Cmd string
		Stdout bool
	}	

	Profile struct {
		Extends []string
		Watchers []string
		Tasks []string
	}

	Config struct {
		Version string
		Verbose bool
		Tasks map[string]*Task
		Watchers map[string]*Watcher
		Profiles map[string]*Profile
	}
)

func New() *Config {
	return &Config{}
}

func FromArgs(args []string) (*Config, error) {
	return nil, errors.New("Not implemented")
}
