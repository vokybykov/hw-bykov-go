package unpack_test

import (
	unpack "github.com/vokybykov/hw02_unpack_string"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abccd", "abccd", nil},
		{"a2b3c4d", "aabbbccccd", nil},
		{"", "", nil},
		{"aaa0b", "aab", nil},
		{"45", "", unpack.ErrInvalidString},
		{"1a2b3c", "", unpack.ErrInvalidString},
	}
	for _, test := range tests {
		result, err := unpack.Unpack(test.input)
		if result != test.expected || err != test.err {
			t.Errorf(
				"unexpected result for input %s: got (%s, %v), expected (%s, %v)",
				test.input,
				result,
				err,
				test.expected,
				test.err)
		}
	}
}
