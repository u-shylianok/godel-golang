package main

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	logger.Info("application started")

	verbosePrt := flag.Bool("v", false, "verbose mode - shows all debug information")

	flag.Parse()

	if verbosePrt != nil && *verbosePrt {
		logger.SetLevel(logrus.DebugLevel)
	}

	logger.Debug("start doing some job..")

	for i := 0; i < 5; i++ {
		logger.Debugf("doing job #%d", i)
		time.Sleep(time.Second / 2)
	}

	logger.Info("application closed")
}
