package manager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// Start creates a new tunnel
func Start(launchArgs string, wg *sync.WaitGroup) {
	defer wg.Done()
	args := strings.Split(launchArgs, " ")
	path := args[0]

	scCmd := exec.Command(path, args[1:]...)
	stdout, _ := scCmd.StdoutPipe()
	err := scCmd.Start()
	if err != nil {
		fmt.Println("Something went wrong while starting the SC binary! ", err)
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
			Stop(scCmd.Process.Pid)
		}
	}
}

// Stop will halt a running process with SIGINT(CTRL-C)
func Stop(Pid int) {
	tunnel, err := os.FindProcess(Pid)
	if err != nil {
		fmt.Printf("Process ID %d does not exist or could not be sent the SIGINT.", Pid)
	} else {
		time.Sleep(6 * time.Second)
		err := tunnel.Signal(os.Interrupt)
		if err != nil {
			fmt.Println("Problem killing Process", Pid, err)
		}
	}
}
