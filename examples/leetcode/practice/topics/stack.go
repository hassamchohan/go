package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(element []string) {
	*s = append(*s, element...)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true

}

func (s *Stack) Print() {
	fmt.Println(*s)
}

func main() {
	s := &Stack{}
	s.Push([]string{"hello", "hello1", "hello2", "hello3", "hello4"})
	s.Print()

	_, _ = s.Pop()
	s.Print()

	_, _ = s.Pop()
	s.Print()

	_, _ = s.Pop()
	s.Print()

	_, _ = s.Pop()
	s.Print()

	_, _ = s.Pop()
	s.Print()
}
