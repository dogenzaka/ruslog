package ruslog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRuslog(t *testing.T) {
	Convey("Another pattern log output", t, func() {
		loggers := []*Logger{
			&Logger{Name: "Test0", Type: APPENDER_DEFAULT, Level: logrus.InfoLevel.String(), Format: FORMATTER_SIMPLE},
			&Logger{Name: "Test1", Type: APPENDER_DEFAULT, Level: "Info", Format: FORMATTER_TEXT},
			&Logger{Name: "Test2", Type: APPENDER_DEFAULT, Level: "warn", Format: FORMATTER_JSON},
			&Logger{Name: "Test3", Type: APPENDER_DEFAULT, Level: "DEBUG", Format: FORMATTER_SIMPLE},
		}

		Convey("Configure run", func() {
			_logging := Configure(loggers)
			_appenders := Appenders
			_formatters := Formatters

			fmt.Println("logging:", _logging)
			fmt.Println("appenders:", _appenders)
			fmt.Println("formatters:", _formatters)

			So(len(_logging.loggers), ShouldEqual, 4)
			So(len(_appenders), ShouldEqual, 3)
			So(len(_formatters), ShouldEqual, 3)
		})

		Convey("Get Appender[s]", func() {
			_appenders := Appenders
			_appender := _appenders[APPENDER_DEFAULT]
			So(len(_appenders), ShouldEqual, 3)
			So(_appender.Name, ShouldEqual, APPENDER_DEFAULT)
		})

		Convey("Get Formatter[s]", func() {
			_formatters := Formatters
			_formatter := _formatters[FORMATTER_SIMPLE]
			So(len(_formatters), ShouldEqual, 3)
			So(_formatter.Name, ShouldEqual, FORMATTER_SIMPLE)
		})

		Convey("Output log", func() {
			fmt.Println("")
			logger := Logging.loggers["Test0"]
			logger.Debug(nil, "output log level:Debug")
			logger.Info(nil, "output log level:Info")
			logger.Warn(nil, "output log level:Warn")
			logger.Error(nil, "output log level:Error")
			//logger.Fatal(nil, "output log level:Fatal")

			ok := reflect.ValueOf(logger).Elem().Type() == reflect.ValueOf(&Logger{}).Elem().Type()

			So(ok, ShouldEqual, true)

		})

		Convey("Output multi log", func() {
			fmt.Println("\ntarget: 0")
			logger0 := Logging.loggers["Test0"]
			logger0.Debug(nil, "output multi log level:Debug 0")
			logger0.Info(nil, "output multi log level:Info 0")
			logger0.Warn(nil, "output multi log level:Warn 0")
			logger0.Error(nil, "output multi log level:Error 0")

			fmt.Println("\ntarget: 1")
			logger1 := Logging.loggers["Test1"]
			logger1.Debug(nil, "output multi log level:Debug 1")
			logger1.Info(nil, "output multi log level:Info 1")
			logger1.Warn(nil, "output multi log level:Warn 1")
			logger1.Error(nil, "output multi log level:Error 1")

			fmt.Println("\ntarget: 2")
			logger2 := Logging.loggers["Test2"]
			logger2.Debug(nil, "output multi log level:Debug 2")
			logger2.Info(nil, "output multi log level:Info 2")
			logger2.Warn(nil, "output multi log level:Warn 2")
			logger2.Error(nil, "output multi log level:Error 2")

			fmt.Println("\ntarget: 3")
			logger3 := Logging.loggers["Test3"]
			logger3.Debug(nil, "output multi log level:Debug 3")
			logger3.Info(nil, "output multi log level:Info 3")
			logger3.Warn(nil, "output multi log level:Warn 3")
			logger3.Error(nil, "output multi log level:Error 3")

		})

		Convey("Get Level", func() {
			So(GetLevel("Debug"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("DEBUG"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("DeBug"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("debug"), ShouldEqual, logrus.DebugLevel)

		})

	})
}
