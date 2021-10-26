package main 

import "fmt"

type Node struct {
	key int 
	left *Node 
	right *Node 
}

func (n *Node) insert(data int) {
	if data < n.key {
		// insert into the left tree
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
		return
	}
	// insert into the right tree
	if n.right == nil {
		n.right = &Node{key: data}
	} else {
		n.right.insert(data)
	}
} 

type Tree struct {
	root *Node
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
		return
	}
	t.root.insert(data)
}

func main() {
	var t Tree 

	t.insert(10)
	t.insert(1)
	t.insert(15)
	t.insert(16)
	t.insert(8)
	t.insert(10)
	t.insert(3)

	fmt.Printf("%+v", t)	
}
