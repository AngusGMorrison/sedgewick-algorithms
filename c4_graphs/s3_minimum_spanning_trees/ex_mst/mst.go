package ex_mst

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s3_minimum_spanning_trees/ex_edge_weighted_graph"

type MST interface {
	Edges() []*ex_edge_weighted_graph.Edge
	Weight() float64
}
