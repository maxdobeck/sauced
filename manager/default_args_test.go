package manager

import (
	"testing"
)

func TestSetLogfileFromPath(t *testing.T) {
	userInput := "/some/random/path"
	newLogfile := setLogfile(userInput)
	lastChars := newLogfile[len(newLogfile)-4:]
	if lastChars != ".log" {
		t.Fail()
	}
}

func TestSetLogfileIgnoresGoodData(t *testing.T) {
	userInput := "/some/random/path/some.log"
	newLogfile := setLogfile(userInput)
	if newLogfile != userInput {
		t.Error(newLogfile, " != ", userInput)
		t.Fail()
	}
}
