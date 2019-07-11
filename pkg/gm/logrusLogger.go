package gm

import (
	"path"

	"github.com/go-mach/gm/config"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logger *logrus.Logger
var logConfig = config.GetConfiguration()

// NewLogger returns the logger instance. Initialize the instance only once.
func NewLogger() Logger {
	if logger == nil {
		logger = logrus.New()

		if config.IsSet("log") {
			// file log with rotation
			rfh, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
				Filename:   path.Join(logConfig.Log.Path, logConfig.Log.Filename),
				MaxSize:    logConfig.Log.MaxSize,
				MaxBackups: logConfig.Log.MaxBackups,
				MaxAge:     logConfig.Log.MaxAge,
				Level:      parseLevel(),
				Formatter:  logFormatter(),
			})

			if err != nil {
				panic(err)
			}

			logger.AddHook(rfh)

			// console log
			if logConfig.Log.Console.Enabled {
				logger.SetLevel(parseLevel())
				logger.SetOutput(colorable.NewColorableStdout())
				logger.SetFormatter(consoleFormatter())
			}
		} else {
			// default logger
			Formatter := new(logrus.TextFormatter)
			Formatter.TimestampFormat = "02-01-2006 15:04:05"
			Formatter.FullTimestamp = true
			logrus.SetFormatter(Formatter)
		}
	}

	logger.Debug("Config and Logger initialized")

	return logger
}

func parseLevel() logrus.Level {
	var logLevel logrus.Level

	logLevel, err := logrus.ParseLevel(logConfig.Log.Level)
	if err != nil {
		panic(err)
	}

	return logLevel
}

func logFormatter() logrus.Formatter {
	if logConfig.Log.JSON {
		return &logrus.JSONFormatter{
			TimestampFormat: logConfig.Log.TimestampFormat,
		}
	}

	return &prefixed.TextFormatter{
		DisableColors:   true,
		ForceColors:     false,
		TimestampFormat: logConfig.Log.TimestampFormat,
		FullTimestamp:   logConfig.Log.FullTimestamp,
		ForceFormatting: logConfig.Log.ForceFormatting,
	}
}

func consoleFormatter() logrus.Formatter {
	return &prefixed.TextFormatter{
		DisableColors:   logConfig.Log.Console.DisableColors,
		ForceColors:     logConfig.Log.Console.Colors,
		TimestampFormat: logConfig.Log.TimestampFormat,
		FullTimestamp:   logConfig.Log.FullTimestamp,
		ForceFormatting: logConfig.Log.ForceFormatting,
	}
}
