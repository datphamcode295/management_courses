package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() func() {
	var logger *zap.Logger
	var encoderCfg zapcore.EncoderConfig
	var consoleEncoder zapcore.Encoder

	switch os.Getenv("RUN_MODE") {
	case "local":
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.FunctionKey = "func"
		consoleEncoder = zapcore.NewConsoleEncoder(encoderCfg)

	default:
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.FunctionKey = "func"
		consoleEncoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	atom := zap.NewAtomicLevel()
	logLevel := os.Getenv("LOG_LEVEL")
	atom.UnmarshalText([]byte(logLevel))

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, atom),
	)

	logger = zap.New(core, zap.AddCaller())
	return zap.ReplaceGlobals(logger)
}
