package main

import "fmt"

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j, j+1)
			}
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(arr) // Before Sort
	BubbleSort(arr)
	fmt.Println(arr) // After Sort
}
