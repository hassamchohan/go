package main

import (
	"fmt"
	"strconv"
)

type Node4 struct {
	Value int
	Left  *Node4
	Right *Node4
}

func (n *Node4) Add(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node4{Value: value}
		} else {
			n.Left.Add(value)
		}
	}

	if value > n.Value {
		if n.Right == nil {
			n.Right = &Node4{Value: value}
		} else {
			n.Right.Add(value)
		}
	}
}

func (n *Node4) Search(value int) bool {
	if n == nil {
		return false
	}
	if value < n.Value {
		return n.Left.Search(value)
	}

	if value > n.Value {
		return n.Right.Search(value)
	}

	return true
}

func (n *Node4) InOrderTraversal() {
	if n == nil {
		return
	}

	n.Left.InOrderTraversal()
	fmt.Print(fmt.Sprintf("%s ", strconv.Itoa(n.Value)))
	n.Right.InOrderTraversal()
}

func (n *Node4) PreOrderTraversal() {
	if n == nil {
		return
	}

	fmt.Print(fmt.Sprintf("%s ", strconv.Itoa(n.Value)))
	n.Left.InOrderTraversal()
	n.Right.InOrderTraversal()
}

func main() {
	tree := &Node4{Value: 10}
	tree.Add(5)
	tree.Add(3)
	tree.Add(7)
	tree.Add(15)
	tree.Add(12)
	tree.Add(20)

	fmt.Println(tree.Search(20))

	tree.InOrderTraversal()
	fmt.Println("")
	tree.PreOrderTraversal()

}
