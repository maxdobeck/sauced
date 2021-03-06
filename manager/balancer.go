package manager

import (
	"github.com/mdsauce/sauced/logger"
)

func balance(lastState []Tunnel, pool string) int {
	logger.Disklog.Debugf("Counting active tunnels from Last Known State: %v", lastState)
	count := 0
	for _, tunnel := range lastState {
		if tunnel.Metadata.Pool == pool {
			count++
		}
	}
	logger.Disklog.Debugf("Currently %d %s open", count, pool)
	return count
}

func vacancy(config Metadata) bool {
	state := GetLastKnownState()
	curBal := balance(state.Tunnels, config.Pool)
	if curBal < config.Size {
		return true
	}
	return false
}
