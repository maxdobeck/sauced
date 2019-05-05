// +build linux darwin

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
func Start(launchArgs string, wg *sync.WaitGroup, meta Metadata) {
	defer wg.Done()
	args := strings.Split(launchArgs, " ")
	path := args[0]

	scCmd := exec.Command(path, args[1:]...)
	stdout, _ := scCmd.StdoutPipe()
	err := scCmd.Start()
	if err != nil {
		logger.Disklog.Warnf("Something went wrong while starting the SC binary! %v", err)
		return
	}

	logger.Disklog.Infof("Tunnel started as process %d - %s\n", scCmd.Process.Pid, launchArgs)
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "Sauce Connect is up") {
			AddTunnel(launchArgs, path, scCmd.Process.Pid, meta)
		}
		if strings.Contains(m, "Log file:") {
			logger.Disklog.Infof("Tunnel log started for tunnel: %s \n %s", launchArgs, m)
		}
	}
	logger.Disklog.Infof("Tunnel %s closed", launchArgs)
	defer scCmd.Wait()
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
			logger.Disklog.Warnf("Problem killing Process %d %v", Pid, err)
		}
		RemoveTunnel(Pid)
	}
}

// StopAll will send a kill or SIGINT signal
// to all tunnels that are running.
func StopAll() {
	last := getLastKnownState()
	for _, tunnel := range last.Tunnels {
		Stop(tunnel.PID)
	}
}
