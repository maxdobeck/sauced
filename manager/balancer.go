package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("Starting to collect metadata!")
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
			name := <-pool
			username := <-owner

			// then append to the metadata map.  And increment the Size
		}
	}
	return meta
}

// PoolName takes the launchArgs and parses for the
// tunnel name, if no name returns 'anonymous'
func PoolName(launchArgs string, pool chan string) {
	//return the -i flag or anonymous if there is no name
	args := strings.Split(launchArgs, " ")
	for index, arg := range args {
		if arg == "-i" {
			fmt.Println("Pool Name", args[index+1])
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
			fmt.Println("Owner Username", args[index+1])
			owner <- args[index+1]
		}
	}
	owner <- "none"
}
