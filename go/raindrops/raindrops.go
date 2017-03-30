package raindrops

import (
	"strconv"
)

const testVersion = 2

type raindrops map[int]string

var mappings raindrops

func init() {
	mappings = make(raindrops, 3)
	mappings[3] = "Pling"
	mappings[5] = "Plang"
	mappings[7] = "Plong"
}

func Convert(n int) string {

	if n == 1 {
		return strconv.Itoa(n)
	}

	var response string

	allFactors := factorsFor(n)

	for _, v := range allFactors {
		if isDigitConvertableToString(v) {
			response += mappings[v]
		}
	}

	if response == "" {
		return strconv.Itoa(n)
	}

	return response

}

func factorsFor(n int) []int {
	var f []int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			f = append(f, i)
		}
	}

	return f
}

func isDigitConvertableToString(n int) bool {
	_, exists := mappings[n]
	return exists
}
