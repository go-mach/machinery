package gm

import (
	"io"
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
