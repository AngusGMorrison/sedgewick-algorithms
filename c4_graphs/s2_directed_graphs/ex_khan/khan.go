package ex_khan

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

func TopoSort(g ex_digraph.Digraph) []int {
	vertices := g.Vertices()
	indeg := indegree(g)
	order := make([]int, 0, vertices)

	worklist := queue.NewSliceQueue[int]()
	for v := 0; v < vertices; v++ {
		if indeg[v] == 0 {
			worklist.Enqueue(v)
		}
	}

	for v, ok := worklist.Dequeue(); ok; v, ok = worklist.Dequeue() {
		order = append(order, v)
		for _, w := range g.Adjacent(v) {
			indeg[w]--
			if indeg[w] == 0 {
				worklist.Enqueue(w)
			}
		}
	}

	if len(order) < vertices {
		return nil
	}
	return order
}

func indegree(g ex_digraph.Digraph) []int {
	vertices := g.Vertices()
	indeg := make([]int, vertices)
	for v := 0; v < vertices; v++ {
		for _, w := range g.Adjacent(v) {
			indeg[w]++
		}
	}
	return indeg
}
