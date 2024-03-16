package problems_tests

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestTopKFrequentWords(t *testing.T) {
	tests := []struct {
		in struct {
			words []string
			k     int
		}
		out []string
	}{
		{
			in: struct {
				words []string
				k     int
			}{
				words: []string{"hello", "hello", "world", "cup", "cup", "bat"},
				k:     2,
			},
			out: []string{"cup", "hello"},
		},
		{
			in: struct {
				words []string
				k     int
			}{
				words: []string{"in", "out", "in", "up", "up", "out", "in", "in", "in", "up", "up"},
				k:     2,
			},
			out: []string{"in", "up"},
		},
	}

	for _, tt := range tests {
		res := topKFrequentWords(tt.in.words, tt.in.k)
		require.Equal(t, tt.out, res)
	}
}

func topKFrequentWords(words []string, k int) []string {
	res := make([]string, 0)
	if len(words) == 0 {
		return res
	}

	//count the words
	m := make(map[string]int, 0)
	for _, w := range words {
		m[w] += 1
	}

	// convert the map to a slice of Pairs for sorting by values
	type Pair struct {
		Key   string
		Value int
	}

	pairs := make([]Pair, 0)
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	//sort the slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	//put top K words in result
	i := 0
	for k != 0 {
		res = append(res, pairs[i].Key)
		i++
		k--
	}

	// sort the result of top K words alphabetically

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
