package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Disklog logs to disk w/ this instance.  Logfile
// location is setup via the "logfile" flag
var Disklog = log.New()

// SetupLogfile will take in the user spec'd flag
// and attempt to create a logfile if one does not exist
func SetupLogfile(logfile string, verbose bool) {
	Disklog.SetFormatter(&log.TextFormatter{})
	if verbose {
		Disklog.SetLevel(log.DebugLevel)
	} else {
		Disklog.SetLevel(log.InfoLevel)
	}

	Disklog.Out = os.Stdout
	// May be a logrus bug but the first log output has to be Info or it doesn't format in a pretty way  :shrug:
	Disklog.Info("Looking for logfile: ", logfile)
	if _, err := os.Stat(logfile); err == nil {
		Disklog.Debug("Using existing file: ", logfile)
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		multi := io.MultiWriter(os.Stdout, file)

		if err != nil {
			log.Warn("Failed to log to file, using default stdout")
		} else {
			Disklog.Out = multi
			Disklog.Debug("Started program and now writing to ", logfile)
		}
	} else if os.IsNotExist(err) {
		log.Debug(logfile, " NOT found: ", err)
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		multi := io.MultiWriter(os.Stdout, file)

		if err != nil {
			log.Warn("Failed to log to file, using default stderr")
		} else {
			Disklog.Out = multi
			Disklog.Debugf("Started program and now writing to %s.", logfile)
		}
	} else {
		log.Warn("unable to set the logfile to", logfile)
	}
}
