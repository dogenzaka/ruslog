package ruslog

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/dogenzaka/rotator"
)

const (
	APPENDER_DEFAULT = "Default"
	APPENDER_SIZE    = "Size"
	APPENDER_DAILY   = "Daily"
)

type Appender struct {
	Name  string
	Setup func(logger *Logger) *Logger
}

func defaultAppender(l *Logger) *Logger {
	log := logrus.New()
	log.Formatter = Formatters[l.Format].Formatter
	log.Level = GetLevel(l.Level)

	l.Call = func(level string, options map[string]interface{}, messages []string) {
		message := strings.Join(messages, " ") // dynamic message
		CallMethod(l, level, message, options)
	}
	l.logrus = log

	return l
}

func sizeRollingFileAppender(l *Logger) *Logger {

	l = defaultAppender(l)

	o := rotator.NewSizeRotator(l.FilePath)
	if l.RotationSize > 0 {
		o.RotationSize = l.RotationSize
	}
	if l.MaxRotation > 0 {
		o.MaxRotation = l.MaxRotation
	}

	l.logrus.Out = o
	return l
}

func dailyRollingFileAppender(l *Logger) *Logger {
	l = defaultAppender(l)
	l.logrus.Out = rotator.NewDailyRotator(l.FilePath)
	return l
}
