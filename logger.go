package logger

import (
	//"os"
	//"path/filepath"
	//"sync"

	//l4g "git.ikuban.com/server/until-repository/log4go"
	"io"

	"github.com/go-sql-driver/mysql"
	//"github.com/robscc/mysql"
)

func SetLogger(log LoggerInterface) {
	CurrLogger = log
}

func init() {
	InitBeegoLogByConsole(7)
	//InitSeelog()
	//SetLogger(SeeLogger)
}

// Tracef formats message according to format specifier
// and writes to log with level = Trace.
func Tracef(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Tracef(format, params...)
	}
}

// Debugf formats message according to format specifier
// and writes to log with level = Debug.
func Debugf(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Debugf(format, params...)
	}
}

// Infof formats message according to format specifier
// and writes to log with level = Info.
func Infof(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Infof(format, params...)
	}
}

// Warnf formats message according to format specifier
// and writes to log with level = Warn.
func Warnf(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Warnf(format, params...)
	}
}

// Errorf formats message according to format specifier
// and writes to log with level = Error.
func Errorf(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Errorf(format, params...)
	}
}

// Criticalf formats message according to format specifier
// and writes to log with level = Critical.
func Criticalf(format string, params ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Criticalf(format, params...)
	}
}

// Trace formats message using the default formats for its operands
// and writes to log with level = Trace
func Trace(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Trace(v...)
	}
}

// Debug formats message using the default formats for its operands
// and writes to log with level = Debug
func Debug(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Debug(v...)
	}
}

// Info formats message using the default formats for its operands
// and writes to log with level = Info
func Info(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Info(v...)
	}
}

// Warn formats message using the default formats for its operands
// and writes to log with level = Warn
func Warn(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Warn(v...)
	}
}

// Error formats message using the default formats for its operands
// and writes to log with level = Error
func Error(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Error(v...)
	}
}

// Critical formats message using the default formats for its operands
// and writes to log with level = Critical
func Critical(v ...interface{}) {
	if CurrLogger != nil {
		CurrLogger.Critical(v...)
	}
}

type iowriter struct{}

func (w *iowriter) Write(p []byte) (n int, err error) {
	CurrLogger.Debugf("iowriter:%s", p)
	return len(p), nil
}

func GetDebugIOWriter() io.Writer {
	return &iowriter{}
}

type mysqlLogger struct{}

func (l *mysqlLogger) Print(v ...interface{}) {

	if CurrLogger != nil {
		if len(v) == 1 {
			CurrLogger.Error(v...)
		} else {
			CurrLogger.Errorf(v[0].(string), v[1:]...)
		}
	}
}

func GetMySQLLogger() mysql.Logger {
	return &mysqlLogger{}
}
