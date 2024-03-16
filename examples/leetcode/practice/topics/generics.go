package main

import "fmt"

type IntAndFloat interface {
	~float64 | int
}

type CustomFloat float64

func FindMax[T IntAndFloat](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var a CustomFloat = 2.3
	max := FindMax(a, 5.7)
	fmt.Println(max)
}
