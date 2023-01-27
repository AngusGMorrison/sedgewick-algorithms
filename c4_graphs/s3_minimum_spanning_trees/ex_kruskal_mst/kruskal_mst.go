package ex_kruskal_mst

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/unionfind"
)

type KruskalMST struct {
	mst    *queue.SliceQueue[*ex_edge_weighted_graph.Edge]
	weight float64
}

func NewKruskalMST(g *ex_edge_weighted_graph.EdgeWeightedGraph) *KruskalMST {
	mst := queue.NewSliceQueue[*ex_edge_weighted_graph.Edge]()
	vertices := g.NVertices()
	uf := unionfind.NewWeightedQuickUnion(vertices)
	minPQ := heap.NewHeapFromSlice(g.Edges(), func(a, b *ex_edge_weighted_graph.Edge) bool {
		return a.Less(b)
	})

	var weight float64
	for edge, ok := minPQ.Pop(); ok && mst.Len() < vertices-1; edge, ok = minPQ.Pop() {
		v := edge.Either()
		w, _ := edge.Other(v)
		if uf.Connected(v, w) {
			continue
		}

		uf.Union(v, w)
		mst.Enqueue(edge)
		weight += edge.Weight()
	}

	return &KruskalMST{
		mst:    mst,
		weight: weight,
	}
}

func (mst *KruskalMST) Edges() *queue.SliceQueue[*ex_edge_weighted_graph.Edge] {
	return mst.mst
}

func (mst *KruskalMST) Weight() float64 {
	return mst.weight
}
