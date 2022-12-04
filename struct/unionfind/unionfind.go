package unionfind

type QuickFind struct {
	id    []int
	count int
}

// NewQuickFind initializes a QuickFind object with id[i] = i.
func NewQuickFind(size int) *QuickFind {
	qf := QuickFind{
		id:    make([]int, size),
		count: size,
	}
	for i := range qf.id {
		qf.id[i] = i
	}
	return &qf
}

// Find returns the ID of the site at id[p].
// O(1) time.
func (qf *QuickFind) Find(p int) int {
	return qf.id[p]
}

// Union joins the two components containing p and q by setting each site in the array that shares
// the same parent as p to have the parent of q. This choice is arbitrary - the parent of p could
// replace each site that has the parent of q.
// O(n) time.
func (qf *QuickFind) Union(p, q int) {
	pParent := qf.Find(p)
	qParent := qf.Find(q)
	if pParent == qParent {
		return
	}

	for i := range qf.id {
		if qf.id[i] == pParent {
			qf.id[i] = qParent
		}
	}
	qf.count--
}

// Connected returns true if sites p and q are part of the same component.
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
}

func (qf *QuickFind) Count() int {
	return qf.count
}

type QuickUnion struct {
	id    []int
	count int
}

// NewQuickUnion initializes a QuickUnion object with id[i] = i.
func NewQuickUnion(size int) *QuickUnion {
	qf := QuickUnion{
		id:    make([]int, size),
		count: size,
	}
	for i := range qf.id {
		qf.id[i] = i
	}
	return &qf
}

// Find returns the root of the component containing site p.
func (qu *QuickUnion) Find(p int) int {
	for qu.id[p] != p {
		p = qu.id[p]
	}
	return p
}

func (qu *QuickUnion) Union(p, q int) {
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	if pRoot == qRoot {
		return
	}

	qu.id[pRoot] = qRoot
	qu.count--
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *QuickUnion) Count() int {
	return qu.count
}

type WeightedQuickUnion struct {
	id    []int // forest of trees
	size  []int
	count int
}

func NewWeightedQuickUnion(size int) *WeightedQuickUnion {
	wqu := WeightedQuickUnion{
		id:    make([]int, size),
		size:  make([]int, size),
		count: size,
	}
	for i := range wqu.id {
		wqu.id[i] = i
		wqu.size[i] = 1
	}
	return &wqu
}

// Find returns the root of the component containing site p. The maximum depth of any node is lg n.
func (wqu *WeightedQuickUnion) Find(p int) int {
	for wqu.id[p] != p {
		p = wqu.id[p]
	}
	return p
}

// Union joins the components of p and q into a single component by connecting the smaller component
// to the root of the larger component, thus minimizing the rate at which the height of the tree
// increases. The worst-case order of growth is lg n.
func (wqu *WeightedQuickUnion) Union(p, q int) {
	pRoot := wqu.Find(p)
	qRoot := wqu.Find(q)
	if pRoot == qRoot { // already connected
		return
	}

	if wqu.size[pRoot] < wqu.size[qRoot] { // connect the smaller tree to the root of the larger tree
		wqu.id[pRoot] = qRoot
	} else {
		wqu.id[qRoot] = pRoot
	}
	wqu.size[pRoot] += wqu.size[qRoot]
	wqu.size[qRoot] += wqu.size[pRoot]
	wqu.count--
}

func (wqu *WeightedQuickUnion) Connected(p, q int) bool {
	return wqu.Find(p) == wqu.Find(q)
}

func (wqu *WeightedQuickUnion) Count() int {
	return wqu.count
}
