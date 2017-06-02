package pangram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsPangram(s string) bool {

	var chars = make(map[string]struct{}, 26)

	strings.Map(func(r rune) rune {

		if upper := unicode.ToUpper(r); upper >= 'A' && upper <= unicode.MaxASCII {
			val := strings.ToLower(string(r))

			if _, ok := chars[val]; !ok {
				chars[val] = struct{}{}
			}

			return r
		}

		return -1

	}, s)

	return len(chars) == 26
}
