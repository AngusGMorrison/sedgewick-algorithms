package unionfind

type WeightedQuickUnion struct {
	id    []int
	size  []int
	count int
}

func (wqu *WeightedQuickUnion) Find(p int) int {
	for wqu.id[p] != p {
		p = wqu.id[p]
	}
	return p
}

func (wqu *WeightedQuickUnion) Union(p, q int) {
	pRoot, qRoot := wqu.Find(p), wqu.Find(q)
	if pRoot == qRoot {
		return
	}

	if wqu.size[pRoot] < wqu.size[qRoot] {
		wqu.id[pRoot] = qRoot
	} else {
		wqu.id[qRoot] = pRoot
	}
	wqu.size[pRoot] += wqu.size[qRoot]
	wqu.size[qRoot] += wqu.size[pRoot]
	wqu.count--
}

func (wqu *WeightedQuickUnion) NewSite() int {
	site := len(wqu.id)
	wqu.id = append(wqu.id, site)
	wqu.size = append(wqu.size, 1)
	wqu.count++
	return site
}
