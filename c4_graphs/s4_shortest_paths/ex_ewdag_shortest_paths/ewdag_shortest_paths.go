package ex_ewdag_shortest_paths

import (
	"fmt"
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_toposort"
)

type AcyclicSP struct {
	edgeTo   []*ex_ewdg.DirectedEdge
	weightTo []float64
}

func NewAcyclicSP(g ex_ewdg.EdgeWeightedDigraph, source int) (*AcyclicSP, error) {
	vertices := g.V()
	asp := &AcyclicSP{
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		weightTo: make([]float64, vertices),
	}

	for v := range asp.weightTo {
		asp.weightTo[v] = math.Inf(1)
	}
	asp.weightTo[source] = 0

	topo, err := ex_toposort.NewTopoSort(g)
	if err != nil {
		return nil, fmt.Errorf("instantiate TopoSort: %w", err)
	}

	// The topological sort for DAGs ensures that, once a vertex has been relaxed, it will never be
	// seen again by any other path, since we process all edges to the vertex before the vertex
	// itself. Therefore, when a given vertex is encountered, we know that there can be no shorter
	// path to it. We can't topologically sort a graph with cycles, so we can't guarantee that we've
	// processed all edges to a vertex before processing a vertex in a cyclic graph. Hence cyclic
	// graphs require that we maintain a heap of path weights to ensure the shortest is selected.
	for v := range topo.Order() {
		asp.relax(g, v)
	}

	return asp, nil
}

func (asp *AcyclicSP) relax(g ex_ewdg.EdgeWeightedDigraph, v int) {
	for _, edge := range g.Adj(v) {
		w := edge.To()
		candidateWeight := asp.weightTo[v] + edge.Weight()
		if candidateWeight < asp.weightTo[w] {
			asp.edgeTo[w] = edge
			asp.weightTo[w] = candidateWeight
		}
	}
}

func (asp *AcyclicSP) HasPathTo(v int) bool {
	return !math.IsInf(asp.weightTo[v], 1)
}

func (asp *AcyclicSP) WeightTo(v int) float64 {
	return asp.weightTo[v]
}

func (asp *AcyclicSP) PathTo(v int) []*ex_ewdg.DirectedEdge {
	if !asp.HasPathTo(v) {
		return nil
	}

	var path []*ex_ewdg.DirectedEdge
	for edge := asp.edgeTo[v]; edge != nil; edge = asp.edgeTo[edge.From()] {
		path = append(path, edge)
	}
	reverse(path)

	return path
}

func reverse(e []*ex_ewdg.DirectedEdge) {
	for i, j := 0, len(e)-1; i < j; i, j = i+1, j-1 {
		e[i], e[j] = e[j], e[i]
	}
}
