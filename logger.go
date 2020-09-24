package logger

import (
	"sync"

	"github.com/go-sql-driver/mysql"
)

func SetLogger(log LoggerInterface) {
	CurrLogger = log
}

var onceDo = sync.Once{}

func initDefaultLog() {
	onceDo.Do(func() {
		InitBeegoLogByConsole(7)
	})
}

// Debugf formats message according to format specifier
// and writes to log with level = Debug.
func Debugf(format string, params ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Debugf(format, params...)
}

// Infof formats message according to format specifier
// and writes to log with level = Info.
func Infof(format string, params ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Infof(format, params...)
}

// Warnf formats message according to format specifier
// and writes to log with level = Warn.
func Warnf(format string, params ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Warnf(format, params...)
}

// Errorf formats message according to format specifier
// and writes to log with level = Error.
func Errorf(format string, params ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Errorf(format, params...)
}

// Debug formats message using the default formats for its operands
// and writes to log with level = Debug
func Debug(v ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Debug(v...)
}

// Info formats message using the default formats for its operands
// and writes to log with level = Info
func Info(v ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Info(v...)
}

// Warn formats message using the default formats for its operands
// and writes to log with level = Warn
func Warn(v ...interface{}) {
	if CurrLogger == nil {
		CurrLogger.Warn(v...)
	}
}

// Error formats message using the default formats for its operands
// and writes to log with level = Error
func Error(v ...interface{}) {
	if CurrLogger == nil {
		initDefaultLog()
	}
	CurrLogger.Error(v...)
}

type mysqlLogger struct{}

func (l *mysqlLogger) Print(v ...interface{}) {

	if CurrLogger == nil {
		initDefaultLog()
	}
	if len(v) == 1 {
		CurrLogger.Error(v...)
	} else {
		CurrLogger.Errorf(v[0].(string), v[1:]...)
	}
}

func GetMySQLLogger() mysql.Logger {
	return &mysqlLogger{}
}
