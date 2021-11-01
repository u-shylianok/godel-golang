package model

const (
	RED = 0
	BLACK = 1 
)

type Node struct {
	Left *Node
	Right *Node
	Parent *Node 
	Color uint 

	Value int
}

type RBTree struct {
	Root *Node
}

func NewRBTree() *RBTree {
	root := &Node{Color: BLACK}
	return &RBTree{
		Root: root,
	}
}

func (t *RBTree) Insert(z *Node) *Node {
	var y *Node

	x := t.Root
	for x != nil {
		y = x
		if z.Value < x.Value {
			x = x.Left
		} else if x.Value < z.Value {
			x = x.Right
		} else {
			return x
		}
	}

	z.Parent = y
	if y == nil {
		t.Root = z
	} else if z.Value < y.Value {
		y.Left = z
	} else {
		y.Right = z
	}

	// t.count++
	t.insertFixup(z)
	return z
}

func (t *RBTree) insertFixup(z *Node) {
	for parent(z).Color == RED {
		// Case 1:
		if uncle(z).Color == RED {
			z.Parent.Color = BLACK
			y.Color = BLACK
			z.Parent.Parent.Color = RED
			z = z.Parent.Parent
			continue
		}

		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if z == z.Parent.Right {
				// Case 2:
				z = parent(z)
				t.rotLeft(z)
			}
			// Case 3:
			z.Parent.Color = BLACK
			z.Parent.Parent.Color = RED
			t.rotRight(grandfather(z))
		} else { 
			// same as above, but the opposite side
			y := grandfather(z).Left
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = grandfather(z)
			} else {
				if z == z.Parent.Left {
					z = parent(z)
					t.rotRight(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.rotLeft(grandfather(z))
			}
		}
	}
	t.Root.Color = BLACK
}

func (t *RBTree) rotRight(n *Node) {
	if n.Left == nil {
		return
	}

	y := n.Left
	n.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = n
	}
	y.Parent = n.Parent

	if n.Parent == nil {
		t.Root = y
	} else if n == n.Parent.Left {
		n.Parent.Left = y
	} else {
		n.Parent.Right = y
	}

	y.Right = n
	n.Parent = y
}

func (t *RBTree) rotLeft(n *Node) {
	if n.Right == nil {
		return
	}

	y := n.Right
	n.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = n
	}

	y.Parent = n.Parent

	if n.Parent == nil {
		t.Root = y
	} else if n == n.Parent.Left {
		n.Parent.Left = y
	} else {
		n.Parent.Right = y
	}

	y.Left = n
	n.Parent = y
}

func parent(n *Node) *Node {
	if n == nil {
		return leaf()
	}

	return n.Parent
} 

func grandfather(n *Node) *Node {
	p := parent(n)
	
	if p == nil {
		return leaf()
	}

	return p.Parent
}

func uncle(n *Node) *Node {
	p := parent(n)
	if p == nil {
		return leaf()
	}

	g := grandfather(n)
	if g == nil {
		return leaf()
	}

	if g.Left == p {
		return g.Right 
	}

	return g.Left
}

func leaf() *Node {
	return &Node{Color: BLACK}
}