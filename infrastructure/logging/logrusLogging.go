package logging

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"portfolio-api/usecases"

	"os"
	"fmt"
)

type LogrusLogging struct {
	Client *logrus.Logger
}

func NewLogrusLogging() *LogrusLogging {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}

	return  &LogrusLogging{
		Client: logger,
	}
}

func (log *LogrusLogging) Info(msg string) {
	log.Client.Infoln(msg)
}

func (log *LogrusLogging) Error(uerr *usecases.UError) {
	trace := uerr.Msg
	_, file, line, ok := runtime.Caller(1)
	if ok {
		trace = trace +fmt.Sprintf("[%s:%d]", file, line)
	}

	log.Client.Errorln(trace)
}

func (log *LogrusLogging) Warning(uerr *usecases.UError) {
	log.Client.Warningln(uerr)
}


func (log *LogrusLogging) Debug(msg string) {
	log.Client.Debugln(msg)
}