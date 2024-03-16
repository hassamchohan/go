package main

import "fmt"

type LinkedList struct {
	value int
	next  *LinkedList
}

func (node *LinkedList) Insert(value int) {
	temp := &LinkedList{}
	temp.value = value

	if node == nil {
		node = temp
	} else {
		p := node
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}

func (node *LinkedList) Display() {
	p := node
	for p != nil {
		fmt.Printf("%d-> ", p.value)
		p = p.next
	}
	fmt.Println("\n")
}

func (node *LinkedList) DeleteFirstNode() {
	if node == nil {
		return
	}
	*node = *node.next
}

func (node *LinkedList) DeleteLastNode() {
	if node == nil {
		return
	}
	p := node
	for p.next.next != nil {
		p = p.next
	}
	p.next = nil
}

func (node *LinkedList) DeleteNode(value int) {
	if node == nil {
		return
	}
	if node.value == value {
		*node = *node.next
		return
	}

	curr := node
	for curr != nil && curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func main() {
	l := &LinkedList{
		value: 12,
		next: &LinkedList{
			value: 29,
			next: &LinkedList{
				value: 30,
				next:  nil,
			},
		},
	}
	l.Insert(50)
	l.Display()

	//l.DeleteLastNode()
	//l.Display()

	//l.DeleteLastNode()
	//l.Display()

	//l.Insert(500)
	//l.Insert(5000)
	//l.Display()

	l.DeleteNode(30)
	l.Display()

	l.DeleteNode(12)
	l.Display()

}
