package unionfind

// To produce a tree of depth 4, we must find only existing roots to avoid path compression, and
// join only trees of the same size, which increases the height of the resulting tree by 1.
// Required unions: (1,0), (7,6), (3,2), (5,4), (3,1), (7,5), (7,3)
type WeightedQuickUnion struct {
	id    []int
	size  []int
	count int
}

// Find returns the root of the component containing p.
func (wqu *WeightedQuickUnion) Find(p int) int {
	root := p
	for wqu.id[root] != root {
		root = wqu.id[root]
	}

	for wqu.id[p] != p {
		next := wqu.id[p]
		wqu.id[p] = root
		p = next
	}

	return root
}

func (wqu *WeightedQuickUnion) Union(p, q int) {
	pRoot, qRoot := wqu.Find(p), wqu.Find(q)
	if wqu.size[pRoot] < wqu.size[qRoot] {
		wqu.id[pRoot] = wqu.id[qRoot]
	} else {
		wqu.id[qRoot] = wqu.id[pRoot]
	}
	wqu.size[pRoot] += wqu.size[qRoot]
	wqu.size[qRoot] += wqu.size[pRoot]

	wqu.count--
}
