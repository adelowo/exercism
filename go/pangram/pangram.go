package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {

	var chars = make(map[string]bool, 26)

	strings.Map(func(r rune) rune {

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			val := strings.ToLower(string(r))

			if _, ok := chars[val]; !ok {
				chars[val] = true
			}

			return r
		}

		return -1

	}, s)

	return len(chars) == 26
}
