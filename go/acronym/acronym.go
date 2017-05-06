package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Abbreviate(s string) string {
	//Bail out quickly
	if strings.Contains(s, ":") {
		return strings.Split(s, ":")[0]
	}

	var acr string

	for k, v := range s {
		if k == 0 {
			acr += strings.ToUpper(string(v))
			continue
		}

		previous := rune(s[k-1])
		current := rune(s[k])

		isLetter := unicode.IsLetter(current) && !unicode.IsLetter(previous)

		isUpper := unicode.IsUpper(current) && !unicode.IsUpper(previous)

		if isLetter || isUpper {
			acr += strings.ToUpper(string(v))
		}
	}

	return acr
}
