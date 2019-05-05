package manager

import (
	"github.com/maxdobeck/sauced/logger"
)

// Balance takes in the config file reduces/increases
// the tunnels in each pool as needed
func Balance(config string) {
	logger.Disklog.Infof("Checking that there are not too few or too many tunnels per config file: %s", config)
}
