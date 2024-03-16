package main

import (
	"fmt"
	"strings"
)

type Stack2 []rune

func (s *Stack2) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack2) Push(value rune) {
	*s = append(*s, value)
}

func (s *Stack2) Pop() rune {
	if s.IsEmpty() {
		return 0
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func validParenthesis(str string) bool {
	stack := Stack2{}
	brackets := ")}]({["
	runes := []rune(str)
	if runes[0] == '}' || runes[0] == ']' || runes[0] == ')' {
		return false
	}

	for _, s := range runes {
		if !strings.Contains(brackets, string(s)) {
			return false
		}

		if s == '{' || s == '[' || s == '(' {
			stack.Push(s)
			continue
		}

		c := stack.Pop()
		if (c == '(' && s != ')') || (c == '{' && s != '}') || (c == '[' && s != ']') || (c == 0) {
			return false
		}

	}
	return len(stack) == 0
}

func main() {
	//s := Stack2{}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)

	// fmt.Println(s.Pop()) // print 3
	// fmt.Println(s.Pop()) // print 2
	// fmt.Println(s.Pop()) // print 1

	fmt.Println(validParenthesis("(}"))

}
