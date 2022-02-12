package lib

import (
	"unicode"
)

func IsHiragana(str string) bool {
	for _, r := range []rune(str) {
		if unicode.In(r, unicode.Hiragana) {
			return true
		}
	}
	return false
}
