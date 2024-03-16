package main

import (
	"fmt"
	"strings"
)

func reverseSentence(str string) string {
	if str == "" {
		return ""
	}
	words := strings.Split(str, " ")
	str = ""
	for i := len(words) - 1; i >= 0; i-- {
		str += words[i]
		str += " "
	}
	str = str[:len(str)-1]
	return str
}

func reverseString(str *string) {
	if *str == "" {
		return
	}
	runes := []rune(*str)
	mid := len(*str) / 2
	for i, j := 0, len(*str)-1; i <= mid && j >= mid; i, j = i+1, j-1 {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp
	}

	*str = string(runes)
}

func isPalindrome(str string) bool {
	if str == "" {
		return false
	}
	runes := []rune(str)
	mid := len(str) / 2
	for i, j := 0, len(str)-1; i <= mid && j >= mid; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func countSubStrings(str, subString string) int {
	if str == "" || subString == "" {
		return 0
	}

	return strings.Count(str, subString)
}

func main() {
	//str := "dlrow olleh"
	//reverseString(&str)
	//fmt.Println(str)
	//
	//str2 := "hassam"
	//fmt.Println(isPalindrome(str2))
	//
	//str3 := "hello world hello man hello boy hello girl hello panda"
	//fmt.Println(countSubStrings(str3, "hello"))

	sen := "Hello! How are you?"
	s := reverseSentence(sen)
	fmt.Println(s)
}
