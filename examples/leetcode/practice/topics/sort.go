package main

import (
	"fmt"
	"sort"
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

	return res
}

func main() {
	// m := map[string]int{
	// 	"Alice":  1,
	// 	"Camara": 3,
	// 	"Bob":    2,
	// 	"Zebra":  10,
	// 	"Lion":   19,
	// 	"Monkey": 22,
	// 	"Orange": 20,
	// }

	nums := []int{1, 1, 1, 2, 2, 3}
	r := topKFrequent(nums, 2)
	fmt.Println(r)

	//sortByKeys(m)
	//sortByValues(m)
}
