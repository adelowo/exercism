package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(s string) bool {

	var stringBuilder strings.Builder

	for _, v := range s {
		isSpace := unicode.IsSpace(v)

		if !isSpace && !unicode.IsNumber(v) {
			return false
		}

		if !isSpace {
			stringBuilder.WriteRune(v)
		}
	}

	s = stringBuilder.String()

	if len(s) <= 1 {
		return false
	}

	var isSecond bool
	var total int

	for i := len(s) - 1; i >= 0; i-- {

		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}

		if !isSecond {
			total += n
			isSecond = true
			continue
		}

		if doubled := n * 2; doubled < 9 {
			total += doubled
		} else {
			total += (doubled - 9)
		}

		isSecond = false
	}

	return total%10 == 0
}
