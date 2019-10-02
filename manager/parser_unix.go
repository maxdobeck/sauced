package manager

import (
	"math/rand"
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

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
