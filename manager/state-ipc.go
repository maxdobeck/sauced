package manager

import (
	"github.com/maxdobeck/sauced/logger"
	"os"
)

// TunnelState is a json representation
// of the a tunnel as an OS process after it has launched
type TunnelState struct {
	PID      int
	SCBinary string
	Args     string
	Launch   string
}

func createIPCFile() *os.File {
	if _, err := os.Stat("/tmp/sauced-state.json"); err == nil {
		logger.Disklog.Debug("Found state file /tmp/sauced-state.json")
		// ADD sauced LOG ROTATION HERE
		file, err := os.OpenFile("/tmp/sauced-state.json", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logger.Disklog.Warn("Failed to open /tmp/sauced-state.json")
			return nil
		}
		return file
	} else if os.IsNotExist(err) {
		logger.Disklog.Info("/tmp/sauced-state.json not found ", err)
		file, err := os.OpenFile("/tmp/sauced-state.json", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logger.Disklog.Warn("Failed to open /tmp/sauced-state.json")
			return nil
		}
		return file
	} else {
		logger.Disklog.Warn("unable to obtain a pointer to /tmp/sauced-state.json")
		return nil
	}
}
