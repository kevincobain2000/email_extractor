package pkg

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	loggerOnce sync.Once
	log        *logrus.Logger
)

func Logger() *logrus.Logger {
	loggerOnce.Do(func() {
		fmt.Println("Logger initialized")
		log = logrus.New()
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
		log.SetReportCaller(false)
		log.SetLevel(logrus.InfoLevel)
	})
	return log
}
