package unpack_test

import (
	"errors"
	unpack "github.com/vokybykov/hw02_unpack_string"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := unpack.Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackStringStartsWithDigit(t *testing.T) {
	invalidStrings := []string{"3abc", "45"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := unpack.Unpack(tc)
			require.Truef(t, errors.Is(err, unpack.ErrStringStartsWithDigit), "actual error %q", err)
		})
	}
}
func TestUnpackStringContainsNumbers(t *testing.T) {
	invalidStrings := []string{"a23b", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := unpack.Unpack(tc)
			require.Truef(t, errors.Is(err, unpack.ErrStringContainsNumbers), "actual error %q", err)
		})
	}
}
