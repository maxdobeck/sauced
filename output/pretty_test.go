package output

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/mdsauce/sauced/manager"
)

//TestPrettyStateFindsTunnel confirms that the target tunnel can be located
func TestPrettyStateFindsTunnel(t *testing.T) {
	target := "test123"
	state := soloState(target)

	var buf bytes.Buffer
	//log in this case is the log used by the Testing package
	log.SetOutput(&buf)

	showStatePretty(state)

	if strings.Contains(buf.String(), target) == false {
		t.Log("Current tunnels state: ", state)
		t.Log("Output from func under test: ", buf.String())
		t.Fail()
	}

}

//TestPrettyStatePrintsEmpty confirms that empty Tunnels[] can be identified as such
func TestPrettyStatePrintsEmpty(t *testing.T) {
	state := emptyState()

	var buf bytes.Buffer
	log.SetOutput(&buf)

	showStatePretty(state)

	if strings.Contains(buf.String(), "No tunnels are running right now!") == false {
		t.Log("Current tunnels state: ", state)
		t.Log("Output from func under test: ", buf.String())
		t.Fail()
	}
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

func emptyState() manager.LastKnownTunnels {
	var state manager.LastKnownTunnels
	return state
}
