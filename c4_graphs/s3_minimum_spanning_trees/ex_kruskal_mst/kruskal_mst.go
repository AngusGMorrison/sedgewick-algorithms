package ex_kruskal_mst

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_mst"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/unionfind"
)

type KruskalMST struct {
	edges  []*ex_edge_weighted_graph.Edge
	weight float64
}

var _ ex_mst.MST = (*KruskalMST)(nil)

func NewKruskalMST(g ex_edge_weighted_graph.EdgeWeightedGraph) *KruskalMST {
	kruskal := &KruskalMST{
		edges: make([]*ex_edge_weighted_graph.Edge, 0, g.NVertices()-1), // there are V-1 edges in an MST
	}
	kruskal.build(g)
	return kruskal
}

func (kruskal *KruskalMST) build(g ex_edge_weighted_graph.EdgeWeightedGraph) {
	nVertices := g.NVertices()
	uf := unionfind.NewWeightedQuickUnion(nVertices)                                          // check if vertices are already part of the MST in near-constant time
	edgePQ := heap.NewHeapFromSlice(g.Edges(), func(a, b *ex_edge_weighted_graph.Edge) bool { // heap of all edges, prioritized by least weight
		return a.Weight() < b.Weight()
	})

	// While there is still an edge on the heap and the MST is not complete, find the next-shortest
	// edge that is not yet part of the MST.
	for edge, ok := edgePQ.Pop(); ok && len(kruskal.edges) < cap(kruskal.edges); edge, ok = edgePQ.Pop() {
		v := edge.Either()
		w, _ := edge.Other(v)
		if uf.Connected(v, w) {
			continue
		}

		uf.Union(v, w)
		kruskal.edges = append(kruskal.edges, edge)
		kruskal.weight += edge.Weight()
	}
}

func (kruskal *KruskalMST) Edges() []*ex_edge_weighted_graph.Edge {
	edges := make([]*ex_edge_weighted_graph.Edge, len(kruskal.edges))
	copy(edges, kruskal.edges)
	return edges
}

func (kruskal *KruskalMST) Weight() float64 {
	return kruskal.weight
}
