package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {

	runes := map[rune]struct{}{}

	for _, c := range strings.ToLower(word) {

		if !unicode.IsLetter(c) {
			continue
		}

		if _, ok := runes[c]; ok {
			return false
		}

		runes[c] = struct{}{}
	}

	return true
}
