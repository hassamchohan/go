package problems_tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMaxProfit(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{10, 7, 5, 8, 11},
			out: 6,
		},
		{
			in:  []int{10, 2, 4, 6, 9, 4, 5, 18, 19, 20},
			out: 18,
		},
	}

	for _, tt := range tests {
		res := getMaxProfit(tt.in)
		require.Equal(t, tt.out, res)
	}
}
func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
func getMaxProfit(stocks []int) int {
	if len(stocks) <= 2 {
		return 0
	}

	maxProfit := stocks[1] - stocks[0]
	minPrice := stocks[0]

	for i := 1; i < len(stocks); i++ {
		curr := stocks[i]
		profit := curr - minPrice
		maxProfit = max(profit, maxProfit)
		minPrice = min(minPrice, curr)
	}

	return maxProfit
}
