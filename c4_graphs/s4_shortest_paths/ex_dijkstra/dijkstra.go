package ex_dijkstra

import (
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_shortest_paths"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
)

type DijkstraSP struct {
	edgeTo   []*ex_ewdg.DirectedEdge
	distTo   []float64
	weightPQ *heap.SymbolHeap[int, prioritizableWeight]
}

func NewDijkstraSP(g ex_ewdg.EdgeWeightedDigraph, source int) *DijkstraSP {
	vertices := g.V()
	dijkstra := &DijkstraSP{
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		distTo:   make([]float64, vertices),
		weightPQ: heap.NewSymbolHeap[int, prioritizableWeight](),
	}

	for v := range dijkstra.distTo {
		dijkstra.distTo[v] = math.Inf(1)
	}
	dijkstra.distTo[source] = 0
	dijkstra.weightPQ.Push(source, 0)

	// On each iteration, process the vertex that is closest to the source. I.e. continually make
	// the locally optimal choice. Because we use non-negative values and choose the lowest distTo
	// value at each step, once popped from the heap, distTo[v] never changes and no subsequent
	// relaxation can set any distTo entry to a lower value than distTo[v].
	for v, _, ok := dijkstra.weightPQ.Pop(); ok; v, _, ok = dijkstra.weightPQ.Pop() {
		dijkstra.relax(g, v)
	}

	return dijkstra
}

func (dsp *DijkstraSP) relax(g ex_ewdg.EdgeWeightedDigraph, v int) {
	for _, edge := range g.Adj(v) {
		w := edge.To()
		candidateWeight := dsp.distTo[v] + edge.Weight()
		if dsp.distTo[w] > candidateWeight {
			dsp.distTo[w] = candidateWeight
			dsp.edgeTo[w] = edge
			dsp.weightPQ.Update(w, prioritizableWeight(dsp.distTo[w]))
		}
	}
}

func (dsp *DijkstraSP) HasPathTo(v int) bool {
	return !math.IsInf(dsp.distTo[v], 1)
}

func (dsp *DijkstraSP) WeightTo(v int) float64 {
	return dsp.distTo[v]
}

func (dsp *DijkstraSP) PathTo(v int) []*ex_ewdg.DirectedEdge {
	if !dsp.HasPathTo(v) {
		return nil
	}

	var edges []*ex_ewdg.DirectedEdge
	for edge := dsp.edgeTo[v]; edge != nil; edge = dsp.edgeTo[edge.From()] {
		edges = append(edges, edge)
	}
	reverse(edges)

	return edges
}

func reverse(edges []*ex_ewdg.DirectedEdge) {
	for i, j := 0, len(edges)-1; i < j; i, j = i+1, j-1 {
		edges[i], edges[j] = edges[j], edges[i]
	}
}

var _ ex_shortest_paths.SP = (*DijkstraSP)(nil)

type prioritizableWeight float64

var _ heap.Prioritizable[prioritizableWeight] = prioritizableWeight(0)

func (a prioritizableWeight) HasPriority(b prioritizableWeight) bool {
	return a < b
}
