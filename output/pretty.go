package output

import (
	"fmt"

	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
)

func showTunnelPretty(ID string, state manager.LastKnownTunnels) {
	if state.Empty() == true {
		noTunnels()
		return
	}
	tunnel, err := state.FindTunnelsByID(ID)
	if err != nil {
		logger.Disklog.Warn("Problem searching Statefile for tunnel", " , ", ID, err)
	}
	printTunnel(tunnel)
}

func showPoolPretty(pool string, state manager.LastKnownTunnels) {
	if state.Empty() == true {
		noTunnels()
		return
	}
}

func showStatePretty(state manager.LastKnownTunnels) {
	if state.Empty() == true {
		noTunnels()
		return
	}
}

func noTunnels() {
	fmt.Printf("\nNo tunnels are running right now!\n\n")
	fmt.Println("Tunnels:")
	fmt.Println("--------")
	fmt.Print("None\n\n")
}

func printTunnel(t manager.Tunnel) {
	fmt.Println("\nTunnel Found")
	fmt.Println("------------")
	fmt.Println("PID: ", t.PID)
	fmt.Println("Assigned ID: ", t.AssignedID)
	fmt.Println("Name: ", t.Metadata.Pool)
	fmt.Println("Tunnel Log Location: ", t.Log)
	//convert to local OS timezone and humanize
	// t.LaunchTime
	fmt.Println("Owner: ", t.Metadata.Owner)
	fmt.Println("Pool Size: ", t.Metadata.Size)
	fmt.Println()
}
