package manager

import (
	"encoding/json"
	"github.com/maxdobeck/sauced/logger"
	"io/ioutil"
	"os"
	"time"
)

// LastKnownTunnels will a json object with
// all tunnels that were previously known to be alive
type LastKnownTunnels struct {
	Tunnels []TunnelState `json:"tunnels"`
}

// TunnelState is a json representation
// of the a tunnel as an OS process after it has launched
type TunnelState struct {
	PID        int       `json:"pid"`
	SCBinary   string    `json:"scbinary"`
	Args       string    `json:"args"`
	LaunchTime time.Time `json:"launchtime"`
}

// AddTunnel will record the state of the tunnel
// to the IPC file after the tunnel has launched as an OS process
func AddTunnel(launchArgs string, path string, PID int) {
	fp := getIPCFile()
	defer fp.Close()
	var tunnels LastKnownTunnels

	tun := TunnelState{PID, path, launchArgs, time.Now().UTC()}

	rawValue, _ := ioutil.ReadAll(fp)
	if rawValue == nil {
		// start a new TunnelState list, marshall the json, and write it to file
	} else {
		json.Unmarshal(rawValue, &tunnels)
		tunnels.Tunnels = append(tunnels.Tunnels, tun)
	}
}

func getIPCFile() *os.File {
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
