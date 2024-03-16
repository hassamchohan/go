package main

import (
	"fmt"
)

func appendZeroesToEnd(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	nz := make([]int, 0)
	zc := 0
	for _, n := range numbers {
		if n != 0 {
			nz = append(nz, n)
			continue
		}
		zc++
	}

	for i := 0; i < zc; i++ {
		nz = append(nz, 0)
	}

	return nz
}

func isAnagram(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}

	if len(str1) != len(str2) {
		return false
	}

	runes1 := []rune(str1)
	m1 := make(map[rune]int, 0)

	for _, r := range runes1 {
		if _, ok := m1[r]; !ok {
			m1[r]++
		}
	}

	runes2 := []rune(str2)
	m2 := make(map[rune]int, 0)

	for _, r := range runes2 {
		if _, ok := m2[r]; !ok {
			m2[r]++
		}
	}

	for r, c := range m1 {
		if _, ok := m2[r]; !ok {
			return false
		}
		if m2[r] != c {
			return false
		}
	}

	return true
}

func main() {
	numbers := []int{1, 2, 9, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 8, 9, 0}
	fmt.Println("Before: ", numbers)
	res := appendZeroesToEnd(numbers)
	fmt.Println("After: ", res)

	fmt.Println(isAnagram("anagram", "aangmar"))
}
