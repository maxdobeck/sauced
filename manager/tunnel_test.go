package manager

import (
	"os"
	"sync"
	"testing"
	"time"
)

// TestSCStarts will start a sc tunnel if the $SAUCE_USERNAME and $SAUCE_ACCESS_KEY environment variables are set with valid data
func TestSCStarts(t *testing.T) {
	// alternatively download sauce connect
	if _, err := os.Stat("/home/max/workspace/tools/sc-4.5.1-linux/bin/sc"); err != nil {
		t.SkipNow()
	}
	scBinary := "/home/max/workspace/tools/sc-4.5.1-linux/bin/sc"
	var wg sync.WaitGroup
	wg.Add(1)
	go Start(scBinary, &wg, Metadata{})
	time.Sleep(5 * time.Second)
	StopAll()
}

// TestSCFailsOnBadInput asserts that only an actual argument to SC can start a tunnel
func TestSCFailsOnBadInput(t *testing.T) {
	badPath := "/some/fake/path"
	var wg sync.WaitGroup
	wg.Add(1)
	go Start(badPath, &wg, Metadata{})
	PruneState()
	tunnels := GetLastKnownState()
	if tunnels.Empty() == false {
		t.Log(tunnels.Tunnels)
		t.Error("Tunnels started even with scBinary path: ", badPath)
	}
	StopAll()
}
