package logger

import (
	"go.uber.org/zap"

	"github.com/ramezanius/octoql/pkg/common/env"
)

var L *zap.Logger

func init() {
	Prepare(env.Environment)
}

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
