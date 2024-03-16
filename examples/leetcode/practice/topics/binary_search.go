package main

import (
	"fmt"
	"strconv"
)

func BinarySearch(numbers []int, value int) (int, bool) {
	min := 0
	max := len(numbers)
	mid := 0

	for min < max {
		mid = (min + max) / 2
		if value == numbers[mid] {
			return mid, true
		} else if value > numbers[mid] {
			min = mid + 1

		} else {
			max = mid - 1
		}
	}

	if min >= len(numbers) {
		return 0, false
	}

	return mid + 1, true
}

func main() {
	n := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	index, _ := BinarySearch(n, 200)
	fmt.Println("Value is present at index " + strconv.Itoa(index))
}
