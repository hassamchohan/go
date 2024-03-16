package problems_tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		in struct {
			a []int
			b []int
		}
		out []int
	}{
		{
			in: struct {
				a []int
				b []int
			}{
				a: []int{1, 2, 3},
				b: []int{4, 5, 6},
			},
			out: []int{1, 2, 3, 4, 5, 6},
		},
		{
			in: struct {
				a []int
				b []int
			}{
				a: []int{1, 3, 5},
				b: []int{2, 6, 7},
			},
			out: []int{1, 2, 3, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		res := mergeSortedArrays(tt.in.a, tt.in.b)
		require.Equal(t, tt.out, res)
	}

}

func mergeSortedArrays(a, b []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	if i < len(a) { // a is bigger than b
		res = append(res, a[i:]...)
	}

	if j < len(b) { // // b is bigger than a
		res = append(res, b[j:]...)
	}

	return res
}
