package manager

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/mdsauce/sauced/logger"
)

// LastKnownTunnels is a slice of all tunnels found in the statefile
type LastKnownTunnels struct {
	Tunnels []Tunnel `json:"tunnels"`
}

// Tunnel is a logical representation of a Sauce Connect tunnel once it has started on an OS
type Tunnel struct {
	PID        int       `json:"pid"`
	AssignedID string    `json:"assignedID"`
	SCBinary   string    `json:"scbinary"`
	Args       string    `json:"args"`
	LaunchTime time.Time `json:"launchtime"`
	Log        string    `json:"log"`
	Metadata   Metadata  `json:"metadata"`
}

// Empty returns true if there are no tunnels, false if one ore more tunnels are present
func (tState LastKnownTunnels) Empty() bool {
	if len(tState.Tunnels) == 0 {
		return true
	}
	return false
}

// add later
// func (tun Tunnel) rotateLog() error {
// 	return nil
// }

// FindTunnelByPID gets a PID int and finds a given tunnel
func (tState LastKnownTunnels) FindTunnelByPID(targetPID int) (Tunnel, error) {
	var tunnel Tunnel
	for i := 0; i < len(tState.Tunnels); i++ {
		tunnel = tState.Tunnels[i]
		if tunnel.PID == targetPID {
			return tunnel, nil

		}
	}
	return tunnel, errors.New("No tunnel found with given PID")
}

// FindTunnelsByPool gets a pool name and returns a list of tunnels
func (tState LastKnownTunnels) FindTunnelsByPool(poolName string) ([]Tunnel, error) {
	var tunnel Tunnel
	tunnels := make([]Tunnel, 0)
	for i := 0; i < len(tState.Tunnels); i++ {
		tunnel = tState.Tunnels[i]
		if tunnel.Metadata.Pool == poolName {
			tunnels = append(tunnels, tunnel)

		}
	}
	if len(tunnels) == 0 {

		return tunnels, errors.New("No tunnel found with given Pool name")
	}
	return tunnels, nil
}

// FindTunnelByID gets a tunnel ID and returns a tunnel
func (tState LastKnownTunnels) FindTunnelByID(assignedID string) (Tunnel, error) {
	var tunnel Tunnel
	for i := 0; i < len(tState.Tunnels); i++ {
		tunnel = tState.Tunnels[i]
		if tunnel.AssignedID == assignedID {
			return tunnel, nil

		}
	}

	return tunnel, errors.New("No tunnel found with given ID")

}

// AddTunnel will record the state of the tunnel
// to the IPC file after the tunnel has launched as an OS process
func AddTunnel(launchArgs string, path string, PID int, meta Metadata, tunnelLog string, asgnID string) {
	createIPCFile(PID)
	// TODO Obtain a lock on the file THEN add to the statefile
	var state LastKnownTunnels
	tun := Tunnel{PID, asgnID, path, launchArgs, time.Now().In(time.Local), tunnelLog, meta}

	rawState, err := ioutil.ReadFile("/tmp/sauced-state.json")
	if err != nil {
		logger.Disklog.Warn("Could not read from statefile: ", err)
	}

	// TODO can this be condensed so that we do the Unmarshall if rawState != nil ?
	if rawState == nil {
		state.Tunnels = append(state.Tunnels, tun)
	} else {
		json.Unmarshal(rawState, &state)
		state.Tunnels = append(state.Tunnels, tun)
	}

	saveState(state)
}

// RemoveTunnel deletes a tunnel entry from the
// Last Known Tunnel object
func RemoveTunnel(targetPID int) {
	state := GetLastKnownState()
	for i := 0; i < len(state.Tunnels); i++ {
		tunnel := state.Tunnels[i]
		if tunnel.PID == targetPID {
			state.Tunnels = append(state.Tunnels[:i], state.Tunnels[i+1:]...)
			break
		}
	}
	saveState(state)
}

// PruneState will access the state file and
// remove any entries that are not found by the OS
func PruneState() {
	logger.Disklog.Debug("Pruning tunnels")
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
	saveState(state)
}

// GetLastKnownState will retrieve a byte[] slice with the contents of the sauced-state.json file
func GetLastKnownState() LastKnownTunnels {
	rawValue, err := ioutil.ReadFile("/tmp/sauced-state.json")
	if err != nil {
		logger.Disklog.Warn("Could not get last known state from file: ", err)
	}
	var state LastKnownTunnels

	json.Unmarshal(rawValue, &state)
	return state
}

func createIPCFile(PID int) {
	if _, err := os.Stat("/tmp/sauced-state.json"); err == nil {
		logger.Disklog.Debugf("Process %d found state file /tmp/sauced-state.json", PID)
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

func saveState(newState LastKnownTunnels) {
	tunnelStateJSON, err := json.Marshal(newState)
	if err != nil {
		logger.Disklog.Warn("Could not marshall the tunnel state data into JSON object: ", err)
	}
	err = ioutil.WriteFile("/tmp/sauced-state.json", tunnelStateJSON, 0755)
	if err != nil {
		logger.Disklog.Warn("Could not write the tunnel state data to the JSON file: ", err)
	}
}
