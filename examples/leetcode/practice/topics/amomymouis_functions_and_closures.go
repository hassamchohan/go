package main

import (
	"fmt"
	"sort"
	"strconv"
)

type OpFunc func(int, int) int

var opMap = map[string]OpFunc{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func add(i int, j int) int { return i + j }

func subtract(i int, j int) int { return i - j }

func multiply(i int, j int) int { return i * j }

func divide(i int, j int) int { return i / j }

func intseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"15", "/", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression", exp)
			continue
		}

		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operation", op)
			continue
		}

		num1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		num2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		res := opFunc(num1, num2)
		fmt.Println(res)
	}

	// Closure Testing

	nextInt := intseq() // returns a closure meaning it returns a funciton that has a copy of i that can be reused and i will keep on changing

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4
	fmt.Println(nextInt()) // 5

	nextInt = intseq()     //reinitialized closure, now value of i will reinitilize to 0
	fmt.Println(nextInt()) // 1

	// Anonymous Function Testing

	people := []struct {
		Name string
		Age  int
	}{
		{
			Name: "Hassam",
			Age:  30,
		},
		{
			Name: "Ali",
			Age:  40,
		},
		{
			Name: "Peter",
			Age:  10,
		},
		{
			Name: "John",
			Age:  7,
		},
		{
			Name: "Andy",
			Age:  88,
		},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted Slice", people)
}
