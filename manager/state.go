package manager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/mdsauce/sauced/logger"
)

// LastKnownTunnels will a json object with
// all tunnels that were previously known to be alive
type LastKnownTunnels struct {
	Tunnels []Tunnel `json:"tunnels"`
}

// Tunnel is a json representation
// of an OS process after SC has launched
type Tunnel struct {
	PID        int       `json:"pid"`
	AssignedID string    `json:"assignedID"`
	SCBinary   string    `json:"scbinary"`
	Args       string    `json:"args"`
	LaunchTime time.Time `json:"launchtime"`
	Log        string    `json:"log"`
	Metadata   Metadata  `json:"metadata"`
}

// add later
// func (tun Tunnel) rotateLog() error {
// 	return nil
// }

// AddTunnel will record the state of the tunnel
// to the IPC file after the tunnel has launched as an OS process
func AddTunnel(launchArgs string, path string, PID int, meta Metadata, tunnelLog string, asgnID string) {
	createIPCFile()

	var tunnelsState LastKnownTunnels
	tun := Tunnel{PID, asgnID, path, launchArgs, time.Now().UTC(), tunnelLog, meta}

	rawValue, err := ioutil.ReadFile("/tmp/sauced-state.json")
	if err != nil {
		logger.Disklog.Warn("Could not read from statefile: ", err)
	}
	if rawValue == nil {
		tunnelsState.Tunnels = append(tunnelsState.Tunnels, tun)
	} else {
		json.Unmarshal(rawValue, &tunnelsState)
		tunnelsState.Tunnels = append(tunnelsState.Tunnels, tun)
	}

	tunnelStateJSON, err := json.Marshal(tunnelsState)
	if err != nil {
		logger.Disklog.Warn("Could not marshall the tunnel state data into JSON object: ", err)
	}
	err = ioutil.WriteFile("/tmp/sauced-state.json", tunnelStateJSON, 0755)
	if err != nil {
		logger.Disklog.Warn("Could not write the tunnel state data to the JSON file: ", err)
	}
}

// RemoveTunnel deletes a tunnel entry from the
// Last Known Tunnel object
func RemoveTunnel(targetPID int) {
	state := GetLastKnownState()
	for i := 0; i < len(state.Tunnels); i++ {
		tunnel := state.Tunnels[i]
		logger.Disklog.Infof("Closing Tunnel %s", tunnel.Args)
		if tunnel.PID == targetPID {
			state.Tunnels = append(state.Tunnels[:i], state.Tunnels[i+1:]...)
			break
		}
	}
	tunnelStateJSON, err := json.Marshal(state)
	if err != nil {
		logger.Disklog.Warn("Could not marshall the tunnel state data into JSON object: ", err)
	}
	err = ioutil.WriteFile("/tmp/sauced-state.json", tunnelStateJSON, 0755)
	if err != nil {
		logger.Disklog.Warn("Could not write the tunnel state data to the JSON file: ", err)
	}
}

// PruneState will access the state file and
// remove any entries that are not found by the OS
func PruneState() {
	state := GetLastKnownState()
	for i := 0; i < len(state.Tunnels); i++ {
		tunnel := state.Tunnels[i]
		proc, _ := os.FindProcess(tunnel.PID)
		syscallErr := proc.Signal(syscall.Signal(0))
		if syscallErr != nil {
			// catch OS permission issues.  Do not change state.Tunnels
			if syscallErr.Error() == "operation not permitted" {
				logger.Disklog.Warnf("Syscall Err on Signal(0) to tunnel PID %d: '%s'\nCheck that your current user has permissions to interact with this process. You may need 'sudo' or admin rights.", tunnel.PID, syscallErr)
				return
			}
			// take everything to the left of i, then conjoin it w/ everything to the right of i
			state.Tunnels = append(state.Tunnels[:i], state.Tunnels[i+1:]...)
			// decrement our for loop scope so we don't go out of bounds. do not remove.
			i--
			logger.Disklog.Infof("Found dead tunnel.  Removing from statefile: %v", tunnel)
			logger.Disklog.Warnf("Syscall Err on Signal(0) '%s' for PID %d", syscallErr, tunnel.PID)
		}
	}
	tunnelStateJSON, err := json.Marshal(state)
	if err != nil {
		logger.Disklog.Warn("Could not marshall the tunnel state data into JSON object: ", err)
	}
	err = ioutil.WriteFile("/tmp/sauced-state.json", tunnelStateJSON, 0755)
	if err != nil {
		logger.Disklog.Warn("Could not write the tunnel state data to the JSON file: ", err)
	}
}

// UpdateState uses the derived metadata
// to correct the statefile
func UpdateState(newMeta map[string]Metadata) {
	newState := GetLastKnownState()
	// loop through the lastknownstate and update the .metadata based on the newMeta.
	for index, tunnel := range newState.Tunnels {
		oldMeta := tunnel.Metadata
		if val, ok := newMeta[tunnel.Metadata.Pool]; ok {
			newState.Tunnels[index].Metadata = val
		} else {
			newState.Tunnels[index].Metadata.Size = 0
		}
		logger.Disklog.Debugf("Updated metadata %v with new metadata, %v", oldMeta, newMeta[tunnel.Metadata.Pool])
	}
	// commit to the statefile
	logger.Disklog.Debugf("Updated metadata marshalling to JSON and saving state: %v", newState)
	tunnelStateJSON, err := json.Marshal(newState)
	if err != nil {
		logger.Disklog.Warn("Could not marshall the tunnel state data into JSON object: ", err)
	}
	err = ioutil.WriteFile("/tmp/sauced-state.json", tunnelStateJSON, 0755)
	if err != nil {
		logger.Disklog.Warn("Could not write the tunnel state data to the JSON file: ", err)
	}
}

// GetLastKnownState will retrieve a byte[] slice with the contents
// of the /tmp/sauced-stat.json file
func GetLastKnownState() LastKnownTunnels {
	rawValue, err := ioutil.ReadFile("/tmp/sauced-state.json")
	if err != nil {
		logger.Disklog.Warn("Could not get last known state from file: ", err)
	}
	var tunnelsState LastKnownTunnels

	json.Unmarshal(rawValue, &tunnelsState)
	return tunnelsState
}

func createIPCFile() {
	if _, err := os.Stat("/tmp/sauced-state.json"); err == nil {
		logger.Disklog.Debug("Found state file /tmp/sauced-state.json")
	} else if os.IsNotExist(err) {
		logger.Disklog.Info("/tmp/sauced-state.json not found ", err)
		file, err := os.OpenFile("/tmp/sauced-state.json", os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			logger.Disklog.Warn("Failed to open /tmp/sauced-state.json")
		}
		defer file.Close()
	} else {
		logger.Disklog.Warn("unable to find /tmp/sauced-state.json")
	}
}
