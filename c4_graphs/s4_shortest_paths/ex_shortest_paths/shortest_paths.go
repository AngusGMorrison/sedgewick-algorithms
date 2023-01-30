package ex_shortest_paths

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"

type SP interface {
	HasPathTo(v int) bool
	PathTo(v int) []*ex_ewdg.DirectedEdge
	WeightTo(v int) float64
}
