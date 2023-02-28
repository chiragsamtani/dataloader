package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func NewLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrusLevel := logrus.WarnLevel
	logger.Out = os.Stdout
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "debug":
		logrusLevel = logrus.DebugLevel
	case "warn":
		logrusLevel = logrus.WarnLevel
	case "error":
		logrusLevel = logrus.ErrorLevel
	}
	logger.SetLevel(logrusLevel)
	return logger
}
