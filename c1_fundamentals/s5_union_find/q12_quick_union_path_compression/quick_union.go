package q12_quick_union_path_compression

// To create a path of length 4, union: (0,1), (1,2), (2,3).
// This produces the path 0 -> 1 -> 2 -> 3.
type QuickUnion struct {
	id    []int
	count int
}

func (qu *QuickUnion) Find(p int) int {
	// Find the root site.
	root := p
	for qu.id[root] != root {
		root = qu.id[root]
	}

	// Set the parent of each node on the path from p to root to be the root node.
	for qu.id[p] != p {
		next := qu.id[p]
		qu.id[p] = root
		p = next
	}

	return root
}

func (qu *QuickUnion) Union(p, q int) {
	pRoot, qRoot := qu.Find(p), qu.Find(q)
	if pRoot == qRoot {
		return
	}
	qu.id[pRoot] = qRoot
	qu.count--
}
