package manager

import (
	"sync"
	"testing"
)

// TestSCStarts will start a sc tunnel if the $SAUCE_USERNAME and $SAUCE_ACCESS_KEY environment variables are set with valid data
func TestSCStarts(t *testing.T) {
	// pass in a valid path here for local testing
	scBinary := "/home/max/workspace/tools/sc-4.5.1-linux/bin/sc"
	var wg sync.WaitGroup
	wg.Add(1)
	go Start(scBinary, &wg)
	wg.Wait()
}
