package manager

import (
	"bufio"
	"os"
)

// Metadata is the collection of items that
// make a tunnel unique and part of a pool
type Metadata struct {
	Pool  string
	Size  int
	Owner string
}

// CollectMetadata parses the config file for important
// data points
func CollectMetadata(config string) map[string]Metadata {
	file, _ := os.Open(config)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		if fscanner.Text() != "" || len(fscanner.Text()) != 0 {
			go PoolName(fscanner.Text())
			go getUser(fscanner.Text())
			// use channels here to wait for the string data to return
			// silently return and fail if getOwner() fails

			// then append to the metadata map.  And increment the Size
		}
	}
}

// PoolName takes the launchArgs and parses for the
// tunnel name, if no name returns 'anonymous'
func PoolName(launchArgs string) string {
	//return the -i flag or anonymous if there is no name
}

func getOwner(launchArgs string) string {
	//return the user that owns this tunnel.
	//err and return if there is no user or -u flag.
}
