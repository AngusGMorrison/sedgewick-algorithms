package ex_directed_cycle

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s2_directed_graphs/ex_digraph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

type DirectedCycleDetector struct {
	seen    []bool
	edgeTo  []int
	onStack []bool
	cycle   *stack.SliceStack[int]
}

func New(g ex_digraph.Digraph) *DirectedCycleDetector {
	vertices := g.Vertices()
	d := &DirectedCycleDetector{
		seen:    make([]bool, vertices),
		edgeTo:  make([]int, vertices),
		onStack: make([]bool, vertices),
	}
	for v := 0; v < vertices; v++ {
		if !d.seen[v] {
			d.dfs(g, v)
			if d.HasCycle() {
				break
			}
		}
	}

	return d
}

func (d *DirectedCycleDetector) dfs(g ex_digraph.Digraph, v int) {
	d.seen[v] = true
	d.onStack[v] = true

	for _, w := range g.Adjacent(v) {
		if !d.seen[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
			if d.HasCycle() {
				return
			}
		} else if d.onStack[w] {
			d.buildCycleThrough(v, w)
			return
		}
	}

	d.onStack[v] = false
}

func (d *DirectedCycleDetector) buildCycleThrough(v, w int) {
	d.cycle = stack.NewSliceStack[int]()
	for x := v; x != w; x = d.edgeTo[x] {
		d.cycle.Push(x)
	}
	d.cycle.Push(w)
	d.cycle.Push(v)
}

func (d *DirectedCycleDetector) HasCycle() bool {
	return d.cycle != nil
}

func (d *DirectedCycleDetector) Cycle() []int {
	cycle := make([]int, 0, d.cycle.Len())
	d.cycle.Each(func(v int) {
		cycle = append(cycle, v)
	})
	return cycle
}
