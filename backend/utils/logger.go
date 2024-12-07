package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)


var Logger *logrus.Logger


func InitLogger() {
	Logger = logrus.New()


	Logger.SetOutput(os.Stdout)

	Logger.SetFormatter(&logrus.JSONFormatter{})

	Logger.SetLevel(logrus.InfoLevel)
}

