package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Log — глобальный логгер, доступный во всем проекте.
var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	Log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Log.Level = logrus.InfoLevel
}
