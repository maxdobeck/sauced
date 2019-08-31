package output

import (
	"log"

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
	tunnels, err := state.FindTunnelsByPool(pool)
	if err != nil {
		logger.Disklog.Warn("Problem searching for tunnels for specific pool", " , ", pool, err)
	}
	for _, tunnel := range tunnels {
		printTunnel(tunnel)
	}
}

func showStatePretty(state manager.LastKnownTunnels) {
	if state.Empty() == true {
		noTunnels()
		return
	}
	for _, tunnel := range state.Tunnels {
		printTunnel(tunnel)
	}
}

func noTunnels() {
	log.Printf("\nNo tunnels are running right now!\n\n")
	log.Println("Tunnels:")
	log.Println("--------")
	log.Print("None\n\n")
}

func printTunnel(t manager.Tunnel) {
	log.Print("Tunnel Found")
	log.Println("------------")
	log.Println("PID: ", t.PID)
	log.Println("Assigned ID: ", t.AssignedID)
	log.Println("Pool Name: ", t.Metadata.Pool)
	log.Println("Tunnel Log Location: ", t.Log)
	//convert to local OS timezone and humanize
	// t.LaunchTime
	log.Println("Owner: ", t.Metadata.Owner)
	log.Print("Pool Size: ", t.Metadata.Size, "\n\n")
}
