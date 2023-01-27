package ex_edge_weighted_graph

import (
	"fmt"
	"strings"
)

type EdgeWeightedGraph interface {
	fmt.Stringer

	NVertices() int
	NEdges() int
	Edges() []*Edge
	Adjacent(v int) []*Edge
	AddEdge(edge *Edge)
	AddVertex()
}

type WeightedAdjList struct {
	edges int
	adj   [][]*Edge
}

var _ EdgeWeightedGraph = (*WeightedAdjList)(nil)

func NewWeightedAdjList(vertices int) *WeightedAdjList {
	return &WeightedAdjList{
		adj: make([][]*Edge, vertices),
	}
}

func (ewg *WeightedAdjList) NVertices() int {
	return len(ewg.adj)
}

func (ewg *WeightedAdjList) NEdges() int {
	return ewg.edges
}

func (ewg *WeightedAdjList) AddEdge(e *Edge) {
	v := e.Either()
	w, _ := e.Other(v)
	ewg.adj[v] = append(ewg.adj[v], e)
	ewg.adj[w] = append(ewg.adj[w], e)
	ewg.edges++
}

func (ewg *WeightedAdjList) AddVertex() {
	ewg.adj = append(ewg.adj, nil)
}

func (ewg *WeightedAdjList) Adjacent(v int) []*Edge {
	return ewg.adj[v]
}

func (ewg *WeightedAdjList) Edges() []*Edge {
	edges := make([]*Edge, ewg.NVertices())
	for v, vertexEdges := range ewg.adj {
		for _, vertexEdge := range vertexEdges {
			if other, _ := vertexEdge.Other(v); other > v { // ensure edges aren't added twice
				edges = append(edges, vertexEdge)
			}
		}
	}

	return edges
}

func (ewg *WeightedAdjList) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Vertices: %d, Edges: %d\n", ewg.NVertices(), ewg.edges)
	for v, edges := range ewg.adj {
		fmt.Fprintf(&sb, "%d: ", v)
		for _, e := range edges {
			other, _ := e.Other(v)
			fmt.Fprintf(&sb, "(%d %.2f) ", other, e.weight)
		}
		fmt.Fprintln(&sb)
	}

	return sb.String()
}

type Edge struct {
	v, w   int // the connected vertices
	weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	return &Edge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Either() int {
	return e.v
}

func (e *Edge) Other(vertex int) (int, error) {
	if vertex == e.v {
		return e.w, nil
	}
	if vertex == e.w {
		return e.v, nil
	}

	return 0, &InvalidVertexError{
		vertex: vertex,
		edge:   e,
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}

type InvalidVertexError struct {
	vertex int
	edge   *Edge
}

func (e *InvalidVertexError) Error() string {
	return fmt.Sprintf("vertex %d not present in edge %s", e.vertex, e.edge)
}
