package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const eventPrefix = "event: "

func Log(level, nameFunc, event string, err error, additionalParams ...interface{}) {
	// Создание файла для логов
	file, errFile := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errFile != nil {
		panic(errFile) // Не удалось открыть файл для логирования
	}
	
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // Формат JSON
		zapcore.AddSync(file),                                    // Запись в файл
		zap.InfoLevel,                                           // Уровень логирования
	)

	logger := zap.New(core)
	defer logger.Sync() // Синхронизация записей перед завершением

	switch level {
	case "Info":
		logger.Info(eventPrefix + event)
	case "Error":
		logger.Error(
			err.Error(),
			zap.String(eventPrefix, event),
			zap.String("func", nameFunc),
			zap.Any("param", additionalParams),
		)
	case "Warning":
		logger.Warn(
			eventPrefix+event,
			zap.String("func", nameFunc),
			zap.Any("param", additionalParams),
		)
	}
}
