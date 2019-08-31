package output

import (
	"encoding/json"

	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
)

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

	tunnel, err := tstate.FindTunnelByID(assignedID)

	if err != nil {
		logger.Disklog.Info(err)
	} else {
		tunnelJSON, err := json.MarshalIndent(tunnel, "", "    ")
		if err != nil {
			logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
		}
		logger.Disklog.Info(string(tunnelJSON))
	}
}

// ShowPool is in charge getting state and searching for a pool of tunnels
func ShowPool(poolName string) {
	tstate := manager.GetLastKnownState()

	tunnels, err := tstate.FindTunnelsByPool(poolName)

	if err != nil {
		logger.Disklog.Info(err)
	} else {
		tunnelsJSON, err := json.MarshalIndent(tunnels, "", "    ")
		if err != nil {
			logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
		}
		logger.Disklog.Info(string(tunnelsJSON))
	}
}
