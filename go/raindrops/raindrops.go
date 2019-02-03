package raindrops

import (
	"fmt"
	"strings"
)

const testVersion = 2

func Convert(num int) string {

	var s strings.Builder

	if num%3 == 0 {
		s.WriteString("Pling")
	}

	if num%5 == 0 {
		s.WriteString("Plang")
	}

	if num%7 == 0 {
		s.WriteString("Plong")
	}

	if s.Len() == 0 {
		return fmt.Sprintf("%d", num)
	}

	return s.String()
}
