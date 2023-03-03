package app_log

import (
	"io"
	"path/filepath"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//全局日志
type AppLogger struct {
	*logrus.Entry
	Out io.Writer
}

func InitAppLog(logPath string) *AppLogger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)

	infoPath := genLogPath(logPath, "request_logistics.log")
	infOut, err := rotate.New(
		infoPath+".%Y-%m-%d-%H", rotate.WithLinkName(infoPath),
		rotate.WithMaxAge(time.Duration(72)*time.Hour),
		rotate.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		return nil
	}

	debugPath := genLogPath(logPath, "biz_logistics.log")
	debugOut, err := rotate.New(
		debugPath+".%Y-%m-%d-%H", rotate.WithLinkName(debugPath),
		rotate.WithMaxAge(time.Duration(72)*time.Hour),
		rotate.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		return nil
	}

	errPath := genLogPath(logPath, "exception_logistics.err.log")
	errOut, err := rotate.New(
		errPath+".%Y-%m-%d-%H",
		rotate.WithLinkName(errPath),
		rotate.WithMaxAge(time.Duration(72)*time.Hour),
		rotate.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		return nil
	}

	hook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: debugOut,
			logrus.InfoLevel:  infOut,
			logrus.WarnLevel:  errOut,
			logrus.ErrorLevel: errOut,
			logrus.FatalLevel: errOut,
			logrus.PanicLevel: errOut,
		},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)

	hooks := make(logrus.LevelHooks)
	hooks.Add(hook)

	logger.ReplaceHooks(hooks)

	return &AppLogger{
		Entry: logger.WithFields(logrus.Fields{
			"module_name": "gin-api-frame",
		}),
		Out: infOut,
	}

}

func genLogPath(logPath, logFile string) string {
	return filepath.Join(logPath, logFile)
}
