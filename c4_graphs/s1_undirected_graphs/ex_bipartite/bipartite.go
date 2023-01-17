package ex_bipartite

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"

type BipartiteDetector struct {
	isBipartite bool
}

func New(g ex_graph.Graph) *BipartiteDetector {
	bd := &BipartiteDetector{
		isBipartite: true,
	}
	bd.detect(g)
	return bd
}

func (bd *BipartiteDetector) IsBipartite() bool {
	return bd.isBipartite
}

func (db *BipartiteDetector) detect(g ex_graph.Graph) {
	vertices := g.Vertices()
	seen := make([]bool, vertices)
	color := make([]bool, vertices)

	var dfs func(v int)
	dfs = func(v int) {
		seen[v] = true
		for _, w := range g.Adjacent(v) {
			if seen[w] {
				if color[w] == color[v] {
					db.isBipartite = false
					return
				}
				continue
			}

			color[w] = !color[v]
			dfs(w)
			// Return early as soon as we determine that g is not bipartite.
			if !db.isBipartite {
				return
			}
		}
	}

	// DFS each of the components in g.
	for v := 0; v < vertices; v++ {
		if seen[v] {
			continue
		}

		dfs(v)
	}
}
