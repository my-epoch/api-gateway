package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var l *logrus.Logger

func init() {
	l = logrus.New()
	l.Out = os.Stdout
	l.SetFormatter(&logrus.TextFormatter{})
}

func Trace(args ...interface{}) {
	l.Trace(args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	l.Fatal(args...)
}
