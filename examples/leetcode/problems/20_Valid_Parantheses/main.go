package problems

import "strings"

// First solution, BETTER ONE
func isValid1(s string) bool {

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune

	for _, v := range s {
		if _, ok := pairs[v]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Second solution
func isValid2(s string) bool {
	var stack []string
	for i := 0; i < len(s); i = i + 1 {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, string(s[i]))
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 {
				return false
			}
			index := len(stack) - 1
			item := stack[index]
			stack = stack[:index]
			if s[i] == ')' && item != string('(') || s[i] == ']' && item != string('[') || s[i] == '}' && item != string('{') {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func IsPalindrome(str string) bool {
	runes := []rune(str)
	mid := len(str) / 2
	for i, j := 0, len(str)-1; i <= mid && j >= mid; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func BinarySearch(numbers []int, value int) (int, bool) {
	min := 0
	max := len(numbers)
	mid := 0

	for min < max {
		mid = (min + max) / 2
		if value == numbers[mid] {
			return mid, true
		} else if value > numbers[mid] {
			min = mid + 1

		} else {
			max = mid - 1
		}
	}

	if min >= len(numbers) {
		return 0, false
	}

	return mid, true
}

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(value rune) {
	*s = append(*s, value)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return 0
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
func ValidParenthesis(str string) bool {
	stack := Stack{}
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
