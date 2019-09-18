package manager

import (
	"math/rand"
	"os"
	"strings"
)

func eatLine(line string) bool {
	// Comment symbol
	if line == "#" || line == "# " || strings.HasPrefix(line, "#") {
		return true
	}
	return false
}

func setDefaults(args []string) []string {
	args = addDefaults(args)

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

func missingRequiredFlags(args []string) bool {
	if len(args) <= 4 {
		return true
	}
	return false
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
