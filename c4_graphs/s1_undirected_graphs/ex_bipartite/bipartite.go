package ex_bipartite

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"

func IsBipartite(g ex_graph.Graph) bool {
	vertices := g.Vertices()
	seen := make([]bool, vertices)
	color := make([]bool, vertices)
	isBipartite := true

	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		for _, w := range g.Adjacent(v) {
			if !seen[w] {
				color[w] = !color[v]
				visit(w)
				if !isBipartite {
					return
				}
			} else if color[w] == color[v] {
				isBipartite = false
				return
			}
		}
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
			if !isBipartite {
				return false
			}
		}
	}

	return true
}
