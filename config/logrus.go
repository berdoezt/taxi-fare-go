package config

import (
	"log"

	"github.com/sirupsen/logrus"
)

func InitLogrus(cfg Config) {
	logger := logrus.StandardLogger()
	log.SetOutput(logger.Writer())

	level, err := logrus.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	if cfg.App.Debug {
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
}
