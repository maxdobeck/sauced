package manager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Start creates a new tunnel
func Start(binary string) error {
	fmt.Println("Launching Sauce Connect Proxy binary at", binary)

	scCmd := exec.Command(binary)
	stdout, _ := scCmd.StdoutPipe()
	err := scCmd.Start()
	if err != nil {
		fmt.Println("Something went wrong with the sc binary! ", err)
		return err
	}

	fmt.Printf("Sauce Connect started as process %d.\n", scCmd.Process.Pid)
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		if strings.Contains(m, "Sauce Connect is up") {
			fmt.Println("Sauce Connect started!  Killing it for you now so you don't forget!")
			// can't send interrupts on Windows!! Beware, must use scCmd.Process.Kill
			scCmd.Process.Signal(os.Interrupt)
			break
		}
	}
	return nil
}

// ReadConfigs uses Viper to get a map of strings that constitute 1 or more tunnels
func ReadConfigs(tunnels map[string]string) {
	for key, tunnel := range tunnels {
		fmt.Println()
		fmt.Println("Tunnel: ", key)
		fmt.Println(tunnel)
	}
	fmt.Println()
}
