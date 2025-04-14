package logger

import (
	"log"
	"os"
	"wallet-app/config"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	logfile, err := os.OpenFile(config.AppConfig.LOG_FILENAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Log.SetOutput(logfile)
	level, err := logrus.ParseLevel(config.AppConfig.LOG_LEVEL)
	if err != nil {
		Log.Warn("Setup incorrected log-level, using default: INFO")
		Log.SetLevel(logrus.InfoLevel)
	} else {
		Log.SetLevel(level)
	}
}
