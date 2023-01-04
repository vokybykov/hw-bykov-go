package unpack

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrStringStartsWithDigit = errors.New("string starts with digit")
	ErrStringContainsNumbers = errors.New("string contains numbers")
)

func Unpack(stringForUnpacking string) (string, error) {
	var result strings.Builder

	if stringForUnpacking == "" {
		return result.String(), nil
	}

	if !unicode.IsLetter(rune(stringForUnpacking[0])) {
		return result.String(), ErrStringStartsWithDigit
	}

	for i := 0; i < len(stringForUnpacking); i++ {
		currentChar := stringForUnpacking[i]
		nextChar := rune(0)
		if i+1 < len(stringForUnpacking) {
			nextChar = rune(stringForUnpacking[i+1])
		}
		switch {
		case unicode.IsLetter(rune(currentChar)):
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
		case unicode.IsDigit(rune(currentChar)):
			if unicode.IsDigit(nextChar) {
				return result.String(), ErrStringContainsNumbers
			}
		}
	}

	return result.String(), nil
}
