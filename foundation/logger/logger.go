// Package logger provides a convenience function to constructing a logger
// for use.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New constructs a Structured Logger that writes to stdout and
// provides human-readable timestamps.
func New(service string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log, nil
}
