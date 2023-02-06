package ex_bellman_ford

import (
	"math"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_cycle_detector"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

type BellmanFordSP struct {
	edgeTo     []*ex_ewdg.DirectedEdge
	weightTo   []float64
	onQueue    []bool
	queue      *queue.SliceQueue[int]
	passes     int
	negativeCD *ex_cycle_detector.CycleDetector
}

func NewBellmanFordSP(g ex_ewdg.EdgeWeightedDigraph, source int) *BellmanFordSP {
	vertices := g.V()
	bf := &BellmanFordSP{
		edgeTo:   make([]*ex_ewdg.DirectedEdge, vertices),
		weightTo: make([]float64, vertices),
		onQueue:  make([]bool, vertices),
		queue:    queue.NewSliceQueue[int](),
	}

	for v := range bf.weightTo {
		bf.weightTo[v] = math.Inf(1)
	}
	bf.weightTo[source] = 0

	bf.queue.Enqueue(source)
	bf.onQueue[source] = true
	bf.queue.Enqueue(math.MinInt) // marker indicating the end of a pass through the graph

	for bf.queue.Len() > 0 {
		v, _ := bf.queue.Dequeue()
		if v == math.MinInt { // we've completed a pass through the queue
			bf.passes++
			// If we've made > V-1 passes, there must be a negative cycle, since any shortest path
			// in a graph without a negative cycle has at most V-1 edges.
			if bf.passes == vertices {
				bf.findNegativeCycle()
				break
			}
			// Mark the end of the next pass.
			bf.queue.Enqueue(math.MinInt)
		} else {
			bf.onQueue[v] = false
			bf.relax(g, v)
		}
	}

	return bf
}

func (bf *BellmanFordSP) HasNegativeCycle() bool {
	return bf.negativeCD.HasCycle()
}

func (bf *BellmanFordSP) relax(g ex_ewdg.EdgeWeightedDigraph, v int) {
	for _, edge := range g.Adj(v) {
		w := edge.To()
		candidateWeight := bf.weightTo[v] + edge.Weight()
		if candidateWeight < bf.weightTo[w] {
			bf.edgeTo[w] = edge
			bf.weightTo[w] = candidateWeight
			if !bf.onQueue[w] {
				bf.queue.Enqueue(w)
				bf.onQueue[w] = true
			}
		}
	}
}

// If the queue is non-empty after the Vth pass, the subgraph of edges in edgeTo contains a negative
// cycle. Hence, we can build a directed graph of edges from edgeTo and search for the cycle using
// the standard CycleDetector implementation.
func (bf *BellmanFordSP) findNegativeCycle() {
	vertices := len(bf.edgeTo)
	cyclicGraph := ex_ewdg.NewWeightedAdjList(vertices)
	for _, edge := range bf.edgeTo {
		if edge != nil {
			cyclicGraph.AddEdge(edge)
		}
	}
	bf.negativeCD = ex_cycle_detector.NewCycleDetector(cyclicGraph)
}

func (bf *BellmanFordSP) NegativeCycle() []*ex_ewdg.DirectedEdge {
	return bf.negativeCD.Cycle()
}
