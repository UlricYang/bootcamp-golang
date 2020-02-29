package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

func initLogger() {
	logfile, err := os.OpenFile("rename.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = logfile
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func main() {
	initLogger()
	logger.WithFields(logrus.Fields{
		"file": "xxx",
	}).Error("cannot open file")
}
