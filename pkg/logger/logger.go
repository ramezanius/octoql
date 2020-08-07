package logger

import (
	"go.uber.org/zap"

	"github.com/ramezanius/octoql/pkg/common/env"
)

// L common logger
var L *zap.Logger

func init() {
	Prepare(env.Environment)
}

// Prepare initializes logger with environment and
// replace global logger (zap)
func Prepare(environment env.Env) {
	switch environment {
	case env.Production:
		L, _ = zap.NewProduction()
	case env.Development:
		L, _ = zap.NewDevelopment()
	default:
		panic("unknown environment")
	}

	zap.ReplaceGlobals(L)
}
