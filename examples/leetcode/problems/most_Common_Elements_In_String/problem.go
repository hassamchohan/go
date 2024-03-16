package most_Common_Elements_In_String

import "sort"

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func mostCommonElementsInAString(str string) map[string]int {
	result := map[string]int{}
	var sortedMap PairList
	stringMap := map[string]int{}

	for _, v := range str {
		if _, ok := stringMap[string(v)]; ok {
			stringMap[string(v)]++
			continue
		}
		stringMap[string(v)] = 1
	}

	for k, v := range stringMap {
		sortedMap = append(sortedMap, struct {
			Key   string
			Value int
		}{Key: k, Value: v})
	}

	sort.Slice(sortedMap, func(i, j int) bool {
		return sortedMap[i].Value > sortedMap[j].Value
	})

	i := 0
	for _, v := range sortedMap {
		if i >= 2 {
			break
		}
		result[v.Key] = v.Value
		i++
	}

	return result

}

func rotate(nums []int, k int) []int {

	for i := 0; i < k; i++ {
		l := nums[len(nums)-1]
		nums = append([]int{l}, nums[:len(nums)-1]...)
	}
	return nums

}

func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		l := nums[len(nums)-1]
		nums = append([]int{l}, nums[:len(nums)-1]...)
	}
}
