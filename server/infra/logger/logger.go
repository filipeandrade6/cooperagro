package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(service string) (*zap.SugaredLogger, error) {
	// TODO descomentar quando estiver em produção
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	return nil, fmt.Errorf("localizando diretorio do usuario: %w", err)
	// }

	// f, err := os.OpenFile(filepath.Join(homeDir, "ginova.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return nil, fmt.Errorf("criando arquivo de log: %w", err)
	// }

	config := zap.NewProductionConfig()
	// config.OutputPaths = []string{f.Name(), "stdout"}
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

	return log.Sugar(), nil
}
