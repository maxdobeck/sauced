package manager

import "strings"

func eatLine(line string) bool {
	// Comment symbol
	if line == "#" || line == "# " || strings.HasPrefix(line, "#") {
		return true
	}
	return false
}
