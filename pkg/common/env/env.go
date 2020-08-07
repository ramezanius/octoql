package env

import (
	"os"
)

type Env string

var (
	Production  Env = "PROD"
	Development Env = "DEV"
)

var Environment Env

func init() {
	Environment = Env(os.Getenv("ENV"))

	// Set default environment (default: development)
	if Environment == "" {
		Environment = Development
	}
}
