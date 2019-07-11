package logger

import (
	"io"
	"log"
	"path"

	"github.com/go-mach/machinery/pkg/config"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Logger represent the minimal set of func to set a logger for the Machinery
type Logger interface {
	SetOutput(w io.Writer)
	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Panic(i ...interface{})
	Panicf(format string, args ...interface{})
}

var loggerInstance *logrus.Logger
var conf config.Log

// NewLogger returns the logger instance. Initialize the instance only once.
func NewLogger(configuration *config.Log) Logger {
	if loggerInstance == nil {
		loggerInstance = logrus.New()
		log.Println("CECKING configuration log != nil")
		if configuration != nil {
			log.Println("configuration log != nil")
			conf := *configuration
			log.Printf("conf: %v", conf)
			// file log with rotation
			rfh, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
				Filename:   path.Join(conf.Path, conf.Filename),
				MaxSize:    conf.MaxSize,
				MaxBackups: conf.MaxBackups,
				MaxAge:     conf.MaxAge,
				Level:      parseLevel(),
				Formatter:  logFormatter(),
			})

			if err != nil {
				panic(err)
			}

			loggerInstance.AddHook(rfh)

			// console log
			if conf.Console.Enabled {
				loggerInstance.SetLevel(parseLevel())
				loggerInstance.SetOutput(colorable.NewColorableStdout())
				loggerInstance.SetFormatter(consoleFormatter())
			}
		} else {
			log.Println("configuration log == nil")
			// default logger
			Formatter := new(logrus.TextFormatter)
			Formatter.TimestampFormat = "02-01-2006 15:04:05"
			Formatter.FullTimestamp = true
			logrus.SetFormatter(Formatter)
		}
	}

	loggerInstance.Debug("Config and Logger initialized")

	return loggerInstance
}

func parseLevel() logrus.Level {
	var logLevel logrus.Level

	logLevel, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		panic(err)
	}

	return logLevel
}

func logFormatter() logrus.Formatter {
	if conf.JSON {
		return &logrus.JSONFormatter{
			TimestampFormat: conf.TimestampFormat,
		}
	}

	return &prefixed.TextFormatter{
		DisableColors:   true,
		ForceColors:     false,
		TimestampFormat: conf.TimestampFormat,
		FullTimestamp:   conf.FullTimestamp,
		ForceFormatting: conf.ForceFormatting,
	}
}

func consoleFormatter() logrus.Formatter {
	return &prefixed.TextFormatter{
		DisableColors:   conf.Console.DisableColors,
		ForceColors:     conf.Console.Colors,
		TimestampFormat: conf.TimestampFormat,
		FullTimestamp:   conf.FullTimestamp,
		ForceFormatting: conf.ForceFormatting,
	}
}
