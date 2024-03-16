package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) AddNode(value int) {
	if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.AddNode(value)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.AddNode(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	}
	if value > n.Value {
		return n.Right.Search(value)
	} else {
		return n.Left.Search(value)
	}
}

func (n *Node) PreOrder() {
	if n != nil {
		fmt.Print(n.Value, " ")
		n.Left.PreOrder()
		n.Right.PreOrder()
	}
}

func (n *Node) InOrder() {
	if n != nil {
		n.Left.InOrder()
		fmt.Print(n.Value, " ")
		n.Right.InOrder()
	}
}

func (n *Node) PostOrder() {
	if n != nil {
		n.Left.PostOrder()
		n.Right.PostOrder()
		fmt.Print(n.Value, " ")
	}
}

func main() {
	n := &Node{Value: 10}
	n.AddNode(15)
	n.AddNode(6)
	n.AddNode(3)
	n.AddNode(5)
	n.AddNode(30)
	n.AddNode(2)
	n.AddNode(9)
	n.AddNode(8)

	n.PreOrder()
	fmt.Println()
	n.InOrder()
	fmt.Println()
	n.PostOrder()
	fmt.Println()
	fmt.Println(n.Search(8))

}
