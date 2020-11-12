package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(message string) bool {
	var chars = make(map[rune]int)
	for _, char := range strings.ToUpper(message) {
		if unicode.IsLetter(char) && chars[char] > 0 {
			return false
		}
		chars[char]++
	}

	return true
}
