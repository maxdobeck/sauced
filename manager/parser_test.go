package manager

import (
	"fmt"
	"strings"
	"testing"
)

func TestEatLine(t *testing.T) {
	if eatLine("#") != true {
		t.Error("# should be a comment character and ignored.")
		t.Fail()
	}

	if eatLine("# ") != true {
		t.Fail()
	}

	if eatLine("# this is- = a 32! test") != true {
		t.Fail()
	}

	if eatLine("Hello!") != false {
		t.Error("This is NOT a comment and should be read by the program. Something is wrong if this failed.")
		t.Fail()
	}

	if eatLine("Test that this is NOT # a comment") != false {
		t.Error("Inline comments may be supported at a future date.  But that date is not today.")
		t.Fail()
	}
}

func TestMissingFlags(t *testing.T) {
	launchArgs := "/Users/maxdobeck/workspace/sauce_connect/sc-4.5.1-osx/bin/sc -i minimal-tunnel"
	args := strings.Split(launchArgs, " ")
	if missingRequiredFlags(args) != true {
		t.Error("missingFlags() should have caught this.  missingFlags returned false")
		t.Fail()
	}
}

func TestAddDefaults(t *testing.T) {
	launchArgs := "/Users/maxdobeck/workspace/sauce_connect/sc-4.5.1-osx/bin/sc -i minimal-tunnel"
	args := strings.Split(launchArgs, " ")
	verboseOutput := addDefaults(args)
	rawArgs := strings.Join(verboseOutput, " ")
	fmt.Println(verboseOutput)

	if strings.Contains(rawArgs, "-v") != true {
		t.Error("no -v found in addVerbosity() output: ", verboseOutput)
		t.Fail()
	}
	if strings.Contains(rawArgs, "--logfile") != true {
		t.Error("no --logfile found in output: ", verboseOutput)
		t.Fail()
	}
	if strings.Contains(rawArgs, "--no-remove-colliding-tunnels") != true {
		t.Error("no HA Mode flag found in output: ", verboseOutput)
		t.Fail()
	}
	if strings.Contains(rawArgs, "-u") != true || strings.Contains(rawArgs, "-u") != true {
		t.Error("no username and access key info found in output: ", verboseOutput)
		t.Fail()
	}
	if strings.Contains(rawArgs, "--pidfile") != true {
		t.Error("no pidfile info found in output: ", verboseOutput)
		t.Fail()
	}
	if strings.Contains(rawArgs, "--se-port") != true {
		t.Error("no se-port info found in output: ", verboseOutput)
		t.Fail()
	}
}
