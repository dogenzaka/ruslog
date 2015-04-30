package ruslog

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

const (
	ruslogFileName   = "github.com/dogenzaka/ruslog/ruslog.go"
	appenderFileName = "github.com/dogenzaka/ruslog/appender.go"
)

type (
	Logger struct {
		Name         string // logger name(uniq,required)
		Type         string // ruslog.APPENDER_XXXX
		Format       string // ruslog.FORMATTER_XXXX
		Level        string // logrus.XXXXLevel.String()
		FilePath     string // outpu file path (optional)
		RotationSize int64  // size threshold of the rotation example 10M) 1024 * 1024 * 10 (optional)
		MaxRotation  int    // maximum count of the rotation (optional)
		AddFileInfo  bool   // add the file info to the log message (optional)

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
		APPENDER_DEFAULT: &Appender{
			Name:  APPENDER_DEFAULT,
			Setup: defaultAppender,
		},
		APPENDER_SIZE: &Appender{
			Name:  APPENDER_SIZE,
			Setup: sizeRollingFileAppender,
		},
		APPENDER_DAILY: &Appender{
			Name:  APPENDER_DAILY,
			Setup: dailyRollingFileAppender,
		},
	}

	// Manage ruslog(logrus) formatters
	Formatters = formatters{
		FORMATTER_SIMPLE: &Formatter{
			Name:      FORMATTER_SIMPLE,
			Formatter: &SimpleFormatter{},
		},
		//logrus
		FORMATTER_TEXT: &Formatter{
			Name:      FORMATTER_TEXT,
			Formatter: &logrus.TextFormatter{},
		},
		// logrus
		FORMATTER_JSON: &Formatter{
			Name:      FORMATTER_JSON,
			Formatter: &logrus.JSONFormatter{},
		},
	}
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
	l := Logging.loggers[name]
	// if name logger is not found, return default logger.
	if l == nil {
		l = &Logger{Name: name, Type: APPENDER_DEFAULT, Level: "Info", Format: FORMATTER_TEXT}
		Configure([]*Logger{l})
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

	if logger.AddFileInfo {
		if options == nil {
			options = map[string]interface{}{}
		}
		var info string
		for depth := 1; ; depth++ {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				info = "unknown"
				break
			}
			if !strings.HasSuffix(file, appenderFileName) && !strings.HasSuffix(file, ruslogFileName) {
				info = fmt.Sprintf("%s:%d", file, line)
				break
			}
		}
		options["file"] = info
	}

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
			fmt.Println("[RUSLOG-INFO] Default logging.", APPENDER_DEFAULT)
		}
		appender = Appenders[APPENDER_DEFAULT]
	}

	return appender.Setup(logger)
}

// Debug log output (goroutine)
func (l *Logger) Debug(options map[string]interface{}, messages ...string) {
	go l.Call("Debug", options, messages)
}

// Info log output (goroutine)
func (l *Logger) Info(options map[string]interface{}, messages ...string) {
	go l.Call("Info", options, messages)
}

// Warn log output (goroutine)
func (l *Logger) Warn(options map[string]interface{}, messages ...string) {
	go l.Call("Warn", options, messages)
}

// Error log output (goroutine)
func (l *Logger) Error(options map[string]interface{}, messages ...string) {
	go l.Call("Error", options, messages)
}

// Fatal log output (goroutine)
func (l *Logger) Fatal(options map[string]interface{}, messages ...string) {
	go l.Call("Fatal", options, messages)
}

///

// Debug log output (not goroutine)
func (l *Logger) DebugSync(options map[string]interface{}, messages ...string) {
	l.Call("Debug", options, messages)
}

// Info log output (not goroutine)
func (l *Logger) InfoSync(options map[string]interface{}, messages ...string) {
	l.Call("Info", options, messages)
}

// Warn log output (not goroutine)
func (l *Logger) WarnSync(options map[string]interface{}, messages ...string) {
	l.Call("Warn", options, messages)
}

// Error log output (not goroutine)
func (l *Logger) ErrorSync(options map[string]interface{}, messages ...string) {
	l.Call("Error", options, messages)
}

// Fatal log output (not goroutine)
func (l *Logger) FatalSync(options map[string]interface{}, messages ...string) {
	l.Call("Fatal", options, messages)
}

///

// log.Logger.Output like (gorutine)
func (l *Logger) Output(calldepth int, s string) error {
	go l.Call("Info", nil, []string{s})
	return nil
}

// io.Write like (gorutine)
func (l *Logger) Write(p []byte) (n int, err error) {
	go l.logrus.Out.Write(p)
	return 0, nil // The exception is ignored
}
