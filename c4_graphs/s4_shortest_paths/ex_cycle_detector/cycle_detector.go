package ex_cycle_detector

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

type CycleDetector struct {
	cycle *stack.SliceStack[*ex_ewdg.DirectedEdge]
}

func NewCycleDetector(g ex_ewdg.EdgeWeightedDigraph) *CycleDetector {
	cd := &CycleDetector{}
	cd.detect(g)
	return cd
}

func (cd *CycleDetector) detect(g ex_ewdg.EdgeWeightedDigraph) {
	vertices := g.V()
	edgeTo := make([]*ex_ewdg.DirectedEdge, vertices)
	seen := make([]bool, vertices)
	onStack := make([]bool, vertices)

	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		onStack[v] = true

		for _, edge := range g.Adj(v) {
			w := edge.To()
			if !seen[w] {
				edgeTo[w] = edge
				visit(w)
				if cd.HasCycle() {
					return
				}
			} else if onStack[w] {
				edgeTo[w] = edge
				cd.buildCycleThrough(v, w, edgeTo)
				return
			}
		}

		onStack[v] = false
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
			if cd.HasCycle() {
				return
			}
		}
	}
}

func (cd *CycleDetector) buildCycleThrough(v, w int, edgeTo []*ex_ewdg.DirectedEdge) {
	cycle := stack.NewSliceStack[*ex_ewdg.DirectedEdge]()
	for edge := edgeTo[v]; edge != edgeTo[w]; edge = edgeTo[edge.From()] {
		cycle.Push(edge)
	}
	cycle.Push(edgeTo[w])
	cycle.Push(edgeTo[v])
	cd.cycle = cycle
}

func (cd *CycleDetector) HasCycle() bool {
	return cd.cycle != nil
}

func (cd *CycleDetector) Cycle() []*ex_ewdg.DirectedEdge {
	if !cd.HasCycle() {
		return nil
	}

	result := make([]*ex_ewdg.DirectedEdge, 0, cd.cycle.Len())
	cd.cycle.Each(func(e *ex_ewdg.DirectedEdge) {
		result = append(result, e)
	})
	return result
}
