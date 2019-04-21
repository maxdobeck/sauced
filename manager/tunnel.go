package manager

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/maxdobeck/sauced/logger"
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
		logger.Disklog.Fatal("Something went wrong while starting the SC binary! ", err)
	}

	logger.Disklog.Infof("Sauce Connect started as process %d.\n", scCmd.Process.Pid)
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "Sauce Connect is up") {
			logger.Disklog.Infof("Sauce Connect started! These arguments launched with success: %s Killing it for you now so you don't forget!", launchArgs)
			// can't send interrupts on Windows!! Beware, must use scCmd.Process.Kill
			Stop(scCmd.Process.Pid)
		}
		if strings.Contains(m, "Log file:") {
			logger.Disklog.Infof("Tunnel log started for tunnel: %s \n %s", launchArgs, m)
		}
	}
}

// Stop will halt a running process with SIGINT(CTRL-C)
func Stop(Pid int) {
	tunnel, err := os.FindProcess(Pid)
	if err != nil {
		logger.Disklog.Warnf("Process ID %d does not exist or could not be sent the SIGINT.", Pid)
	} else {
		time.Sleep(6 * time.Second)
		err := tunnel.Signal(os.Interrupt)
		if err != nil {
			logger.Disklog.Warnf("Problem killing Process", Pid, err)
		}
	}
}
