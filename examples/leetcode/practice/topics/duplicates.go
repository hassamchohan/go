package main

import "fmt"

type Custom interface {
	int | string
}

func hasDuplicateStrings(strings []string) bool {
	m := make(map[string]int, 0)

	for _, str := range strings {
		if _, ok := m[str]; !ok {
			m[str] += 1
			continue
		}
		return true
	}
	return false
}

func hasDuplicateIntegers(numbers []int) bool {
	m := make(map[int]int, 0)

	for _, n := range numbers {
		if _, ok := m[n]; !ok {
			m[n] += 1
			continue
		}
		return true
	}
	return false
}

func hasDuplicates[T Custom](values []T) bool {
	m := make(map[T]int, 0)

	for _, n := range values {
		if _, ok := m[n]; !ok {
			m[n] += 1
			continue
		}
		return true
	}
	return false

}

func main() {
	s := []string{"1", "2", "3", "4", "5"}
	fmt.Println(hasDuplicateStrings(s))

	n := []int{1, 2, 3, 4, 5, 1, 1}
	fmt.Println(hasDuplicateIntegers(n))

	fmt.Println(hasDuplicates(s))
	fmt.Println(hasDuplicates(n))

}
