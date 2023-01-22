package ex_kosaraju_scc

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_dfs_vertex_ordering"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
)

type StrongComponents struct {
	componentIDs []int
	seen         []bool
	count        int
}

func NewStrongComponents(g ex_digraph.Digraph) *StrongComponents {
	sc := &StrongComponents{
		componentIDs: make([]int, g.Vertices()),
		seen:         make([]bool, g.Vertices()),
	}
	reverse := g.Reverse()
	postorder := ex_dfs_vertex_ordering.NewDirectedDFSReversePostorder(reverse)
	for _, v := range postorder.Order() {
		if !sc.seen[v] {
			sc.dfs(g, v)
			sc.count++
		}
	}

	return sc
}

func (sc *StrongComponents) dfs(g ex_digraph.Digraph, v int) {
	sc.seen[v] = true
	sc.componentIDs[v] = sc.count
	for _, neighbor := range g.Adjacent(v) {
		if !sc.seen[neighbor] {
			sc.dfs(g, neighbor)
		}
	}
}

func (sc *StrongComponents) Count() int {
	return sc.count
}

func (sc *StrongComponents) StronglyConnected(v, w int) bool {
	return sc.componentIDs[v] == sc.componentIDs[w]
}

func (sc *StrongComponents) ID(v int) int {
	return sc.componentIDs[v]
}

func (sc *StrongComponents) Components() [][]int {
	comps := make([][]int, sc.count)
	for v, compID := range sc.componentIDs {
		comps[compID] = append(comps[compID], v)
	}
	return comps
}
