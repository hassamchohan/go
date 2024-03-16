package main

import "fmt"

func mostOccuredCharInAstring(str string) string {
	runes := []rune(str)
	m := make(map[rune]int, 0)

	for _, r := range runes {
		m[r] += 1
	}

	max := 0
	var res rune = -1
	for k, v := range m {
		if v > max {
			max = v
			res = k
		}
	}

	return string(res)
}

func main() {

	str := "oooooppqwqwqwqwqwqqqqqqq"
	fmt.Println(mostOccuredCharInAstring(str))
}
