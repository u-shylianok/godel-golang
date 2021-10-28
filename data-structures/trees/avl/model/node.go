package model 

import "fmt"

type Node struct {
	Key int 
	Height int
	Left *Node 
	Right *Node 
}

func NewNode(key int) *Node {
	return &Node{
		Key: key,
		Height: 1,
	}
}

func insertNode(n *Node, key int) *Node {
	if n == nil {
		return NewNode(key)
	}

	if n.Key > key {
		n.Left = insertNode(n.Left, key)
	} else if n.Key < key {
		n.Right = insertNode(n.Right, key)
	}

	n.Height = 1 + getNodeHeight(n.Left) + getNodeHeight(n.Right);

	return balance(n)
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%d", n.Key)
}

func (n *Node) fixHeight() {
	fmt.Printf("\nfixing height for node [%s]\n", n)

	hl := getNodeHeight(n.Left)
	hr := getNodeHeight(n.Right)

	if hl > hr {
		n.Height = hl + 1
	} else {
		n.Height = hr + 1
	}

	fmt.Printf("\n  node height is %d now\n", n.Height)
}
