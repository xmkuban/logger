package logger

import (
	"errors"
	"io"

	"github.com/cihub/seelog"
)

const (
	DEFAULT_SEELOG_FORMAT = `
	<seelog type="sync">
		<outputs formatid="main">
			<console />
    	</outputs>
    	<formats>
    		<format id="main" format="[%Date(2006-01-02 15:04:05)][%Lev][%File:%Line] %Msg%n"/>
    	</formats>
	</seelog>
	`
)

var SeeLogger LoggerInterface

func InitSeelog() (err error) {
	_logger, err := seelog.LoggerFromConfigAsString(DEFAULT_SEELOG_FORMAT)
	if err != nil {
		return err
	}
	_logger.SetAdditionalStackDepth(2)
	SeeLogger = GetWrapSeelogger(_logger)
	return nil
}

type seelogWrapLogger struct {
	handler seelog.LoggerInterface
}

type seelogWrapLoggerIOWriter struct {
	handler seelog.LoggerInterface
}

func (w *seelogWrapLoggerIOWriter) Write(p []byte) (n int, err error) {
	if w.handler != nil {
		w.handler.Debugf("logger writer:%s", p)
		return len(p), nil
	}
	return 0, errors.New("handler nil")
}

func (l *seelogWrapLogger) Tracef(format string, params ...interface{}) {
	l.handler.Tracef(format, params...)
}

func (l *seelogWrapLogger) Debugf(format string, params ...interface{}) {
	l.handler.Debugf(format, params...)
}

func (l *seelogWrapLogger) Infof(format string, params ...interface{}) {
	l.handler.Infof(format, params...)
}

func (l *seelogWrapLogger) Warnf(format string, params ...interface{}) error {
	return l.handler.Warnf(format, params...)
}

func (l *seelogWrapLogger) Errorf(format string, params ...interface{}) error {
	return l.handler.Errorf(format, params...)
}

func (l *seelogWrapLogger) Criticalf(format string, params ...interface{}) error {
	return l.handler.Criticalf(format, params...)
}

func (l *seelogWrapLogger) Trace(v ...interface{}) {
	l.handler.Trace(v...)
}

func (l *seelogWrapLogger) Debug(v ...interface{}) {
	l.handler.Debug(v...)
}

func (l *seelogWrapLogger) Info(v ...interface{}) {
	l.handler.Info(v...)
}

func (l *seelogWrapLogger) Warn(v ...interface{}) error {
	return l.handler.Warn(v...)
}

func (l *seelogWrapLogger) Error(v ...interface{}) error {
	return l.handler.Error(v...)
}

func (l *seelogWrapLogger) Critical(v ...interface{}) error {
	return l.handler.Critical(v...)
}

func (l *seelogWrapLogger) GetStandardIOWriter() io.Writer {
	return &seelogWrapLoggerIOWriter{
		handler: l.handler,
	}
}

func GetWrapSeelogger(l seelog.LoggerInterface) LoggerInterface {
	//l.SetAdditionalStackDepth(1)
	return &seelogWrapLogger{
		handler: l,
	}
}
