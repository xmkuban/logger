package logger

import (
	"io"
	"os"
	"strings"

	"fmt"

	"time"

	"git.ikuban.com/server/until-repository/logger/beego/logs"
)

func InitBeegoLogByConsole(level int) {
	if level == 0 {
		level = logs.LevelDebug
	}
	InitBeegoLog(logs.AdapterConsole, fmt.Sprintf(`{"level":%d,"color":true}`, level))
}

func InitBeegoLogByConsoleDinding(webhookURL string, level int, dLevel int) {
	if level == 0 {
		level = logs.LevelDebug
	}
	if dLevel == 0 {
		dLevel = logs.LevelDebug
	}
	InitBeegoLogByConsole(level)
	err := logs.SetLogger(logs.AdapterDingding, fmt.Sprintf(`{"webhook_url":"%s","level":%d}`, webhookURL, dLevel))
	if err != nil {
		fmt.Println(err)
	}
}

func InitBeegoLogByDinding(webhookURL string, level int) {
	if level == 0 {
		level = logs.LevelDebug
	}
	InitBeegoLog(logs.AdapterDingding, fmt.Sprintf(`{"webhook_url":%s,"level":%d}`, webhookURL, level))
}

func InitBeegoLogByFile(fileName string, level int) {
	if level == 0 {
		level = logs.LevelDebug
	}
	if fileName == "" {
		fileName = fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	}
	if !strings.HasSuffix(fileName, ".log") {
		fileName = fileName + ".log"
	}
	fileConf := fmt.Sprintf(`{"level":%d,"filename":"%s","daily":false,"maxlines":100000,"color":true}`, level, fileName)

	InitBeegoLog(logs.AdapterFile, fileConf)
}

func InitBeegoLog(adapter string, config string) {
	beelogger := BeeLogger{}
	logs.Reset()
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(5)

	logs.SetLogger(adapter, config)
	SetLogger(&beelogger)
}

type BeeLogger struct {
	logs.BeeLogger
}

func (l *BeeLogger) Tracef(format string, params ...interface{}) {
	logs.Trace(fmt.Sprintf(format, params...))
}

func (l *BeeLogger) Debugf(format string, params ...interface{}) {
	logs.Debug(fmt.Sprintf(format, params...))
}

func (l *BeeLogger) Infof(format string, params ...interface{}) {
	logs.Info(fmt.Sprintf(format, params...))
}

func (l *BeeLogger) Warnf(format string, params ...interface{}) error {
	logs.Warn(fmt.Sprintf(format, params...))
	return nil
}

func (l *BeeLogger) Errorf(format string, params ...interface{}) error {
	logs.Error(fmt.Sprintf(format, params...))
	return nil
}

func (l *BeeLogger) Criticalf(format string, params ...interface{}) error {
	logs.Critical(fmt.Sprintf(format, params...))
	return nil
}

func (l *BeeLogger) Trace(v ...interface{}) {
	logs.Trace(fmt.Sprint(v...))
}

func (l *BeeLogger) Debug(v ...interface{}) {
	logs.Debug(fmt.Sprint(v...))
}

func (l *BeeLogger) Info(v ...interface{}) {
	logs.Info(fmt.Sprint(v...))
}

func (l *BeeLogger) Warn(v ...interface{}) error {
	logs.Warn(fmt.Sprint(v...))
	return nil
}

func (l *BeeLogger) Error(v ...interface{}) error {
	logs.Error(fmt.Sprint(v...))
	return nil
}

func (l *BeeLogger) Critical(v ...interface{}) error {
	logs.Critical(fmt.Sprint(v...))
	return nil
}

func (l *BeeLogger) GetStandardIOWriter() io.Writer {
	return os.Stdout
}
