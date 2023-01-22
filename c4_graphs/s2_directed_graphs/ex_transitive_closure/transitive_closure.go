package ex_transitive_closure

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_directed_dfs"
)

type TransitiveClosure struct {
	searches []*ex_directed_dfs.DirectedDFS
}

func NewTransitiveClosure(g ex_digraph.Digraph) *TransitiveClosure {
	vertices := g.Vertices()
	tc := &TransitiveClosure{
		searches: make([]*ex_directed_dfs.DirectedDFS, vertices),
	}
	for v := 0; v < vertices; v++ {
		tc.searches[v] = ex_directed_dfs.New(g, v)
	}
	return tc
}

func (tc *TransitiveClosure) HasPathBetween(v, w int) bool {
	return tc.searches[v].HasPathTo(w)
}

func (tc *TransitiveClosure) PathBetween(v, w int) []int {
	return tc.searches[v].PathTo(w)
}
