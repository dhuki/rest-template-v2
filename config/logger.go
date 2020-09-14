package config

import (
	"os"

	"github.com/dhuki/rest-template-v2/custom"
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
)

func NewLogger() log.Logger {
	var loggerLogrus *logrus.Logger
	{
		// set configuration default for logrus
		loggerLogrus = logrus.New()
		// set my format for log logrus
		// kinda a fancy using logrus
		loggerLogrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "02-01-2006 15:04:05",
			FullTimestamp:   true,
		})
		// it is used to set output log to file with standardoutput(stdout)
		// by default it set output as standarderror (stderr)
		// usually when we run in linux system we use command "nohup ./main.go > log.log 2>&1 &"
		// it will write to file log.log merged with stderr from system if there is panics
		// source : https://stackoverflow.com/questions/3385201/confused-about-stdin-stdout-and-stderr#:~:text=If%20my%20understanding%20is%20correct,all%20the%20exceptions%20are%20entered.
		// #0 stdin, #1 stdout, #2 stderr
		loggerLogrus.SetOutput(os.Stdout)
	}

	var loggerKit log.Logger
	{
		// change loggrus to log go kit
		loggerKit = custom.NewLogrusLogger(loggerLogrus)
		// it will blocked if there is multiple go routing use logger
		// only one goroutine will be allowed to log to the wrapped logger at that time.
		// until it available again
		loggerKit = log.NewSyncLogger(loggerKit)
		// it will add default specs format error
		loggerKit = log.With(loggerKit,
			"caller", log.DefaultCaller,
		)
	}
	return loggerKit
}
