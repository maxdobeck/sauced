package output

import (
	"fmt"

	"github.com/mdsauce/sauced/manager"
)

// humanOutput takes the last known state for the tunnels and outputs it in a concise & readable way
func prettyOutputTunnel(id string) {
	state := manager.GetLastKnownState()
	if state.Empty() == true {
		fmt.Println("0 tunnels found by Sauced!")
	} else {
		fmt.Println(state)
	}
}

func prettyOutputPool(pool string) {
	state := manager.GetLastKnownState()
	if state.Empty() == true {
		fmt.Println("0 tunnels found by Sauced!")
	} else {
		fmt.Println(state)
	}
}

func prettyOutputState() {
	state := manager.GetLastKnownState()
	if state.Empty() == true {
		fmt.Println("0 tunnels found by Sauced!")
	} else {
		fmt.Println(state)
	}
}
