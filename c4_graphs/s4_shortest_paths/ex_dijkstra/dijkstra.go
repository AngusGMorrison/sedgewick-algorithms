package ex_dijkstra

import (
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_shortest_paths"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
)

type DijkstraSP struct {
	edgeTo   []*ex_ewdg.DirectedEdge
	weightTo []float64
	weightPQ *heap.SymbolHeap[int, prioritizableWeight]
}

func NewDijkstraSP(g ex_ewdg.EdgeWeightedDigraph, source int) *DijkstraSP {
	vertices := g.V()
	dijkstra := &DijkstraSP{
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		weightTo: make([]float64, vertices),
		weightPQ: heap.NewSymbolHeap[int, prioritizableWeight](),
	}

	for v := range dijkstra.weightTo {
		dijkstra.weightTo[v] = math.Inf(1)
	}
	dijkstra.weightTo[source] = 0
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
		candidateWeight := dsp.weightTo[v] + edge.Weight()
		if dsp.weightTo[w] > candidateWeight {
			dsp.weightTo[w] = candidateWeight
			dsp.edgeTo[w] = edge
			dsp.weightPQ.Update(w, prioritizableWeight(dsp.weightTo[w]))
		}
	}
}

func (dsp *DijkstraSP) HasPathTo(v int) bool {
	return !math.IsInf(dsp.weightTo[v], 1)
}

