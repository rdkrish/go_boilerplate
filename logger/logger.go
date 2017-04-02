package logger

import (
	"github.com/rifflock/lfshook"
	log "github.com/Sirupsen/logrus"
	"go_boilerplate/config"
)

// Log object
var Log *log.Logger

// Init function for initializing the logger
func Init() *log.Logger {
	if Log != nil {
		return Log
	}
	c := config.GetConfig()
	log.SetLevel(log.DebugLevel)
	Log = log.New()
	logFile := c.GetString("LOG_FILE")
	CustomFormatter := new(log.TextFormatter)
	CustomFormatter.TimestampFormat = "2006-01-02 15:04:05"
	CustomFormatter.FullTimestamp = true
	Log.Formatter = CustomFormatter

	Log.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  logFile,
		log.ErrorLevel: logFile,
		log.WarnLevel:  logFile,
		log.DebugLevel: logFile,
	}))
	return Log
}

// GetLogger function to expose the Log object
func GetLogger() *log.Logger {
	return Log
}