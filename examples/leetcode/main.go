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

func main() {
	//fmt.Println(problems.ValidParenthesis("[][]]]"))

	numbers := []int{1, 2, 9, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 8, 9, 0}
	fmt.Println("Before: ", numbers)
	res := appendZeroesToEnd(numbers)
	fmt.Println("After: ", res)
}
