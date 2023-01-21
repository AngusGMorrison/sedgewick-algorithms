package ex_directed_dfs

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"

type DirectedDFS struct {
	seen   []bool
	edgeTo []int
	source int
}

func New(g ex_digraph.Digraph, v int) *DirectedDFS {
	vertices := g.Vertices()
	d := &DirectedDFS{
		seen:   make([]bool, vertices),
		edgeTo: make([]int, vertices),
		source: v,
	}
	d.dfs(g, v)
	return d
}

func (d *DirectedDFS) dfs(g ex_digraph.Digraph, v int) {
	d.seen[v] = true
	for _, neighbor := range g.Adjacent(v) {
		if !d.seen[neighbor] {
			d.edgeTo[neighbor] = v
			d.dfs(g, neighbor)
		}
	}
}

func (d *DirectedDFS) HasPathTo(v int) bool {
	return d.seen[v]
}

func (d *DirectedDFS) PathTo(v int) []int {
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
