package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel string) (*zap.Logger, func() error, error) {
	// перевод уровня логгирования из строкового формата в структуру
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	// создание директории для хранения логов
	if err := os.MkdirAll("../out/logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	// создание пути для файла с логами
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.00000")
	logFilePath := filepath.Join("../out/logs", fmt.Sprintf("%s.log", timestamp))

	// открываем поток записи файла
	logfile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}

	// создаем конфиг для encoder
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.00000")

	// создаем сам encoder
	encoder := zapcore.NewConsoleEncoder(cfg)

	// настраиваем вывод в консоль и в файл
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl),
		zapcore.NewCore(encoder, zapcore.AddSync(logfile), lvl),
	)

	// создаем logger
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return logger, logfile.Close, nil
}
