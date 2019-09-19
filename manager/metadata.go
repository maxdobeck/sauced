package manager

import (
	"bufio"
	"os"
	"strings"

	"github.com/mdsauce/sauced/logger"
)

// Metadata is the collection of items that
// make a tunnel unique and part of a pool
type Metadata struct {
	Pool  string
	Size  int
	Owner string
}

// CollectMetadata parses the config file for important
// data points and returns the formatted metadata
func CollectMetadata(config string) map[string]Metadata {
	logger.Disklog.Infof("Started collecting metadata from %s", config)
	meta := make(map[string]Metadata)
	file, _ := os.Open(config)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		if fscanner.Text() != "" || len(fscanner.Text()) != 0 {
			pool := make(chan string)
			owner := make(chan string)
			go PoolName(fscanner.Text(), pool)
			go getOwner(fscanner.Text(), owner)
			// use channels here to wait for the string data to return
			// silently return and fail if getOwner() fails
			tunnelName := <-pool
			username := <-owner

			// then add to the metadata map.  And increment the Size as needed
			if val, ok := meta[tunnelName]; ok {
				val.Size = val.Size + 1
				meta[tunnelName] = val
			} else {
				val = Metadata{Size: 1, Pool: tunnelName, Owner: username}
				meta[val.Pool] = val
			}
		}
	}
	logger.Disklog.Infof("Found metadata from %s: %v", config, meta)
	return meta
}

// PoolName takes the launchArgs and returns the
// tunnel name to the channel, if no name returns 'anonymous'
func PoolName(launchArgs string, pool chan string) {
	//return the -i flag or anonymous if there is no name
	args := strings.Split(launchArgs, " ")
	for index, arg := range args {
		if arg == "-i" || arg == "--tunnel-identifier" {
			pool <- args[index+1]
		}
	}
	pool <- "anonymous"
}

func getOwner(launchArgs string, owner chan string) {
	//return the user that owns this tunnel.
	//err and return if there is no user or -u flag.
	args := strings.Split(launchArgs, " ")
	for index, arg := range args {
		if arg == "-u" {
			owner <- args[index+1]
		}
	}
	owner <- "not found"
}
