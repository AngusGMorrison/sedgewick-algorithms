package q14_weighted_quick_union_by_height

type QuickUnion struct {
	id     []int
	height []int
	count  int
}

func (qu *QuickUnion) Find(p int) int {
	for qu.id[p] != p {
		p = qu.id[p]
	}
	return p
}

func (qu *QuickUnion) Union(p, q int) {
	pRoot, qRoot := qu.Find(p), qu.Find(q)
	if qu.height[pRoot] < qu.height[qRoot] {
		qu.id[pRoot] = qRoot
		qu.height[pRoot] = qu.height[qRoot]
	} else if qu.height[qRoot] < qu.height[pRoot] {
		qu.id[qRoot] = pRoot
		qu.height[qRoot] = qu.height[pRoot]
	} else {
		// Heights are equal, so tree height must increase by 1.
		qu.id[pRoot] = qRoot
		qu.height[pRoot]++
		qu.height[qRoot]++
	}
}
