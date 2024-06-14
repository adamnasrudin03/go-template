package driver

import (
	"github.com/adamnasrudin03/go-template/app/configs"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func Logger(config *configs.Configs) *logrus.Logger {
	logger := logrus.New()

	level := logrus.InfoLevel
	switch config.App.Env {
	case "dev":
		level = logrus.TraceLevel
	}

	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	logger.Out = file
	// }

	logger.SetLevel(level)
	logger.SetFormatter(&ecslogrus.Formatter{})

	return logger
}
