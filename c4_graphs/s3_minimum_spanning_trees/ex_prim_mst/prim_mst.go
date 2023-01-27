package ex_prim_mst

import (
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
)

type PrimMST struct {
	edgeTo     []*ex_edge_weighted_graph.Edge             // shortest edge to tree vertex
	distTo     []float64                                  // distTo[v] = edgeTo[v].weight()
	inTree     []bool                                     // true if v is in the MST
	weightHeap *heap.SymbolHeap[int, prioritizableWeight] // min heap that maps the index of an edge to its weight
	weight     float64
}

type prioritizableWeight float64

func (a prioritizableWeight) HasPriority(b prioritizableWeight) bool {
	return a < b
}

func NewPrimMST(g *ex_edge_weighted_graph.EdgeWeightedGraph) *PrimMST {
	vertices := g.NVertices()
	mst := &PrimMST{
		edgeTo:     make([]*ex_edge_weighted_graph.Edge, vertices),
		distTo:     make([]float64, vertices),
		inTree:     make([]bool, vertices),
		weightHeap: heap.NewSymbolHeap[int, prioritizableWeight](),
	}
	for v := range mst.distTo {
		mst.distTo[v] = math.Inf(1)
	}

	// Initialize the MST with edge 0, weight 0
	mst.distTo[0] = 0
	mst.weightHeap.Push(0, 0)
	for v, _, ok := mst.weightHeap.Pop(); ok; v, _, ok = mst.weightHeap.Pop() {
		mst.visit(g, v)
	}

	return mst
}

func (mst *PrimMST) visit(g *ex_edge_weighted_graph.EdgeWeightedGraph, v int) {
	mst.inTree[v] = true
	mst.weight += mst.distTo[v]
	for _, edge := range g.Adjacent(v) {
		w, _ := edge.Other(v)
		if mst.inTree[w] { // v-w is ineligible
			continue
		}

		if edge.Weight() < mst.distTo[w] {
			mst.edgeTo[w] = edge // edge to double has the smallest weight encountered so far. Push it onto the heap as a candidate edge.
			mst.distTo[w] = edge.Weight()
			mst.weightHeap.Update(w, prioritizableWeight(edge.Weight()))
		}
	}
}

func (mst *PrimMST) Edges() []*ex_edge_weighted_graph.Edge {
	edges := make([]*ex_edge_weighted_graph.Edge, 0, len(mst.edgeTo)-1)
	for v := 1; v < len(mst.edgeTo); v++ {
		edges = append(edges, mst.edgeTo[v])
	}
	return edges
}

func (mst *PrimMST) Weight() float64 {
	return mst.weight
}
