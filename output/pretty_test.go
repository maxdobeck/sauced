package output

import (
	"testing"
	"time"

	"github.com/mdsauce/sauced/manager"
)

//TestPrettyState tests the
func TestPrettyStateFindsTunnel(t *testing.T) {
	target := "test123"
	state := soloState(target)
	showStatePretty(state)
}

func TestPrettyStatePrintsEmpty(t *testing.T) {
	state := emptystate()
	showStatePretty(state)
}

func soloState(target string) manager.LastKnownTunnels {
	meta := manager.Metadata{Size: 1, Pool: target, Owner: "me.user"}

	var tunnel1 manager.Tunnel
	tunnel1.PID = 12345
	tunnel1.AssignedID = "ab1lk5b1glah9s"
	tunnel1.SCBinary = "some/path/to/sc.exe"
	tunnel1.Args = "-v -u me.user -k some-secret-key"
	tunnel1.LaunchTime = time.Now().UTC()
	tunnel1.Log = "path/to/tunnels/log.log"
	tunnel1.Metadata = meta

	var state manager.LastKnownTunnels
	state.Tunnels = append(state.Tunnels, tunnel1)

	return state
}

func emptystate() manager.LastKnownTunnels {
	var state manager.LastKnownTunnels
	return state
}
