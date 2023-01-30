package ex_ewdg

import (
	"fmt"
	"strings"
)

type DirectedEdge struct {
	from, to int
	weight   float64
}

func (de *DirectedEdge) From() int {
	return de.from
}

func (de *DirectedEdge) To() int {
	return de.to
}

func (de *DirectedEdge) Weight() float64 {
	return de.weight
}

func (de *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %.2f", de.from, de.to, de.weight)
}

type EdgeWeightedDigraph interface {
	fmt.Stringer

	V() int
	E() int
	Edges() []*DirectedEdge
	Adj(v int) []*DirectedEdge
	AddEdge(e *DirectedEdge)
	AddVertex()
}

type AdjListEWDG struct {
	adjList [][]*DirectedEdge
	e       int
}

var _ EdgeWeightedDigraph = (*AdjListEWDG)(nil)

func NewWeightedAdjList(v int) *AdjListEWDG {
	return &AdjListEWDG{
		adjList: make([][]*DirectedEdge, v),
	}
}

func (al *AdjListEWDG) V() int {
	return len(al.adjList)
}
func (al *AdjListEWDG) E() int {
	return al.e
}

func (al *AdjListEWDG) Edges() []*DirectedEdge {
	edges := make([]*DirectedEdge, 0, al.e)
	for _, vEdges := range al.adjList {
		edges = append(edges, vEdges...)
	}
	return edges
}

func (al *AdjListEWDG) Adj(v int) []*DirectedEdge {
	return al.adjList[v]
}

func (al *AdjListEWDG) AddEdge(e *DirectedEdge) {
	from := e.From()
	al.adjList[from] = append(al.adjList[from], e)
	al.e++
}

func (al *AdjListEWDG) AddVertex() {
	al.adjList = append(al.adjList, nil)
}

func (al *AdjListEWDG) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Vertices: %d, Edges: %d\n", al.V(), al.e)
	for v, vEdges := range al.adjList {
		fmt.Fprintf(&sb, "%d: ", v)
		for _, vEdge := range vEdges {
			fmt.Fprintf(&sb, "(%d %.2f) ", vEdge.To(), vEdge.Weight())
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
