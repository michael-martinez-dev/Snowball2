// internal/logger/logger.go

package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// InitLogger sets up the logrus logger
func InitLogger(logFilePath string) {
	// Open file in append mode, create if not exists
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Warnf("Could not create or open log file: %s", logFilePath)
	} else {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
	}

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
