package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func (ll *ListNode) add(node *ListNode) {
	if ll == nil {
		return
	}

	curr := ll
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
}

func (ll *ListNode) printLinkedList() {
	if ll == nil {
		return
	}

	curr := ll
	for curr != nil {
		fmt.Print(strconv.Itoa(curr.Value) + "->")
		curr = curr.Next
	}
	fmt.Println("")
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{}
	}

	curr := head
	var prev *ListNode
	for curr != nil {
		n := curr.Next
		curr.Next = prev
		prev = curr
		curr = n
	}

	return prev
}

func mergeTwoSortedLinkedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	for list1 != nil && list2 != nil {
		if list1.Value <= list2.Value {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	} else {
		curr.Next = list1
	}

	return result.Next

}

func hasCycle(list *ListNode) bool {
	if list == nil {
		return false
	}

	m := make(map[*ListNode]int, 0)
	curr := list
	for curr != nil {
		if _, ok := m[curr]; ok {
			return true
		}
		m[curr]++
		curr = curr.Next
	}

	return false
}

func main() {

	l := &ListNode{
		Value: 1,
		Next:  nil,
	}
	l.add(&ListNode{
		Value: 2,
		Next:  nil,
	})
	l.add(&ListNode{
		Value: 3,
		Next:  nil,
	})
	l.add(&ListNode{
		Value: 4,
		Next:  nil,
	})

	l.add(&ListNode{
		Value: 5,
		Next:  nil,
	})

	//l.printLinkedList()
	//l = reverseList(l)
	//l.printLinkedList()

	i := 42
	str := strconv.Itoa(i) //Int to String
	fmt.Println(str)

	str = "55"
	n, _ := strconv.Atoi(str) //String to int
	fmt.Println(n)

	i = 42
	s := float64(i) //Int to Float
	fmt.Println(s)

	A := &ListNode{
		Value: 4,
		Next:  nil,
	}
	A.add(&ListNode{
		Value: 8,
		Next:  nil,
	})
	A.add(&ListNode{
		Value: 9,
		Next:  nil,
	})

	B := &ListNode{
		Value: 4,
		Next:  nil,
	}
	B.add(&ListNode{
		Value: 5,
		Next:  nil,
	})
	B.add(&ListNode{
		Value: 6,
		Next:  nil,
	})

	//merged := mergeTwoSortedLinkedLists(A, B)
	//merged.printLinkedList()

	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	f := &ListNode{6, nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

	cycle := hasCycle(a)
	fmt.Println(cycle)

}
