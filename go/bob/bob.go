package bob

import (
	"unicode"
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	
	switch {
	case IsEmpty(remark): return "Fine. Be that way!"
	case IsUpper(remark) && IsQuestion(remark): return "Calm down, I know what I'm doing!"
	case IsUpper(remark): return "Whoa, chill out!"
	case IsQuestion(remark): return "Sure." 
	default: return "Whatever."
	}
}

func IsEmpty(remark string) bool {
	return len(remark) == 0
}

func IsQuestion(remark string) bool {
	return remark[len(remark)-1:] == "?"
}

func IsUpper(remark string) bool {
	return ContainsLetter(remark) && strings.ToUpper(remark) == remark
}

func ContainsLetter(remark string) bool {
	for _, char := range remark {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}