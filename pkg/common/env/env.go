package env

import (
	"os"
)

// Env environment type
type Env string

var (
	// Production environment
	Production Env = "PROD"
	// Development environment
	Development Env = "DEV"
)

// Environment common environment
var Environment Env

func init() {
	Environment = Env(os.Getenv("ENV"))

	// Set default environment (default: development)
	if Environment == "" {
		Environment = Development
	}
}
