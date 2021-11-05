package model

type Btree struct {
	Root *Node

	degree int
}

type Node struct {
	leaf   bool
	keys   []int
	childs []*Node
}

func (n *Node) Len() int {
	return len(n.keys)
}

func BTreeSearch(x *Node, k int) (n *Node, idx int) {
	i := 1

	// linear search of minimal index i
	for i <= x.Len() && k > x.keys[i] {
		i = i + 1
	}

	if i <= x.Len() && k == x.keys[i] {
		// check do we have key in current node
		return x, i
	} else if x.leaf {
		// if it's a leaf - failed to find
		return nil, 0
	} else {
		// load more data from disk
		// DISK_READ(x.childs[i])
	}

	// recursively search in the sub tree
	return BTreeSearch(x.childs[i], k)
}

func AllocateNode(degree int) *Node {
	return &Node{
		keys:   make([]int, 2*degree-1),
		childs: make([]*Node, 2*degree),
	}
}

func BTreeCreate(degree int) *Btree {
	x := AllocateNode(degree)
	x.leaf = true
	//DISK_WRITE(x)

	return &Btree{Root: x, degree: degree}
}

// x - незаполненный внутренний узел
// x.c[i] - заполненный дочерний узел х
// splitChild - вырезать и вставить
func (t *Btree) splitChild(x *Node, i int) {
	y := x.childs[i]

	z := AllocateNode(t.degree) // todo: AllocateNode(t-1)
	z.leaf = y.leaf
	//z.n = t - 1

	for j := 1; j < t.degree-1; j++ {
		x.keys[j] = y.keys[j+t.degree]
	}

	if !y.leaf {
		for j := 1; j < t.degree; j++ {
			z.childs[j] = y.childs[j+t.degree]
		}
	}

	// y.n = t -1
	for j := x.Len() + 1; j > i+1; j-- {
		x.childs[j+1] = x.childs[j]
	}

	x.childs[i+1] = z

	for j := x.Len(); j > i; j-- {
		x.keys[j+1] = x.keys[j]
	}

	x.keys[i] = y.keys[t.degree]

	//x.n = x.n + 1
	//DISK_WRITE(y)
	//DISK_WRITE(z)
	//DISK_WRITE(x)
}

func (t *Btree) Insert(k int) {
	r := t.Root
	if r.Len() != 2*t.degree-1 {
		t.insertNonFull(r, k)
		return
	}

	s := AllocateNode(t.degree)
	t.Root = s
	s.leaf = false
	s.childs[1] = r

	t.splitChild(s, 1)

	t.insertNonFull(s, k)
}

func (t *Btree) insertNonFull(x *Node, k int) {
	i := x.Len()

	if x.leaf {
		for i >= 1 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i = i - 1
		}

		x.keys[i+1] = k
		//DISK_WRITE(x)
		return
	}

	for i >= 1 && k < x.keys[i] {
		i = i - 1
	}

	i = i + 1
	//DISK_READ(x.c[i])

	if x.childs[i].Len() == 2*t.degree-1 {
		t.splitChild(x, i)
		if k > x.keys[i] {
			i = i + 1
		}
	}

	t.insertNonFull(x.childs[i], k)
}
