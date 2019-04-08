package manager

import (
	"fmt"
)

// Tunnel represents the core values of a tunnel
// and any of its flag values.  Created from parsing the yaml config file
type Tunnel struct {
	Identifier string
	User       string
	Key        string
	NoBump     []string
	Verbose    bool
}

// ParseConfigs to get an array that constitutes 1 or more tunnels
func ParseConfigs(tunnels map[string]interface{}) {
	var target []Tunnel
	fmt.Println(tunnels)
	for key, tunnel := range tunnels {
		fmt.Println("Tunnel: ", key)
		if arg, ok := tunnel.(map[string]interface{}); ok {
			var temp Tunnel
			for key, val := range arg {
				fmt.Println(key, val)
				switch key {
				case "user":
					temp.User = val.(string)
				case "verbose":
					temp.Verbose = val.(bool)
				case "key":
					temp.Key = val.(string)
				case "nobump":
					switch val.(type) {
					case string:
						noBumpList := make([]string, 1)
						noBumpList[0] = val.(string)
						temp.NoBump = noBumpList
					case []string:
						temp.NoBump = val.([]string)
					}
				}
			}
			temp.Identifier = key
			target = append(target, temp)
			fmt.Println()
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", tunnel)
		}
	}
	fmt.Println()
	fmt.Println(target)
}