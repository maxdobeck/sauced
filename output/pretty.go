package output

import (
	"fmt"

	"github.com/mdsauce/sauced/manager"
)

// humanOutput takes the last known state for the tunnels and outputs it in a concise & readable way
func humanOutput() {
	state := manager.GetLastKnownState()
	fmt.Println(state)
}
