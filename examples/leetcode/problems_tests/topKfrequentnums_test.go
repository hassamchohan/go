package problems_tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func sortByKeys(m map[string]int) {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, ": ", m[key])
	}
}

func sortByValues(m map[string]int) {
	type Pair struct {
		Key   string
		Value int
	}

	pairs := make([]Pair, len(m))

	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	fmt.Println(pairs)
}

func topKFrequent(nums []int, K int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}

	// count
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] += 1
	}

	// convert map to a slice of Pairs
	type Pair struct {
		Key   int
		Value int
	}
	pairs := make([]Pair, 0)

	i := 0
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	//sort slice of pairs based on the values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	i = 0
	for K != 0 {
		res = append(res, pairs[i].Key)
		i++
		K--
	}

	// sort the result slice
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func TestTopKFrequentNumbers(t *testing.T) {
	tests := []struct {
		in struct {
			nums []int
			k    int
		}
		out []int
	}{
		{
			in: struct {
				nums []int
				k    int
			}{
				nums: []int{9, 9, 9, 9, 9, 6, 5, 3, 5, 2, 6, 5, 7, 6, 5, 4, 5, 3, 5, 5, 5, 5, 1, 1, 1, 1, 1},
				k:    3,
			},
			out: []int{1, 5, 9},
		},
		{
			in: struct {
				nums []int
				k    int
			}{
				nums: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 6, 5, 3, 5, 2, 6, 5, 7, 6, 5, 4, 5, 3, 5, 5, 5, 5, 1, 1, 1, 1, 1},
				k:    2,
			},
			out: []int{5, 9},
		},
		{
			in: struct {
				nums []int
				k    int
			}{
				nums: []int{9, 9, 34, 4, 3, 6, 7, 3, 3, 6, 6, 8, 9, 5, 2, 6, 68, 9, 2, 5, 8, 23, 54, 4, 23, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 25, 6, 7, 8, 9, 2, 2, 2, 2, 2, 2, 2, 2, 2},
				k:    1,
			},
			out: []int{2},
		},
	}

	for _, tt := range tests {
		res := topKFrequent(tt.in.nums, tt.in.k)
		require.Equal(t, tt.out, res)
	}
}
