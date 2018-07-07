package twofer

import (
	"fmt"
	"strings"
)

// ShareWith should have a comment documenting it.
func ShareWith(replacement string) string {
	if len(strings.TrimSpace(replacement)) == 0 {
		replacement = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", replacement)
}
