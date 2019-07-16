package cmd

import (
	"testing"
)

// TestSCStarts will start a sc tunnel if the $SAUCE_USERNAME and $SAUCE_ACCESS_KEY environment variables are set with valid data
func TestBadConfigs(t *testing.T) {
	if configUsable("") != false {
		t.Fail()
	}

	if configUsable(" ") != false {
		t.Fail()
	}

	if configUsable("bad file path") != false {
		t.Fail()
	}

	if configUsable("/path/to/non-existent/file.txt") != false {
		t.Fail()
	}
}
