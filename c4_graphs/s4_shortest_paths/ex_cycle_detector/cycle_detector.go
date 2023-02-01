package ex_cycle_detector

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"

type CycleDetector struct {
	cycle []*ex_ewdg.DirectedEdge
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
				cd.cycle = append(cd.cycle, edge)
				for prev := edgeTo[v]; prev != edge; prev = edgeTo[prev.From()] {
					cd.cycle = append(cd.cycle, prev)
				}
				cd.cycle = append(cd.cycle, edge)
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

func (cd *CycleDetector) HasCycle() bool {
	return cd.cycle != nil
}

func (cd *CycleDetector) Cycle() []*ex_ewdg.DirectedEdge {
	return cd.cycle
}
