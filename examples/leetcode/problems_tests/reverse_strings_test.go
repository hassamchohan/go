package problems_tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseStrings(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		{
			in:  []string{},
			out: []string{},
		},
		{
			in:  []string{"a", "b", "c", "d"},
			out: []string{"d", "c", "b", "a"},
		},
		{
			in:  []string{"aa", "bb"},
			out: []string{"bb", "aa"},
		},
		{
			in:  []string{"ee", "ww", "rr", "oo", "ll"},
			out: []string{"ll", "oo", "rr", "ww", "ee"},
		},
	}

	for _, tt := range tests {
		res := ReverseStrings(tt.in)
		require.Equal(t, tt.out, res)
	}
}

func TestReverseStrings2(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "hassam",
			out: "massah",
		},
		{
			in:  "popopopo",
			out: "opopopop",
		},
		{
			in:  "abcdefghijkl",
			out: "lkjihgfedcba",
		},
	}

	for _, tt := range tests {
		res := ReverseStrings2(tt.in)
		require.Equal(t, tt.out, res)
	}
}

func Swap(a, b *string) {
	*a, *b = *b, *a
}

func ReverseStrings(str []string) []string {
	for i, j := 0, len(str)-1; i <= len(str)/2 && j >= len(str)/2; i, j = i+1, j-1 {
		Swap(&str[i], &str[j])
	}
	return str
}

func ReverseStrings2(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
