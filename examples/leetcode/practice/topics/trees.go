package main

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeTree struct {
	value int
	left  *NodeTree
	right *NodeTree
}

type BST struct {
	root   *NodeTree
	length int
}

func (n NodeTree) String() string {
	return strconv.Itoa(n.value)
}

func (b BST) String() string {
	sb := strings.Builder{}
	b.InOrderTraversal(&sb)
	return sb.String()

}

func (b *BST) InOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)

}

func (b *BST) inOrderTraversalByNode(sb *strings.Builder, root *NodeTree) {
	if root == nil {
		return
	}

	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *BST) Add(value int) {
	b.root = b.addByNode(b.root, value)
	b.length++
}

func (b *BST) addByNode(root *NodeTree, value int) *NodeTree {
	if root == nil {
		return &NodeTree{value: value}
	}

	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}

func main() {
	n1 := &NodeTree{1, nil, nil}
	n1.left = &NodeTree{0, nil, nil}
	n1.right = &NodeTree{2, nil, nil}

	b1 := BST{n1, 3}

	//fmt.Println(b1)

	b1.Add(5)
	b1.Add(4)
	b1.Add(6)

	//fmt.Println(b1)

	n2 := &NodeTree{1, nil, nil}
	n2.left = &NodeTree{0, nil, nil}
	n2.right = &NodeTree{2, nil, nil}

	b2 := BST{n2, 3}
	b2.Add(5)
	b2.Add(4)
	b2.Add(6)
	b2.Add(8)

	fmt.Println(b1.String())
	fmt.Println(b2.String())
	mirror := b1.String() == b2.String()
	if mirror {
		fmt.Printf("Both trees are same.")
	} else {
		fmt.Printf("Both trees are not same.")
	}

}
