package main

import "fmt"

// arr is sorted
func BinarySearch(arr []int, value int) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 && arr[0] == value {
		return true
	}
	min, mid := 0, 0
	max := len(arr) - 1
	for min < max {
		mid = (min + max) / 2
		if value == mid {
			return true
		} else if value > arr[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	if min >= len(arr) {
		return false
	}

	return true

}

func maxAverageofSubArray(arr []int, k int) int {
	if len(arr) <= 4 {
		return 0
	}

	maxSum := arr[0] + arr[1] + arr[2] + arr[3]
	if len(arr) == 4 {
		return maxSum
	}

	for i := 0; i <= len(arr)-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum

}

func main() {

	n := []int{1, 1, 1, 20000, 50, 3, 100, 200, 300, 800}
	//fmt.Println(BinarySearch(n, 200))
	fmt.Println(maxAverageofSubArray(n, 4))
}
