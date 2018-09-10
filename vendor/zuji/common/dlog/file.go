package dlog

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/x-cray/logrus-prefixed-formatter"
	"encoding/json"
)

var logEnable bool

func OpenLog(path string, name string) {
	logs, err := rotatelogs.New(
		path+name+"-%Y%m%d.log",
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(90),
	)

	if err != nil {
		return
	}

	logEnable = true

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logs,
		logrus.InfoLevel:  logs,
		logrus.WarnLevel:  logs,
		logrus.ErrorLevel: logs,
		logrus.FatalLevel: logs,
		logrus.PanicLevel: logs,
	}, &prefixed.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 -07",
	})
	logrus.AddHook(lfHook)
}

func WriteLog(a ...interface{}) {
	if logEnable {
		logrus.Info(a...)
	}
}

func WriteLogAsync(a ...interface{}) {
	go WriteLog(a...)
}

func WriteJson(v interface{}) {
	if logEnable {
		b, err := json.Marshal(v)
		if err == nil {
			logrus.Info(string(b))
		}
	}
}

func WriteJsonAsync(v interface{}) {
	go WriteJson(v)
}
