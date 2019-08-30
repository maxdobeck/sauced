package output

import (
	"fmt"

	"github.com/mdsauce/sauced/manager"
)

func showTunnelPretty(ID string, state manager.LastKnownTunnels) {
	if state.Empty() == true {
		noTunnels()
		return
	}

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
