package manager

import (
	"github.com/maxdobeck/sauced/logger"
)

func balance(lastState []Tunnel, pool string) int {
	logger.Disklog.Infof("Counting active tunnels from Last Known State: %v", lastState)
	count := 0
	for _, tunnel := range lastState {
		if tunnel.Metadata.Pool == pool {
			count++
		}
	}
	logger.Disklog.Debugf("Currently %d %s open", count, pool)
	return count
}

func startTunnel(config Metadata) bool {
	state := getLastKnownState()
	curBal := balance(state.Tunnels, config.Pool)
	if curBal < config.Size {
		return true
	}
	return false
}
