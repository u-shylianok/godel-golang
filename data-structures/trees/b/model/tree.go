package model 

func BTreeSearch(x, k) {
	i := 1

	// linear search of minimal index i 
	while i <= x.n && k > x.key[i] {
		i = i + 1
	}

	
	if i <= x.n && k == x.key[i] {
		// check do we have key in current node
		return (x, i)
	} else if x.leaf {
		// if it's a leaf - failed to find
		return NIL 
	} else {
		// load more data from disk
		DISK_READ(x.c[i])
	}

	// recursively search in the sub tree
	return BTreeSearch(x.c[i], k)
}

func BTreeCreate(T) {
	x = AllocateNode() 
	x.leaf = true 
	x.n = 0 
	DISK_WRITE(x)
	T.root = x
}

// x - незаполненный внутренний узел 
// x.c[i] - заполненный дочерний узел х
func BTreeSplitChild(x, i) {
	z := AllocateNode()
	y := x.c[i]

	z.leaf = y.leaf
	z.n = t - 1 

	for j = 1 to t-1 {
		x.key[j] = y.key[j+t]
	}

	if !y.leaf {
		for j = 1 to t {
			z.c[j] = y.c[j+t]
		}
	}

	y.n = t -1 

	for j := x.n+1 downto i+1 {
		x.c[j+1] = x.c[j]
	}

	x.c[i+1] = z

	for j := x.n downto i {
		x.key[j+1] = x.key[j]
	}

	x.key[i] = y.key[t]
	x.n = x.n + 1

	DISK_WRITE(y)
	DISK_WRITE(z)
	DISK_WRITE(x)
} 