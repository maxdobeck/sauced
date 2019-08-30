package output

import (
	"encoding/json"

	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
)

// PrettyPrint is called by a user to output some sort of detailed info on the active tunnels
func PrettyPrint(id string, pool string) {
	state := manager.GetLastKnownState()
	if id != "" {
		showTunnelPretty(id, state)
	} else if pool != "" {
		showPoolPretty(pool)
	} else {
		showStatePretty()
	}
}

// ShowStateJSON will pretty print JSON
func ShowStateJSON() {
	last := manager.GetLastKnownState()
	lastJSON, err := json.MarshalIndent(last, "", "    ")
	if err != nil {
		logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
	}
	logger.Disklog.Info(string(lastJSON))
}

// ShowTunnelJSON is in charge getting state and searching for a single tunnel by ID.
func ShowTunnelJSON(assignedID string) {
	tstate := manager.GetLastKnownState()

	tunnels, err := tstate.FindTunnelsByID(assignedID)

	if err != nil {
		logger.Disklog.Warn(err, ", Tunnel Assigned ID: ", assignedID)
	} else {
		tunnelsJSON, err := json.MarshalIndent(tunnels, "", "    ")
		if err != nil {
			logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
		}
		logger.Disklog.Info(string(tunnelsJSON))
	}
}

// ShowPool is in charge getting state and searching for a pool of tunnels
func ShowPool(poolName string) {
	tstate := manager.GetLastKnownState()

	tunnels, err := tstate.FindTunnelsByPool(poolName)

	if err != nil {
		logger.Disklog.Warn(err, ", Pool Name: ", poolName)
	} else {
		tunnelsJSON, err := json.MarshalIndent(tunnels, "", "    ")
		if err != nil {
			logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
		}
		logger.Disklog.Info(string(tunnelsJSON))
	}
}
