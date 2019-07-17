// +build linux darwin

package manager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/mdsauce/sauced/logger"
)

// Start creates a new tunnel
func Start(launchArgs string, wg *sync.WaitGroup, meta Metadata) {
	defer wg.Done()
	args := strings.Split(launchArgs, " ")
	path := args[0]

	if path == "#" || strings.Contains(path, "#") {
		fmt.Println("path we're using here>>>>>>>>>>>>", path)
		return
	}

	if vacancy(meta) != true {
		logger.Disklog.Infof("Too many tunnels open.  Not opening %s \n %v", meta.Pool, launchArgs)
		return
	}
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

	var tunLog string
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "Log file:") {
			ll := strings.Split(m, " ")
			tunLog = ll[len(ll)-1]
			logger.Disklog.Infof("Tunnel log started for tunnel: %s \n %s", launchArgs, m)
		}
		if strings.Contains(m, "Sauce Connect is up") {
			AddTunnel(launchArgs, path, scCmd.Process.Pid, meta, tunLog)
		}
	}
	logger.Disklog.Debugf("Tunnel %s closed", launchArgs)
	defer scCmd.Wait()
}

// Stop will halt a running process with SIGINT(CTRL-C)
func Stop(Pid int) {
	tunnel, err := os.FindProcess(Pid)
	if err != nil {
		logger.Disklog.Warnf("Process ID %d does not exist or was not accessible for this user. Error: %v", Pid, err)
	} else {
		err := tunnel.Signal(os.Interrupt)
		if err != nil {
			logger.Disklog.Warnf("Problem killing Process %d %v.  The user may not have permissions to send a SIGINT or SIGKILL to the listed process.", Pid, err)
		}
		// Only amend statefile if there wasn't an error
		if err == nil {
			RemoveTunnel(Pid)
		}
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
