package problems_tests

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestLongestWordLength(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "my apple",
			out: 2,
		},
		{
			in:  "hello my name is hassam chohan",
			out: 6,
		},
		{
			in:  "ww iii oooo yuyuyu uiui jha yuuuuuuuuu io io uu y hjih jkh jkhk j",
			out: 10,
		},
	}

	for _, tt := range tests {
		res := longestEvenWordLength(tt.in)
		require.Equal(t, tt.out, res)
	}
}

func longestEvenWordLength(str string) int {
	words := strings.Split(str, " ")

	maxLength := 0
	for _, w := range words {
		if len(w)%2 == 0 && len(w) > maxLength {
			maxLength = len(w)
		}
	}

	return maxLength
}
