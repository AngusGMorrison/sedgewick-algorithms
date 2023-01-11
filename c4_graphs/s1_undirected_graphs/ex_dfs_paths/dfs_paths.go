package ex_dfs_paths

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"
)

type DepthFirstPaths struct {
	seen   []bool
	edgeTo []int
	source int
}

func New(g ex_graph.Graph, source int) *DepthFirstPaths {
	vertices := g.Vertices()
	dfp := DepthFirstPaths{
		seen:   make([]bool, vertices),
		edgeTo: make([]int, vertices),
		source: source,
	}
	dfp.dfs(g, source)
	return &dfp
}

func (dfp *DepthFirstPaths) dfs(g ex_graph.Graph, vertex int) {
	// Marking the vertex as seen when we search on it rather than when we first process it
	// prevents the creation of paths that link back to the source. If this is desirable, mark the
	// vertex as seen when iterating through the list of neighbors that contains it.
	dfp.seen[vertex] = true
	for _, neighbor := range g.Adjacent(vertex) {
		if dfp.seen[neighbor] {
			continue
		}

		dfp.edgeTo[neighbor] = vertex
		dfp.dfs(g, neighbor)
	}
}

// HasPathTo returns true if there is a path from the source to v in the graph.
func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	return dfp.seen[v]
}

// PathTo returns a path to the given node, if such a path exists.
func (dfp *DepthFirstPaths) PathTo(v int) []int {
	if !dfp.HasPathTo(v) {
		return nil
	}
	var path []int
	for x := v; x != dfp.source; x = dfp.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, dfp.source)
	reverse(path)
	return path
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
