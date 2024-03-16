package main

import "fmt"

func calculateFibonnaci(position int) []int {
	res := make([]int, 0)

	res = append(res, 0)
	if position == 1 {
		return res
	}

	res = append(res, 1)
	if position == 2 {
		return res
	}

	for i := 2; i < position; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

func recursiveFibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return recursiveFibonnaci(n-1) + recursiveFibonnaci(n-2)
}

func calculateFactorial(value int) int {
	if value <= 1 {
		return 1
	}

	res := 0
	for i := 2; i <= value; i++ {
		res *= i
	}

	return res
}

func recursiveFactorial(value int) int {
	if value <= 1 {
		return 1
	}

	return value * recursiveFactorial(value-1)
}

func main() {
	n := 10

	fmt.Println("Calculating fibonnaci")
	fmt.Println(calculateFibonnaci(n))

	fmt.Println("Printing recursive now...")
	for i := 0; i < n; i++ {
		fmt.Print(recursiveFibonnaci(i))
		fmt.Print(" ")
	}

	fmt.Println("")
	fmt.Println("Calculating factorial")
	fmt.Println(recursiveFactorial(5))

	fmt.Println("Calculating recursice factorial")
	fmt.Println(recursiveFactorial(5))
}
