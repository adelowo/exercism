package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	counter := make(Frequency)

	for _, word := range splitWords(phrase) {
		if len(strings.Trim(word, "")) == 0 {
			continue
		}

		counter[strings.ToLower(strings.Trim(word, `'`))]++
	}

	return counter
}

func splitWords(s string) []string {
	var words []string
	var word string

	for _, v := range s + " " {
		if unicode.IsLetter(v) || unicode.IsDigit(v) || string(v) == `'` {
			word = word + string(v)
		} else {
			words = append(words, word)
			word = ""
		}
	}

	return words
}
