package ruslog

import (
	"github.com/Sirupsen/logrus"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type (
	SimpleFormatter struct {}
)

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "[%s]", entry.Time.String())
	fmt.Fprintf(b, " [%s]", strings.ToUpper(entry.Level.String()))
	fmt.Fprintf(b, " %s", entry.Message)

	for key, value := range entry.Data {
		if key != "time" && key != "level" && key != "msg" {
			fmt.Fprintf(b, " %s=%s", key, value)
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func SimpleAppender(logger *Logger) *Logger{

	logrusLogger := logrus.New()

	//logrusLogger.Formatter = &SimpleFormatter{}
	formatter, _ := GetFormatter("Default").Formatter.(logrus.Formatter)
	logrusLogger.Formatter = formatter

	logrusLogger.Level = GetLevel(logger.Level)
	logrusLogger.Out = os.Stdout

	logger.Call = func(level string, options map[string]interface {}, messages []string) {
		message := strings.Join(messages, " ") // dynamic message
		CallMethod(logger, level, message, options)
	}

	logger.Logrus = logrusLogger

	return logger
}
