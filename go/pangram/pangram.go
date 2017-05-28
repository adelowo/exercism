package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {

	var chars = make(map[string]bool, 26)

	strings.Map(func(r rune) rune {

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {

			if _, ok := chars[strings.ToLower(string(r))]; !ok {
				chars[strings.ToLower(string(r))] = true
			}

			return r
		}

		return -1

	}, s)

	if len(chars) == 26 {
		return true
	}

	return false

}
