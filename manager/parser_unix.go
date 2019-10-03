package manager

import (
	"strings"
)

func eatLine(line string) bool {
	// Comment symbol
	if line == "#" || line == "# " || strings.HasPrefix(line, "#") {
		return true
	}
	return false
}

func missingRequiredFlags(args []string) bool {
	if len(args) <= 4 {
		return true
	}
	return false
}
