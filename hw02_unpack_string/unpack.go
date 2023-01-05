package hw02unpackstring

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
	if stringForUnpacking == "" {
		return "", nil
	}
	runes := []rune(stringForUnpacking)
	if !unicode.IsLetter(runes[0]) {
		return "", ErrStringStartsWithDigit
	}

	var result strings.Builder
	for i := 0; i < len(runes); i++ {
		currentChar := runes[i]
		nextChar := rune(0)
		if i+1 < len(runes) {
			nextChar = runes[i+1]
		}
		if unicode.IsLetter(currentChar) {
			if unicode.IsDigit(nextChar) {
				repeatCount := int(nextChar - '0')
				result.WriteString(strings.Repeat(string(currentChar), repeatCount))
			} else {
				result.WriteString(string(currentChar))
			}
		} else if unicode.IsDigit(currentChar) && unicode.IsDigit(nextChar) {
			return result.String(), ErrStringContainsNumbers
		}
	}

	return result.String(), nil
}
