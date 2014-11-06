package ruslog

import (
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/dogenzaka/rotator"
)

func defaultAppender(l *Logger) *Logger {

	logrus := logrus.New()

	logrus.Formatter = &SimpleFormatter{}

	logrus.Level = GetLevel(l.Level)
	logrus.Out = os.Stdout

	l.Call = func(level string, options map[string]interface{}, messages []string) {
		message := strings.Join(messages, " ") // dynamic message
		CallMethod(l, level, message, options)
	}

	l.Logrus = logrus

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

	l.Logrus.Out = o
	return l
}

func dailyRollingFileAppender(l *Logger) *Logger {
	l = defaultAppender(l)
	l.Logrus.Out = rotator.NewDailyRotator(l.FilePath)
	return l
}
