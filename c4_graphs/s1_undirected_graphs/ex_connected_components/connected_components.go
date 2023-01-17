package ex_connected_components

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"

type ConnectedComponents struct {
	seen  []bool
	id    []int // site-indexed array of vertex component ids
	count int   // count of components
}

func New(g ex_graph.Graph) *ConnectedComponents {
	vertices := g.Vertices()
	cc := &ConnectedComponents{
		seen: make([]bool, vertices),
		id:   make([]int, vertices),
	}

	for v := 0; v < vertices; v++ {
		if cc.seen[v] {
			continue
		}

		cc.dfs(g, v)
		cc.count++
	}

	return cc
}

func (cc *ConnectedComponents) dfs(g ex_graph.Graph, v int) {
	cc.seen[v] = true
	cc.id[v] = cc.count
	for _, w := range g.Adjacent(v) {
		if cc.seen[w] {
			continue
		}

		cc.dfs(g, w)
	}
}

// Connected returns true if vertices v and w are part of the same component.
func (cc *ConnectedComponents) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

// ID returns the ID of the component to which the vertex belongs.
func (cc *ConnectedComponents) ID(v int) int {
	return cc.id[v]
}

// Count returns the number of components in the graph.
func (cc *ConnectedComponents) Count() int {
	return cc.count
}

// ComponentsOf returns the connected components of the given graph as a slice of bags containing
// the integer vertices making up each component.
func ComponentsOf(g ex_graph.Graph) [][]int {
	cc := New(g)
	components := make([][]int, cc.Count())
	vertices := g.Vertices()
	for v := 0; v < vertices; v++ {
		compID := cc.ID(v)
		components[compID] = append(components[compID], v)
	}

	return components
}
