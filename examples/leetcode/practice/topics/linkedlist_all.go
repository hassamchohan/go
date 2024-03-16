package main

import (
	"fmt"
	"sort"
)

type Node struct {
	value int
	Next  *Node
}

func (n *Node) Add(value int) {
	node := &Node{value: value}
	if n == nil {
		n = node
	}

	curr := n
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
}

func (n *Node) Print() {
	curr := n
	for curr != nil {
		fmt.Print(curr.value, " ")
		curr = curr.Next
	}
}

func (n *Node) Find(value int) bool {
	if n == nil {
		return false
	}

	curr := n
	for curr != nil {
		if value == curr.value {
			return true
		}
		curr = curr.Next
	}

	return false
}

func (n *Node) hasCycle() bool {
	m := make(map[*Node]int)

	curr := n
	for curr != nil {
		if _, ok := m[curr]; ok {
			return true
		}
		m[curr] = 0
		curr = curr.Next
	}

	return false
}

func (n *Node) ReverseLinkedList() *Node {
	if n == nil {
		return nil
	}

	if n.Next == nil {
		return n
	}

	var p *Node
	curr := n
	for curr != nil {
		n := curr.Next
		curr.Next = p
		p = curr
		curr = n
	}

	return p
}

func (n *Node) Remove(value int) {
	if n == nil {
		return
	}

	curr := n
	for curr.Next != nil {
		if curr.Next.value == value {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
}

func mergeTwoSortedLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &Node{}
	temp := dummy
	for list1 != nil && list2 != nil {
		if list1.value <= list2.value {
			temp.Next = list1
			temp = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}

	return dummy.Next
}

func mergeTwoSortedListsUsingArray(list1 *Node, list2 *Node) *Node {

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	arr := make([]int, 0)
	curr1 := list1
	for curr1 != nil {
		arr = append(arr, curr1.value)
		curr1 = curr1.Next
	}

	curr2 := list2
	for curr2 != nil {
		arr = append(arr, curr2.value)
		curr2 = curr2.Next
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	return arrayToLinkedList(arr)

}

func arrayToLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	temp := &Node{value: arr[0]}
	head := temp
	for i := 1; i < len(arr); i++ {
		n := &Node{
			value: arr[i],
		}
		temp.Next = n
		temp = temp.Next
	}

	return head
}

func main() {
	l := &Node{value: 1}
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	l.Print()

	fmt.Println(l.Find(5))

	lr := l.ReverseLinkedList()
	fmt.Println("Print Reverse Linked List")
	lr.Print()

	lr.Remove(3)
	fmt.Println("")
	lr.Print()

	lr.Remove(2)
	fmt.Println("")
	lr.Print()

	m1 := &Node{value: 1}
	m1.Add(2)
	m1.Add(3)
	m1.Add(4)

	m2 := &Node{value: 5}
	m2.Add(6)
	m2.Add(7)
	m2.Add(8)

	r := mergeTwoSortedListsUsingArray(m1, m2)
	//r := mergeTwoSortedLists(m1, m2)
	fmt.Println("")
	fmt.Println("Merged list")
	r.Print()

}
