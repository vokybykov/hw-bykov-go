package unpack

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(stringForUnpacking string) (string, error) {
	var result strings.Builder

	if stringForUnpacking == "" {
		return result.String(), nil
	}

	if !unicode.IsLetter(rune(stringForUnpacking[0])) {
		return result.String(), ErrInvalidString
	}

	for i := 0; i < len(stringForUnpacking); i++ {
		currentChar := stringForUnpacking[i]
		nextChar := rune(0)
		if i+1 < len(stringForUnpacking) {
			nextChar = rune(stringForUnpacking[i+1])
		}
		if unicode.IsLetter(rune(currentChar)) {
			if unicode.IsDigit(nextChar) {
				if nextChar > 0 {
					repeatCount := int(nextChar - '0')
					result.WriteString(strings.Repeat(string(currentChar), repeatCount))
				} else {
					continue
				}
			} else {
				result.WriteString(string(currentChar))
			}
		} else if unicode.IsDigit(rune(currentChar)) {
			if unicode.IsDigit(nextChar) {
				return result.String(), ErrInvalidString
			}
		}
	}

	return result.String(), nil
}
