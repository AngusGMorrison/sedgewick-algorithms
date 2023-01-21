package ex_directed_bfs

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

type DirectedBFS struct {
	seen   []bool
	edgeTo []int
	source int
}

func New(g ex_digraph.Digraph, source int) *DirectedBFS {
	vertices := g.Vertices()
	d := &DirectedBFS{
		seen:   make([]bool, vertices),
		edgeTo: make([]int, vertices),
		source: source,
	}
	d.bfs(g, source)
	return d
}

func (d *DirectedBFS) bfs(g ex_digraph.Digraph, v int) {
	d.seen[v] = true
	worklist := queue.NewSliceQueue[int]()
	worklist.Enqueue(v)

	for v, ok := worklist.Dequeue(); ok; v, ok = worklist.Dequeue() {
		adj := g.Adjacent(v)
		for _, neighbor := range adj {
			if !d.seen[neighbor] {
				d.seen[neighbor] = true
				d.edgeTo[neighbor] = v
				worklist.Enqueue(neighbor)
			}
		}
	}
}

func (d *DirectedBFS) HasPathTo(v int) bool {
	return d.seen[v]
}

func (d *DirectedBFS) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return nil
	}

	var path []int
	for v != d.source {
		path = append(path, v)
		v = d.edgeTo[v]
	}
	path = append(path, v)
	reverse(path)
	return path
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
