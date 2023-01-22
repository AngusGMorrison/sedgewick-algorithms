package ex_dfs_vertex_ordering

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

type DirectedDFSPreorder struct {
	order []int
}

func NewDirectedDFSPreorder(g ex_digraph.Digraph) *DirectedDFSPreorder {
	preorder := &DirectedDFSPreorder{
		order: make([]int, g.Vertices()),
	}
	preorder.dfs(g)
	return preorder
}

func (d *DirectedDFSPreorder) dfs(g ex_digraph.Digraph) {
	vertices := g.Vertices()
	seen := make([]bool, vertices)

	var visit func(v int)
	visit = func(v int) {
		d.order = append(d.order, v)
		seen[v] = true
		for _, neighbor := range g.Adjacent(v) {
			if !seen[neighbor] {
				visit(neighbor)
			}
		}
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
		}
	}
}

func (d *DirectedDFSPreorder) Order() []int {
	order := make([]int, len(d.order))
	copy(order, d.order)
	return order
}

type DirectedDFSPostorder struct {
	order []int
}

func NewDirectedDFSPostorder(g ex_digraph.Digraph) *DirectedDFSPostorder {
	postorder := &DirectedDFSPostorder{
		order: make([]int, g.Vertices()),
	}
	postorder.dfs(g)
	return postorder
}

func (d *DirectedDFSPostorder) dfs(g ex_digraph.Digraph) {
	vertices := g.Vertices()
	seen := make([]bool, vertices)

	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		for _, neighbor := range g.Adjacent(v) {
			if !seen[neighbor] {
				visit(neighbor)
			}
		}
		d.order = append(d.order, v)
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
		}
	}
}

func (d *DirectedDFSPostorder) Order() []int {
	order := make([]int, len(d.order))
	copy(order, d.order)
	return order
}

type DirectedDFSReversePostorder struct {
	order *stack.SliceStack[int]
}

func NewDirectedDFSReversePostorder(g ex_digraph.Digraph) *DirectedDFSReversePostorder {
	reversePost := &DirectedDFSReversePostorder{
		order: stack.NewSliceStack[int](),
	}
	reversePost.dfs(g)
	return reversePost
}

func (d *DirectedDFSReversePostorder) dfs(g ex_digraph.Digraph) {
	vertices := g.Vertices()
	seen := make([]bool, vertices)

	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		for _, neighbor := range g.Adjacent(v) {
			if !seen[neighbor] {
				visit(neighbor)
			}
		}
		d.order.Push(v)
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
		}
	}
}

func (d *DirectedDFSReversePostorder) Order() []int {
	order := make([]int, 0, d.order.Len())
	d.order.Each(func(v int) {
		order = append(order, v)
	})
	return order
}
