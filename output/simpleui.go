package output

import (
	"encoding/json"

	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
)

// PrintState is called by a user to output some sort of detailed info on the active tunnels
func PrintState(human bool) {
	if !human {
		// print the default
		logger.Disklog.Debug("show called by the user.  Pruning then listing all tunnels.")
		manager.PruneState()
		showStateJSON()
	}
	if human {
		humanOutput()
	}
}

// showStateJSON will pretty print JSON
func showStateJSON() {
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
		logger.Disklog.Info(err)
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
		logger.Disklog.Info(err)
	} else {
		tunnelsJSON, err := json.MarshalIndent(tunnels, "", "    ")
		if err != nil {
			logger.Disklog.Warnf("Could not format JSON with Indents: %v", err)
		}
		logger.Disklog.Info(string(tunnelsJSON))
	}
}
