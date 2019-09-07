package manager

import (
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
