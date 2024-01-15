package logger

import (
	"log"
	"path/filepath"
	"std_exporter/common"

	"github.com/sirupsen/logrus"
)

const (
	InfoLevel  = logrus.InfoLevel
	ErrorLevel = logrus.ErrorLevel
	DebugLevel = logrus.DebugLevel
)

type Logger struct {
	*logrus.Logger
}

var StdLogger *Logger

func initStdLogger() {
	logger := logrus.StandardLogger()
	logger.SetLevel(InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{})
	StdLogger = &Logger{
		logger,
	}
	log.SetOutput(logger.Writer())
}

func GetStdLogger() *Logger {
	return StdLogger
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "info":
		l.Logger.SetLevel(InfoLevel)
	case "debug":
		l.Logger.SetLevel(DebugLevel)
	default:
		l.Logger.SetLevel(ErrorLevel)
	}
}

var LogPath string

func initLogPath() {
	LogPath = DefaultLogPath
	dir, file := filepath.Split(LogPath)
	if err := common.MakeFile(dir, file); err != nil {
		LogPath = common.NameSpace + ".log"
	}
}

func init() {
	initLogPath()
	initStdLogger()
	std := GetStdLogger()
	newLfsHook()
	std.AddHook(GetDefaultHook())
}
