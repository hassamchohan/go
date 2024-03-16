package main

import "fmt"

type NODE struct {
	next  *NODE
	value int
}

type LINKEDLIST struct {
	head   *NODE
	length int
}

func (l *LINKEDLIST) Prepend(value int) {
	node := &NODE{nil, value}
	if l.head == nil {
		l.head = node
		l.length++
		return
	}

	node.next = l.head
	l.head = node
	l.length++
}

func (l *LINKEDLIST) Delete(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	curr := l.head
	for curr != nil && curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			l.length--
		}
		curr = curr.next
	}
}

func (l *LINKEDLIST) Print() {
	if l.length == 0 {
		return
	}

	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	l := LINKEDLIST{}
	l.Prepend(88)
	l.Prepend(78)
	l.Prepend(68)
	l.Prepend(58)
	l.Prepend(48)

	l.Delete(88)

	l.Print()

}
