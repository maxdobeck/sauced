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
			tunnelName := PoolName(fscanner.Text())
			username := GetOwner(fscanner.Text())
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
// tunnel name, if no name returns 'anonymous'
func PoolName(launchArgs string) string {
	//return the -i flag or anonymous if there is no name
	args := strings.Split(launchArgs, " ")
	for index, arg := range args {
		if arg == "-i" || arg == "--tunnel-identifier" {
			return args[index+1]
		}
	}
	return "anonymous"
}

// GetOwner takes the launch args and returns the owner of said tunnel if username flag is present
func GetOwner(launchArgs string) string {
	//return the user that owns this tunnel.
	//err and return if there is no user or -u flag.
	args := strings.Split(launchArgs, " ")
	for index, arg := range args {
		if arg == "-u" {
			return args[index+1]
		}
	}
	return "not found"
}
