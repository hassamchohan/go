package main

import (
	"fmt"
	"sort"
)

func isNumberPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n = n / 3
	}

	if n == 1 {
		return true
	}

	return false
}
func highestProduct(num []int) int {
	if len(num) < 3 {
		return 0
	}

	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})

	l := len(num)
	a := num[l-3] * num[l-2] * num[l-1]
	b := num[0] * num[1] * num[l-1]
	return max(a, b)
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

func maxSubArray(nums []int) int {
	currMax, total := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currMax = max(currMax+nums[i], nums[i])
		total = max(currMax, total)
	}

	return total
}

func getMaxProfit(stocks []int) int {
	maxProfit := stocks[1] - stocks[0]
	minPrice := stocks[0]

	for i := 1; i < len(stocks); i++ {
		currPrice := stocks[i]
		potentialProfit := currPrice - minPrice
		maxProfit = max(maxProfit, potentialProfit)
		minPrice = min(minPrice, currPrice)
	}

	return maxProfit
}

func getMaxSumOfSubArrayOfLengthK(nums []int, k int) int {
	maxSum := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		for j := i + 1; j < i+k && j < len(nums)-1; j++ {
			sum += nums[j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	//fmt.Println(isNumberPowerOfThree(10))

	_ = []int{0, -1, 10, 7, 5}
	//fmt.Println(highestProduct(n))

	_ = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//fmt.Println(maxSubArray(m))

	_ = []int{10, 7, 5, 8, 11, 9}
	//fmt.Println(getMaxProfit(s))

	a := []int{2, 4, 5, 1, 3, 4, 9, 9, 9, 1, 2, 4, 3}
	fmt.Println(getMaxSumOfSubArrayOfLengthK(a, 3))
}
