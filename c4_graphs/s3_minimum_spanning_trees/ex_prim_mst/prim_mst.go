package ex_prim_mst

import (
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_mst"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
)

type PrimMST struct {
	edgeTo   []*ex_edge_weighted_graph.Edge             // edgeTo[i] is the shortest-known edge terminating at vertex i
	distTo   []float64                                  // distTo[i] is the shortest-known distance of the edge terminating at vertex i
	inMST    []bool                                     // inMST[i] is true if the vertex i is in the MST
	weightPQ *heap.SymbolHeap[int, prioritizableWeight] // weightPQ is a min heap that associates an edge weight with the index of the vertex where the edge terminates
	weight   float64                                    // the total weight of the MST
}

var _ ex_mst.MST = (*PrimMST)(nil)

func NewPrimMST(g ex_edge_weighted_graph.EdgeWeightedGraph) *PrimMST {
	nVertices := g.NVertices()
	prim := &PrimMST{
		edgeTo:   make([]*ex_edge_weighted_graph.Edge, nVertices),
		distTo:   make([]float64, nVertices),
		inMST:    make([]bool, nVertices),
		weightPQ: heap.NewSymbolHeap[int, prioritizableWeight](),
	}

	for v := range prim.distTo {
		prim.distTo[v] = math.Inf(1)
	}

	// Vertex 0 is always the origin of the MST.
	prim.distTo[0] = 0
	prim.weightPQ.Push(0, 0)

	// Process each vertex that is not yet in the MST, selecting the closest vertex to the tree to
	// merge next.
	for v, _, ok := prim.weightPQ.Pop(); ok; v, _, ok = prim.weightPQ.Pop() {
		prim.visit(g, v)
	}

	return prim
}

func (prim *PrimMST) visit(g ex_edge_weighted_graph.EdgeWeightedGraph, v int) {
	prim.inMST[v] = true
	prim.weight += prim.distTo[v]

	for _, vEdge := range g.Adjacent(v) {
		w, _ := vEdge.Other(v)
		if prim.inMST[w] || vEdge.Weight() >= prim.distTo[w] {
			continue
		}

		prim.edgeTo[w] = vEdge
		prim.distTo[w] = vEdge.Weight()
		prim.weightPQ.Update(w, prioritizableWeight(vEdge.Weight()))
	}
}

func (prim *PrimMST) Edges() []*ex_edge_weighted_graph.Edge {
	// There is no edge to vertex 0, so return all but the first entry in prim.edgeTo.
	edges := make([]*ex_edge_weighted_graph.Edge, len(prim.edgeTo)-1)
	copy(edges, prim.edgeTo[1:])
	return edges
}

func (prim *PrimMST) Weight() float64 {
	return prim.weight
}

type prioritizableWeight float64

func (a prioritizableWeight) HasPriority(b prioritizableWeight) bool {
	return a < b
}
