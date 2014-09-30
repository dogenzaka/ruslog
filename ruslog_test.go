package ruslog

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"reflect"
	"github.com/Sirupsen/logrus"
)

func TestRuslog(t *testing.T) {
	Convey("Another pattern log output", t, func() {
		loggers := []*Logger{
			&Logger{ Name: "Test0", Type: DEFAULT, Level: "INFO", Format: DEFAULT},
			&Logger{ Name: "Test1", Type: DEFAULT, Level: "Info", Format: TEXT},
			&Logger{ Name: "Test2", Type: DEFAULT, Level: "info", Format: JSON},
			&Logger{ Name: "Test3", Type: DEFAULT, Level: "DEBUG", Format: DEFAULT},
		}

		Convey("Configure run", func() {
			_loggings := Configure(loggers)
			_appenders :=  GetAppenderAll()
			_formatters := GetFormatterAll()

			fmt.Println("logging:", _loggings)
			fmt.Println("appenders:", _appenders)
			fmt.Println("formatters:", _formatters)

			So(len(_loggings), ShouldEqual, 4)
			So(len(_appenders), ShouldEqual, 1)
			So(len(_formatters), ShouldEqual, 3)
		})

		Convey("Get Appender[s]", func() {
			_appenders :=  GetAppenderAll()
			_appender := _appenders[DEFAULT]
			So(len(_appenders), ShouldEqual, 1)
			So(_appender.Name, ShouldEqual, DEFAULT)
		})

		Convey("Get Formatter[s]", func() {
			_formatters :=  GetFormatterAll()
			_formatter := _formatters[DEFAULT]
			So(len(_formatters), ShouldEqual, 3)
			So(_formatter.Name, ShouldEqual, DEFAULT)
		})

		Convey("Output log", func() {

			logger := GetLogger("Test0")
			logger.Debug(nil, "output log level:Debug")
			logger.Info(nil, "output log level:Info")
			logger.Warn(nil, "output log level:Warn")
			logger.Error(nil, "output log level:Error")
			//logger.Fatal(nil, "output log level:Fatal")

			ok := reflect.ValueOf(logger).Elem().Type() == reflect.ValueOf(&Logger{}).Elem().Type()

			So(ok, ShouldEqual, true)

		})

		Convey("Get Level", func () {
			So(GetLevel("Debug"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("DEBUG"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("DeBug"), ShouldEqual, logrus.DebugLevel)
			So(GetLevel("debug"), ShouldEqual, logrus.DebugLevel)

		})

		Convey("Get Level", func () {
			So(GetLevelStr(logrus.DebugLevel), ShouldEqual, "debug")
			So(GetLevelStr(logrus.InfoLevel), ShouldEqual, "info")
			So(GetLevelStr(logrus.WarnLevel), ShouldEqual, "warn")
			So(GetLevelStr(logrus.ErrorLevel), ShouldEqual, "error")
			So(GetLevelStr(logrus.FatalLevel), ShouldEqual, "fatal")
			So(GetLevelStr(logrus.PanicLevel), ShouldEqual, "panic")


		})
	})
}
