package ruslog

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Sirupsen/logrus"
)

type (
	Logger struct {
		Name         string
		Type         string
		Format       string
		Level        string
		FilePath     string
		RotationSize int64
		MaxRotation  int

		Call func(level string, options map[string]interface{}, messages []string)

		logrus *logrus.Logger
	}

	logging struct {
		loggers map[string]*Logger
	}

	appenders  map[string]*Appender
	formatters map[string]*Formatter
)

var (
	// debug flag
	DEBUG bool = false

	// ruslog package instance
	Logging *logging = &logging{
		loggers: make(map[string]*Logger),
	}

	// Manage ruslog(logrus) Appenders
	Appenders = appenders{
		DEFAULT: &Appender{
			Name:  DEFAULT,
			Setup: defaultAppender,
		},
		SIZE: &Appender{
			Name:  SIZE,
			Setup: sizeRollingFileAppender,
		},
		DAILY: &Appender{
			Name:  DAILY,
			Setup: dailyRollingFileAppender,
		},
	}

	// Manage ruslog(logrus) formatters
	Formatters = func() formatters {
		ret := formatters{
			SIMPLE: &Formatter{
				Name:      SIMPLE,
				Formatter: &SimpleFormatter{},
			},
			//logrus
			TEXT: &Formatter{
				Name:      TEXT,
				Formatter: &logrus.TextFormatter{},
			},
			// logrus
			JSON: &Formatter{
				Name:      JSON,
				Formatter: &logrus.JSONFormatter{},
			},
		}

		return ret
	}()
)

// load ruslog
func Configure(loggers []*Logger) *logging {
	for _, logger := range loggers {
		Logging.loggers[logger.Name] = logger.Setup()
		if DEBUG {
			fmt.Printf("[RUSLOG-INFO] Add logging. %s=%s\n", logger.Name, GetLevel(logger.Level))
		}
	}

	return Logging
}

// Added the Formatter to manage
func AddFormatter(formatter *Formatter) *Formatter {
	Formatters[formatter.Name] = formatter
	return Formatters[formatter.Name]
}

// Added the Appender to manage
func AddAppender(appender *Appender) *Appender {
	Appenders[appender.Name] = appender
	return Appenders[appender.Name]

}

func GetLogger(name string) *Logger {
	l := Logging[name]
	// if name logger is not found, return default logger.
	if l == nil {
		l = &Logger{Type: DEFAULT}
		return l.Setup()
	}
	return l
}

// Get the logging level value
func GetLevel(level string) logrus.Level {
	l, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		panic(err)
	}
	return l
}

// Call logger method for a given level
func CallMethod(logger *Logger, level string, message string, options map[string]interface{}) {
	//fmt.Println("CallMethod", logger, level, message, options)
	loggerLogrus := logger.logrus

	entry := loggerLogrus.WithFields(options)
	methodName := level
	method := reflect.ValueOf(entry).MethodByName(methodName)

	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf(message)}
		method.Call(args)
	} else {
		entry.Debug(message)
	}
}

// -- Logger

// Setup appender
func (logger *Logger) Setup() *Logger {

	appender := Appenders[logger.Type]
	if appender == nil {
		if DEBUG {
			fmt.Println("[LOGRUSH-INFO] Default logging.", DEFAULT)
		}
		appender = Appenders[DEFAULT]
	}

	return appender.Setup(logger)
}

// Debug log output
func (l *Logger) Debug(options map[string]interface{}, messages ...string) {
	l.Call("Debug", options, messages)
}

// Info log output
func (l *Logger) Info(options map[string]interface{}, messages ...string) {
	l.Call("Info", options, messages)
}

// Warn log output
func (l *Logger) Warn(options map[string]interface{}, messages ...string) {
	l.Call("Warn", options, messages)
}

// Error log outputz
func (l *Logger) Error(options map[string]interface{}, messages ...string) {
	l.Call("Error", options, messages)
}

// Fatal log output
func (l *Logger) Fatal(options map[string]interface{}, messages ...string) {
	l.Call("Fatal", options, messages)
}
