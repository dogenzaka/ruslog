package ruslog

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"reflect"
	"strings"
)

const (
	// formatter types
	JSON    = "Json"
	TEXT    = "Text"
	DEFAULT = "Default"
)

type (
	Logging struct {
		loggers map[string]*Logger
	}

	Logger struct {
		Name   string
		Type   string
		Level  string
		Format string
		//		Directory string
		//		Pattern string
		//		Interval int64
		//		FileName string

		Logrus *logrus.Logger
		Call   func(level string, options map[string]interface{}, messages []string)
	}

	Appenders map[string]*Appender

	Appender struct {
		Name  string
		Setup func(logger *Logger) *Logger
	}

	Formatters map[string]*Formatter

	Formatter struct {
		Name      string
		Formatter interface{}
	}
)

var (
	// debug flag
	DEBUG bool = false

	// ruslog package instance
	logging *Logging = &Logging{
		loggers: make(map[string]*Logger),
	}

	//var Logging map[string]*Logger = make(map[string]*Logger)

	// Manage ruslog(logrus) Appenders
	appenders = Appenders{
		DEFAULT: &Appender{
			Name:  DEFAULT,
			Setup: SimpleAppender,
		},
	}

	// Manage ruslog(logrus) formatters
	formatters = func() Formatters {
		ret := Formatters{
			DEFAULT: &Formatter{
				Name:      DEFAULT,
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
func Configure(loggers []*Logger) *Logging {
	for _, logger := range loggers {
		logging.loggers[logger.Name] = logger.Setup()
		if DEBUG {
			fmt.Printf("[LOGRUSH-INFO] Add logging. %s=%s\n", logger.Name, GetLevel(logger.Level))
		}
	}

	return logging
}

func GetLogging() *Logging {
	return logging
}

// Added the Formatter to manage
func AddFormatter(formatter *Formatter) *Formatter {
	formatters[formatter.Name] = formatter
	return formatters[formatter.Name]
}

// Added the Appender to manage
func AddAppender(appender *Appender) *Appender {
	appenders[appender.Name] = appender
	return appenders[appender.Name]

}

// Get the Formatter to manage
func GetFormatter(name string) *Formatter {
	f := formatters[name]
	return f
}

// Get the Appender to manage
func GetAppender(name string) *Appender {
	return appenders[name]
}

// Get all the Appender  to manage
func GetAppenderAll() Appenders {
	return appenders
}

// Get all the Formatter  to manage
func GetFormatterAll() Formatters {
	return formatters
}

// Get the logging level value
func GetLevel(level string) logrus.Level {
	levelStr := strings.ToLower(level)

	ret, err := logrus.ParseLevel(levelStr)

	if err != nil {
		fmt.Println(err)
	}

	return ret

}

//// Get the logging level string value
func GetLevelStr(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "debug"
	case logrus.InfoLevel:
		return "info"
	case logrus.WarnLevel:
		return "warn"
	case logrus.ErrorLevel:
		return "error"
	case logrus.FatalLevel:
		return "fatal"
	case logrus.PanicLevel:
		return "panic"
	default:
		return ""
	}
}

// Get the logging logger
func GetLogger(name string) *Logger {
	ret := logging.loggers[name]
	return ret
}

// Call logger method for a given level
func CallMethod(logger *Logger, level string, message string, options map[string]interface{}) {
	//fmt.Println("CallMethod", logger, level, message, options)
	loggerLogrus := logger.Logrus

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

	appender := appenders[logger.Type]
	if appender == nil {
		if DEBUG {
			fmt.Println("[LOGRUSH-INFO] Default logging.", DEFAULT)
		}
		appender = GetAppender(DEFAULT)
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
