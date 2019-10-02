package manager

import (
	"os"

	"github.com/mdsauce/sauced/logger"
)

func setDefaults(args []string) []string {
	args = addDefaults(args)

	logfileArg := ""
	logfileIndex := 0
	for index, arg := range args {
		if arg == "-l" || arg == "--logfile" {
			logfileIndex = index + 1
			logfileArg = args[index+1]
		}
	}

	args[logfileIndex] = setLogfile(logfileArg)

	return args
}

func addDefaults(tunnelArgs []string) []string {
	verbose, username, key, sePort, pidfile, haMode, logfile := false, false, false, false, false, false, false
	for _, arg := range tunnelArgs {
		switch arg {
		case "-v":
			verbose = true
		case "--verbose":
			verbose = true
		case "-vv":
			verbose = true
		case "-u":
			username = true
		case "--user":
			username = true
		case "-k":
			key = true
		case "--api-key":
			key = true
		case "--se-port":
			sePort = true
		case "-P":
			sePort = true
		case "--pidfile":
			pidfile = true
		case "-d":
			pidfile = true
		case "--no-remove-colliding-tunnels":
			haMode = true
		case "-l":
			logfile = true
		case "--logfile":
			logfile = true
		}
	}
	_, noFile := os.Stat("/tmp/sauceconnect")
	if noFile != nil {
		if os.IsNotExist(noFile) {
			err := os.Mkdir("/tmp/sauceconnect", 0744)
			if err != nil {
				logger.Disklog.Error("Could not make default directory /tmp/sauceconnect: ", err)
			}
		}
	} else if noFile == nil {
		logger.Disklog.Debug("Found /tmp/sauceconnect.  Adding Defaults to launch args")
	}

	if !verbose {
		tunnelArgs = append(tunnelArgs, "-v")
	}
	if !username {
		sauceuser := os.Getenv("SAUCE_USERNAME")
		tunnelArgs = append(tunnelArgs, "-u")
		tunnelArgs = append(tunnelArgs, sauceuser)
	}
	if !key {
		saucekey := os.Getenv("SAUCE_ACCESS_KEY")
		tunnelArgs = append(tunnelArgs, "-k")
		tunnelArgs = append(tunnelArgs, saucekey)
	}
	if !sePort {
		tunnelArgs = append(tunnelArgs, "--se-port")
		tunnelArgs = append(tunnelArgs, "0")
	}
	if !pidfile {
		tunnelArgs = append(tunnelArgs, "--pidfile")
		pidfileString := "/tmp/sauceconnect/" + randomString(5) + ".pid"
		tunnelArgs = append(tunnelArgs, pidfileString)
	}
	if !haMode {
		tunnelArgs = append(tunnelArgs, "--no-remove-colliding-tunnels")
	}
	if !logfile {
		tunnelArgs = append(tunnelArgs, "--logfile")
		rndLogfile := "/tmp/sauceconnect/" + randomString(5) + ".log"
		tunnelArgs = append(tunnelArgs, rndLogfile)
	}

	return tunnelArgs
}

func setLogfile(logfileArg string) string {
	defaultLogfile := ""

	// set a logfile name if there isn't one
	lastChars := logfileArg[len(logfileArg)-4:]
	if lastChars != ".log" {
		lastChar := logfileArg[len(logfileArg)-1]
		if lastChar == '/' {
			defaultLogfile = logfileArg + randomString(5) + ".log"
		} else {
			defaultLogfile = logfileArg + "/" + randomString(5) + ".log"
		}
	} else {
		return logfileArg
	}

	return defaultLogfile
}
