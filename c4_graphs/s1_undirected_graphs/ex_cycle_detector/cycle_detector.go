package ex_cycle_detector

import (
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

type CycleDetector struct {
	cycle    []int
	hasCycle bool
}

func New(g ex_graph.Graph) *CycleDetector {
	cd := &CycleDetector{}
	cd.detectCycle(g)
	return cd
}

// The method of checking whether a vertex has been seen at any point works only for undirected
// graphs, since for directed graphs two different edges to the same vertex does not necessarily
// indicated a cycle. For example, A->B, B->C, A->C is not cyclic. When detecting cycles in directed
// graphs, we must iterate through the current recursion stack to detect if a given vertex has
// already been seen *on the current path*.
func (c *CycleDetector) detectCycle(g ex_graph.Graph) {
	vertices := g.Vertices()
	seen := make([]bool, vertices)
	// path is the current path being traced by the DFS. For undirected graphs, path is only
	// required if you want to report the vertices of the cycle to the user. Otherwise, it is enough
	// to simply check whether a vertex has been previously seen.
	path := stack.NewSliceStack[int]()

	var dfs func(cur, prev int)
	dfs = func(cur, prev int) {
		seen[cur] = true
		// Add the current vertex to the path.
		path.Push(cur)

		for _, v := range g.Adjacent(cur) {
			if seen[v] {
				// If we've previously seen the vertex and it's not the vertex we've just come from,
				// we've found a cycle.
				if v != prev {
					// Add the repeated vertex to the cycle for inspection by the user.
					path.Push(v)
					c.hasCycle = true
					// Copy the cycle to a slice for easy manipulation by the user.
					c.cycle = make([]int, path.Len())
					path.Each(func(elem int) {
						c.cycle = append(c.cycle, elem)
					})
					// Terminate as soon as we detect a cycle.
					return
				}

				continue
			}

			// Recursively search the immediate neighbors of the current vertex, returning early if
			// a cycle is detected.
			dfs(v, cur)
			if c.hasCycle {
				return
			}
		}

		// Remove the current vertex from the path, since cur may be one of many neighbors of prev,
		// and those neighbors do not (yet) contain cur in their paths.
		_, _ = path.Pop()
	}

	// Visit all components of the graph.
	for v := 0; v < vertices; v++ {
		if !seen[v] {
			dfs(v, v)
		}
	}
}

// HasCycle returns true if the graph contains a cycle.
func (c *CycleDetector) HasCycle() bool {
	return c.hasCycle
}

// Cycle returns the cycle as a slice of vertex IDs. The vertex through which the cycle occurs
// appears twice, including once as the final element of the cycle.
func (c *CycleDetector) Cycle() []int {
	return c.cycle
}
