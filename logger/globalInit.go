package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Disklog logs to disk w/ this instance.  Logfile
// location is setup via the "logfile" flag
var Disklog = log.New()

// SetupLogfile will take in the user spec'd flag
// and attempt to create a logfile if one does not exist
func SetupLogfile(logfile string) {
	Disklog.SetFormatter(&log.TextFormatter{})
	Disklog.SetLevel(log.DebugLevel)
	Disklog.Out = os.Stdout
	Disklog.Info("Looking for logfile: ", logfile)
	if _, err := os.Stat(logfile); err == nil {
		Disklog.Debug("Found logfile: ", logfile)
		// ADD sauced LOG ROTATION HERE
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Info("Failed to log to file, using default stdout")
		} else {
			Disklog.Out = file
			Disklog.Info("Started program and now writing to file.")
		}
	} else if os.IsNotExist(err) {
		log.Info(logfile, " NOT found: ", err)
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Info("Failed to log to file, using default stderr")
		} else {
			Disklog.Out = file
			Disklog.Info("Started program and now writing to file.")
		}
	} else {
		log.Warn("unable to set the logfile to", logfile)
	}
}
