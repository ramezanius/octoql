package config

import (
	"fmt"

	"go.uber.org/config"
)

type cfg struct {
	Debug bool
	Cache struct {
		Host string
	}
	Server struct {
		Host string
	}
	Database struct {
		Host     string
		User     string
		Pass     string
		Dialect  string
		Database string
	}
}

var C *cfg

const DefaultPath = ".octoql/config.yaml"

func init() {
	Load(config.Root)
}

func Load(path string) {
	// Set default configuration path
	if path == config.Root {
		path = DefaultPath
	}

	provider, err := config.NewYAML(config.File(path))
	if err != nil {
		panic(fmt.Errorf("error in initializing provider: %+v", err))
	}

	if err := provider.Get(config.Root).Populate(&C); err != nil {
		panic(fmt.Errorf("error in reading configuration %+v", err))
	}
}
