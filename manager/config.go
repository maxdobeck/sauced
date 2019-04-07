package manager

import (
	"fmt"
)

type tunnel struct {
	Identifier string
	User       string
	Key        string
	NoBump     []string
	Verbose    bool
}

// ParseConfigs uses Viper to get a map of strings that constitute 1 or more tunnels
func ParseConfigs(tunnels map[string]interface{}) {
	fmt.Println(tunnels)
	for key, tunnel := range tunnels {
		fmt.Println("Tunnel: ", key)
		if arg, ok := tunnel.(map[string]interface{}); ok {
			for key, val := range arg {
				fmt.Println(key, val)
			}
			fmt.Println()
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", tunnel)
		}
	}
	fmt.Println()
}
