package logger

import (
	"go.uber.org/zap"
)

const eventPrefix = "event: "

func Log(level, nameFunc, event string, err error, additionalParams ...interface{}) {
	logger, errZap := zap.NewDevelopment()
	if errZap != nil {
		panic(errZap) // Не удалось создать логгер
	}

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
