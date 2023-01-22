package ex_topological_sort

import (
	"errors"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_dfs_vertex_ordering"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_directed_cycle"
)

var ErrCycle = errors.New("input digraph must be acyclic, but it contains a cycle")

func TopologicalSort(g ex_digraph.Digraph) ([]int, error) {
	cycleDetector := ex_directed_cycle.New(g)
	if cycleDetector.HasCycle() {
		return nil, ErrCycle
	}

	topoSorter := ex_dfs_vertex_ordering.NewDirectedDFSReversePostorder(g)
	return topoSorter.Order(), nil
}
