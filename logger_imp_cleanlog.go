package logger

import (
	"fmt"
	"io"
	"os"
)

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

func SetCleanlogger() {
	SetLogger(&CleanLogger{})
}

type CleanLogger struct {
}

func (l *CleanLogger) Tracef(format string, params ...interface{}) {
	fmt.Print("t")
}

func (l *CleanLogger) Debugf(format string, params ...interface{}) {
	fmt.Print(".")
}

func (l *CleanLogger) Infof(format string, params ...interface{}) {
	fmt.Print("i")
}

func (l *CleanLogger) Warnf(format string, params ...interface{}) error {
	fmt.Print("w")
	return nil
}

func (l *CleanLogger) Errorf(format string, params ...interface{}) error {
	fmt.Print("e")
	return nil
}

func (l *CleanLogger) Criticalf(format string, params ...interface{}) error {
	fmt.Print("c")
	return nil
}

func (l *CleanLogger) Trace(v ...interface{}) {
	fmt.Print("t")
}

func (l *CleanLogger) Debug(v ...interface{}) {
	fmt.Print(".")
}

func (l *CleanLogger) Info(v ...interface{}) {
	fmt.Print("i")
}

func (l *CleanLogger) Warn(v ...interface{}) error {
	fmt.Print("w")
	return nil
}

func (l *CleanLogger) Error(v ...interface{}) error {
	fmt.Print("e")
	return nil
}

func (l *CleanLogger) Critical(v ...interface{}) error {
	fmt.Print("c")
	return nil
}

func (l *CleanLogger) GetStandardIOWriter() io.Writer {
	return os.Stdout
}
