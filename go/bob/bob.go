package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(greeting string) string {
	greeting = strings.TrimSpace(greeting)

	if greeting == "" {
		return "Fine. Be that way!"
	}

	if isBobYelledAt(greeting) {
		return "Whoa, chill out!"
	}

	if greeting[len(greeting)-1] == '?' {
		return "Sure."
	}

	return "Whatever."

}

func isBobYelledAt(s string) bool {
	var upper, lower bool

	for _, v := range s {
		if unicode.IsUpper(v) {
			upper = true
			break
		}
	}

	for _, v := range s {
		if unicode.IsLower(v) {
			lower = true
			break
		}
	}

	return upper && !lower
}
