package ex_bfs_paths

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

type BreadthFirstPaths struct {
	seen   []bool
	edgeTo []int
	source int
}

func New(g ex_graph.Graph, source int) *BreadthFirstPaths {
	vertices := g.Vertices()
	bfp := BreadthFirstPaths{
		seen:   make([]bool, vertices),
		edgeTo: make([]int, vertices),
		source: source,
	}
	bfp.bfs(g, source)
	return &bfp
}

func (bfp *BreadthFirstPaths) bfs(g ex_graph.Graph, source int) {
	q := queue.NewSliceQueue[int]()
	bfp.seen[source] = true
	q.Enqueue(source)

	for v, ok := q.Dequeue(); ok; v, ok = q.Dequeue() {
		for _, w := range g.Adjacent(v) {
			if bfp.seen[w] {
				continue
			}

			bfp.seen[w] = true
			bfp.edgeTo[w] = v
			q.Enqueue(w)
		}
	}
}

func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	return bfp.seen[v]
}

func (bfp *BreadthFirstPaths) PathTo(v int) []int {
	if !bfp.HasPathTo(v) {
		return nil
	}

	var path []int
	for x := v; x != bfp.source; x = bfp.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, bfp.source)
	reverse(path)
	return path
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i-1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
