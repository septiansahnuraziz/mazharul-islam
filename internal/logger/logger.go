package logger

import (
	"github.com/mazharul-islam/config"
	"github.com/sirupsen/logrus"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
)

func SetupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		File:           true,
	}

	if config.EnvironmentMode() != "prod" {
		formatter = runtime.Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors:   true,
				FullTimestamp: true,
			},
			Line: true,
			File: true,
		}
	}

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
