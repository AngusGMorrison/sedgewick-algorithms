package ex_toposort

import (
	"errors"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_cycle_detector"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

var ErrCyclic = errors.New("input graph is cyclic")

type TopoSort struct {
	order *stack.SliceStack[int]
}

func NewTopoSort(g ex_ewdg.EdgeWeightedDigraph) (*TopoSort, error) {
	cycleDetector := ex_cycle_detector.NewCycleDetector(g)
	if cycleDetector.HasCycle() {
		return nil, ErrCyclic
	}

	ts := &TopoSort{
		order: stack.NewSliceStack[int](),
	}
	ts.dfs(g)
	return ts, nil
}

func (ts *TopoSort) dfs(g ex_ewdg.EdgeWeightedDigraph) {
	vertices := g.V()
	seen := make([]bool, vertices)

	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		for _, edge := range g.Adj(v) {
			w := edge.To()
			if !seen[w] {
				visit(w)
			}
		}
		ts.order.Push(v)
	}

	for v := 0; v < vertices; v++ {
		if !seen[v] {
			visit(v)
		}
	}
}

func (ts *TopoSort) Order() []int {
	order := make([]int, 0, ts.order.Len())
	ts.order.Each(func(v int) {
		order = append(order, v)
	})
	return order
}
