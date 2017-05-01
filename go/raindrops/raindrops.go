package raindrops

import (
	"fmt"
)

const testVersion = 2

func Convert(num int) string {

	var out string

	if num%3 == 0 {
		out += "Pling"
	}

	if num%5 == 0 {
		out += "Plang"
	}

	if num%7 == 0 {
		out += "Plong"
	}

	if out == "" {
		return fmt.Sprintf("%d", num)
	}

	return out
}