func (dsp *DijkstraSP) WeightTo(v int) float64 {
	return dsp.weightTo[v]
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

var _ ex_shortest_paths.SP = (*DijkstraSP)(nil)

// Find the shortest path from a source vertex to all other vertices on an undirected edge-weighted
// graph by first building a directed edge-weighted graph from the undirected input, then running
// Dijkstra's algorithm.
type UndirectedDijkstraSP struct {
	edgeTo   []*ex_ewdg.DirectedEdge
	weightTo []float64
	weightPQ *heap.SymbolHeap[int, prioritizableWeight]
}

var _ ex_shortest_paths.SP = (*UndirectedDijkstraSP)(nil)

func NewUndirectedDijkstraSP(g ex_edge_weighted_graph.EdgeWeightedGraph, source int) *UndirectedDijkstraSP {
	vertices := g.NVertices()
	dg := newDirectedEWGFromUndirectedEWG(g)
	udsp := &UndirectedDijkstraSP{
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		weightTo: make([]float64, vertices),
		weightPQ: heap.NewSymbolHeap[int, prioritizableWeight](),
	}

	for v := range udsp.weightTo {
		udsp.weightTo[v] = math.Inf(1)
	}

	udsp.weightTo[source] = 0
	udsp.weightPQ.Push(source, 0)

	for v, _, ok := udsp.weightPQ.Pop(); ok; v, _, ok = udsp.weightPQ.Pop() {
		udsp.relax(dg, v)
	}

	return udsp
}

func (udsp *UndirectedDijkstraSP) relax(g ex_ewdg.EdgeWeightedDigraph, v int) {
	for _, edge := range g.Adj(v) {
		w := edge.To()
		candidateWeight := udsp.weightTo[v] + edge.Weight()
		if candidateWeight < udsp.weightTo[w] {
			udsp.edgeTo[w] = edge
			udsp.weightTo[w] = candidateWeight
			udsp.weightPQ.Update(w, prioritizableWeight(candidateWeight))
		}
	}
}

func newDirectedEWGFromUndirectedEWG(ug ex_edge_weighted_graph.EdgeWeightedGraph) *ex_ewdg.AdjListEWDG {
	dg := ex_ewdg.NewWeightedAdjList(ug.NVertices())
	for v, uEdge := range ug.Edges() {
		to, _ := uEdge.Other(v)
		dEdge := ex_ewdg.NewDirectedEdge(v, to, uEdge.Weight())
		dg.AddEdge(dEdge)
	}
	return dg
}

func (udsp *UndirectedDijkstraSP) HasPathTo(v int) bool {
	return !math.IsInf(udsp.weightTo[v], 1)
}

func (udsp *UndirectedDijkstraSP) PathTo(v int) []*ex_ewdg.DirectedEdge {
	if !udsp.HasPathTo(v) {
		return nil
	}

	var edges []*ex_ewdg.DirectedEdge
	for edge := udsp.edgeTo[v]; edge != nil; edge = udsp.edgeTo[edge.From()] {
		edges = append(edges, edge)
	}
	reverse(edges)

	return edges
}

func (udsp *UndirectedDijkstraSP) WeightTo(v int) float64 {
	return udsp.weightTo[v]
}

type SourceSinkDijkstraSP struct {
	sink     int
	edgeTo   []*ex_ewdg.DirectedEdge
	weightTo []float64
	weightPQ *heap.SymbolHeap[int, prioritizableWeight]
}

func NewSourceSinkDijkstraSP(g ex_ewdg.EdgeWeightedDigraph, source, sink int) *SourceSinkDijkstraSP {
	vertices := g.V()
	ssdsp := &SourceSinkDijkstraSP{
		sink:     sink,
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		weightTo: make([]float64, vertices),
		weightPQ: heap.NewSymbolHeap[int, prioritizableWeight](),
	}

	for v := range ssdsp.weightTo {
		ssdsp.weightTo[v] = math.Inf(1)
	}
	ssdsp.weightTo[source] = 0
	ssdsp.weightPQ.Push(source, 0)

	for v, _, ok := ssdsp.weightPQ.Pop(); ok; v, _, ok = ssdsp.weightPQ.Pop() {
		if v == sink {
			break
		}

		ssdsp.relax(g, v)
	}

	return ssdsp
}

func (ssdsp *SourceSinkDijkstraSP) HasPathToSink() bool {
	return !math.IsInf(ssdsp.weightTo[ssdsp.sink], 1)
}

func (ssdsp *SourceSinkDijkstraSP) PathToSink() []*ex_ewdg.DirectedEdge {
	if !ssdsp.HasPathToSink() {
		return nil
	}

	var edges []*ex_ewdg.DirectedEdge
	for edge := ssdsp.edgeTo[ssdsp.sink]; edge != nil; edge = ssdsp.edgeTo[edge.From()] {
		edges = append(edges, edge)
	}
	reverse(edges)

	return edges
}

func (ssdsp *SourceSinkDijkstraSP) WeightToSink() float64 {
	return ssdsp.weightTo[ssdsp.sink]
}

func (ssdsp *SourceSinkDijkstraSP) PathTo(v int)

func (ssdsp *SourceSinkDijkstraSP) relax(g ex_ewdg.EdgeWeightedDigraph, v int) {
	for _, edge := range g.Adj(v) {
		w := edge.To()
		candidateWeight := ssdsp.weightTo[v] + edge.Weight()
		if candidateWeight < ssdsp.weightTo[w] {
			ssdsp.edgeTo[w] = edge
			ssdsp.weightTo[w] = candidateWeight
			ssdsp.weightPQ.Update(w, prioritizableWeight(candidateWeight))
		}
	}
}

// Calculates the shortest path between any pair of vertices using V^2 space (each DijkstraSP
// requires V space, and we need one for each vertex), and VElogV time (we calculate V DijkstraSPs,
// each of which may perform an operation on a heap of V weights for each edge in the graph in the
// worst case).
type AllPairsDijkstraSP struct {
	dijkstraSPs []*DijkstraSP
}

func NewAllPairsDijkstraSP(g ex_ewdg.EdgeWeightedDigraph) *AllPairsDijkstraSP {
	vertices := g.V()
	allPairs := &AllPairsDijkstraSP{
		dijkstraSPs: make([]*DijkstraSP, vertices),
	}

	for v := 0; v < vertices; v++ {
		allPairs.dijkstraSPs[v] = NewDijkstraSP(g, v)
	}

	return allPairs
}

func (ap *AllPairsDijkstraSP) HasPath(from, to int) bool {
	return ap.dijkstraSPs[from].HasPathTo(to)
}

func (ap *AllPairsDijkstraSP) PathBetween(from, to int) []*ex_ewdg.DirectedEdge {
	return ap.dijkstraSPs[from].PathTo(to)
}

func (ap *AllPairsDijkstraSP) WeightBetween(from, to int) float64 {
	return ap.dijkstraSPs[from].WeightTo(to)
}

type prioritizableWeight float64

var _ heap.Prioritizable[prioritizableWeight] = prioritizableWeight(0)

func (a prioritizableWeight) HasPriority(b prioritizableWeight) bool {
	return a < b
}

func reverse[S ~[]E, E any](edges S) {
	for i, j := 0, len(edges)-1; i < j; i, j = i+1, j-1 {
		edges[i], edges[j] = edges[j], edges[i]
	}
}
