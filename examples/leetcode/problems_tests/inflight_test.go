package problems_tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFillMoviesForFlight(t *testing.T) {
	tests := []struct {
		in struct {
			movies       []int
			flightLength int
		}
		out bool
	}{
		{
			in: struct {
				movies       []int
				flightLength int
			}{
				movies:       []int{},
				flightLength: 0,
			},
			out: false,
		},
		{
			in: struct {
				movies       []int
				flightLength int
			}{
				movies:       []int{2, 3, 4},
				flightLength: 6,
			},
			out: true,
		},
		{
			in: struct {
				movies       []int
				flightLength int
			}{
				movies:       []int{5, 2, 19, 9, 0, 1, 4, 8, 2, 4},
				flightLength: 8,
			},
			out: true,
		},
	}

	for _, tt := range tests {
		res := fillMoviesForFlight(tt.in.movies, tt.in.flightLength)
		require.Equal(t, tt.out, res)
	}
}

func fillMoviesForFlight(movies []int, flightLength int) bool {
	if len(movies) == 0 {
		return false
	}

	m := make(map[int]int, 0)
	for _, mov := range movies {
		if _, ok := m[mov]; !ok {
			m[mov] = 1
		}

		n := flightLength - mov
		if _, ok := m[n]; ok {
			return true
		}
	}

	return false
}
