package logger

import (
	"go.uber.org/zap"
)

//TODO: we can use some logger

func LoggerInit() {
	log, _ := zap.NewDevelopment()
	log.Debug("This is a DEBUG message")
	log.Info("This is an INFO message")
}

