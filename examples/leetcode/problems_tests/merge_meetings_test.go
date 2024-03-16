package problems_tests

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

type Meeting struct {
	Start int
	End   int
}

func TestMergeMeetings(t *testing.T) {
	test := []struct {
		in  []Meeting
		out []Meeting
	}{
		{
			in:  []Meeting{},
			out: []Meeting{},
		},
		{
			in:  []Meeting{{1, 2}, {2, 3}, {4, 5}},
			out: []Meeting{{1, 3}, {4, 5}},
		},
		{
			in:  []Meeting{{1, 5}, {2, 3}},
			out: []Meeting{{1, 5}},
		},
	}

	for _, tt := range test {
		result := MergeMeetings(tt.in)
		require.Equal(t, tt.out, result)
	}
}

func GetMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func MergeMeetings(meetings []Meeting) []Meeting {
	// sort the meetings slice
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})

	result := []Meeting{}
	for i := range meetings {
		if i == 0 {
			result = append(result, meetings[i])
			continue
		}

		a := result[len(result)-1].End
		if a >= meetings[i].Start {
			result[len(result)-1].End = GetMax(a, meetings[i].End)
		} else {
			result = append(result, meetings[i])
		}
	}

	return result
}
