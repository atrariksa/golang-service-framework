package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func GetLogger(logDir string) *logrus.Logger {
	if logger != nil {
		logger.Warn("GetLogger is recalled!")
		return logger
	}
	logger = logrus.New()
	ljack := &lumberjack.Logger{
		Filename:   logDir,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,     //days
		Compress:   false, // disabled by default
	}
	mWriter := io.MultiWriter(os.Stdout, ljack)
	logger.SetOutput(mWriter)

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	return logger
}
