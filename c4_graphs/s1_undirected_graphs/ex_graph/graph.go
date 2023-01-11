package ex_graph

import (
	"fmt"
	"io"
	"strings"
)

// Graph represents an abstract graph of integers.
type Graph interface {
	Edges() int
	Vertices() int
	AddEdge(v, w int)
	Adjacent(v int) []int
	String() string
}

// New creates a v-vertex graph with no edges.
func New(v int) Graph {
	return &adjacencyList{
		adj: make([][]int, v),
	}
}

// ReadFrom builds a graph from the given input stream.
func ReadFrom(in io.Reader) Graph {
	return &adjacencyList{
		adj: make([][]int, 0),
	}
}

// Degree returns the degree of the given vertex.
func Degree(g Graph, v int) int {
	return len(g.Adjacent(v))
}

// MaxDegree returns the greatest degree of any vertex in the graph.
func MaxDegree(g Graph) int {
	var max int
	vertices := g.Vertices()
	for v := 0; v < vertices; v++ {
		if degree := Degree(g, v); degree > max {
			max = degree
		}
	}
	return max
}

// AvgDegree returns the average degree of each node in the graph.
func AvgDegree(g Graph) float64 {
	// The number of edges is doubled, since in an undirected graph each edge points to two nodes.
	return (2 * float64(g.Edges())) / float64(g.Vertices())
}

// Self loops returns the number of time any node in g links to itself.
func SelfLoops(g Graph) int {
	var selfLoops int
	vertices := g.Vertices()
	for v := 0; v < vertices; v++ {
		adj := g.Adjacent(v)
		for _, w := range adj {
			if v == w {
				selfLoops++
			}
		}
	}
	return selfLoops
}

// adjacencyList implements Graph using an adjancency list, where the edges of node i are found at index i.
type adjacencyList struct {
	adj   [][]int
	edges int
}

// Edges returns the number of edges in the graph.
func (g *adjacencyList) Edges() int {
	return g.edges
}

// Vertices returns the number of vertices in the graph.
func (g *adjacencyList) Vertices() int {
	return len(g.adj)
}

// AddEdge adds edge v-w to the graph.
func (g *adjacencyList) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.edges++
}

// Adjacent returns a copy of the list of nodes adjacent to v.
func (g *adjacencyList) Adjacent(v int) []int {
	cp := make([]int, len(g.adj[v]))
	copy(cp, g.adj[v])
	return cp
}

// String returns a string representation of the graph's adjacency list.
func (g *adjacencyList) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "%d vertices, %d edges\n", g.Vertices(), g.edges)
	for vertex, edges := range g.adj {
		fmt.Fprintf(&builder, "%d: ", vertex)
		for edge := range edges {
			fmt.Fprintf(&builder, "%d ", edge)
		}
		fmt.Fprintln(&builder)
	}

	return builder.String()
}
