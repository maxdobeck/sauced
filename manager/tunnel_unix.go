// +build linux darwin

package manager

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/mdsauce/sauced/logger"
)

// Start creates a new tunnel from the metadata and launch arguments
func Start(launchArgs string, wg *sync.WaitGroup, meta Metadata) {
	defer wg.Done()
	args := strings.Split(launchArgs, " ")
	path := args[0]

	if eatLine(path) {
		return
	}
	if vacancy(meta) != true {
		logger.Disklog.Warnf("Too many tunnels open.  Not opening %s \n %v", meta.Pool, launchArgs)
		return
	}
	// setDefaults() should go here.  take launchArgs and add all necessary default args/flags.
	manufacturedArgs := setDefaults(args)
	logger.Disklog.Debug("Created new set of args with sensible defaults that will be passed to exec.Command: ", manufacturedArgs)
	// tunnel is actually launched here.  new process is spawned
	scCmd := exec.Command(path, manufacturedArgs[1:]...)
	stdout, _ := scCmd.StdoutPipe()
	err := scCmd.Start()
	if err != nil {
		logger.Disklog.Warnf("Something went wrong while starting the SC binary! %v", err)
		return
	}

	logger.Disklog.Infof("Tunnel started as process %d - %s\n", scCmd.Process.Pid, launchArgs)
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	// this parsing should be moved to its own funciton.
	// everything should be parsed then supplied to the AddTunnel() func
	var tunLog string
	var asgnID string
	for scanner.Scan() {
		m := scanner.Text()
		// should be a func that is unit tested
		if strings.Contains(m, "Log file:") {
			ll := strings.Split(m, " ")
			tunLog = ll[len(ll)-1]
			logger.Disklog.Infof("Tunnel log started for tunnel: %s \n %s", launchArgs, m)
		}
		// should be a func that is unit tested
		if strings.Contains(m, "Tunnel ID:") {
			idLine := strings.Split(m, " ")
			asgnID = idLine[len(idLine)-1]
			logger.Disklog.Infof("Tunnel: %s has AssignedID %s", launchArgs, asgnID)
		}
		if strings.Contains(m, "Sauce Connect is up") {
			AddTunnel(launchArgs, path, scCmd.Process.Pid, meta, tunLog, asgnID)
		}
	}
	logger.Disklog.Infof("Sauce Connect client with PID %d shutting down!  Goodbye!", scCmd.Process.Pid)
	RemoveTunnel(scCmd.Process.Pid)
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
	}
}

// StopTunnelByID will stop a single tunnel that matches a given ID
func StopTunnelByID(assignedID string) {
	tstate := GetLastKnownState()
	tunnel, err := tstate.FindTunnelByID(assignedID)

	if err != nil {
		logger.Disklog.Info(err)
	} else {
		Stop(tunnel.PID)
	}

}

// StopTunnelsByPool will stop a tunnel pool matching the given pool name
func StopTunnelsByPool(poolName string) {
	tstate := GetLastKnownState()
	tunnels, err := tstate.FindTunnelsByPool(poolName)

	if err != nil {
		logger.Disklog.Info(err)
	} else {
		for _, tunnel := range tunnels {
			Stop(tunnel.PID)
		}
	}
}

// StopAll will send a kill or SIGINT signal
// to all tunnels that are running.
func StopAll() {
	last := GetLastKnownState()
	for _, tunnel := range last.Tunnels {
		Stop(tunnel.PID)
		// add a stop via REST API func here.  So rest api gets signal as well.
	}
}
