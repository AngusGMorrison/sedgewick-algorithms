package ex_digraph

type Digraph interface {
	Edges() int
	Vertices() int
	AddEdge(v, w int)
	AddVertex()
	Adjacent(v int) []int
	Reverse() Digraph
}

type AdjacencyList struct {
	list  [][]int
	edges int
}

func NewAdjacencyList(v int) *AdjacencyList {
	return &AdjacencyList{
		list: make([][]int, v),
	}
}

func (al *AdjacencyList) Edges() int {
	return al.edges
}

func (al *AdjacencyList) Vertices() int {
	return len(al.list)
}

func (al *AdjacencyList) AddEdge(v int, w int) {
	al.list[v] = append(al.list[v], w)
	al.edges++
}

func (al *AdjacencyList) AddVertex() {
	al.list = append(al.list, nil)
}

func (al *AdjacencyList) Adjacent(v int) []int {
	adj := make([]int, len(al.list[v]))
	copy(adj, al.list[v])
	return adj
}

func (al *AdjacencyList) Reverse() Digraph {
	reversed := NewAdjacencyList(al.Vertices())
	for v, neighbors := range al.list {
		for _, neighbor := range neighbors {
			reversed.list[neighbor] = append(reversed.list[neighbor], v)
		}
	}
	return reversed
}
