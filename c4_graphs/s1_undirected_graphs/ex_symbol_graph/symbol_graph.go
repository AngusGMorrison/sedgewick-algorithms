package ex_symbol_graph

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"

type SymbolGraph[K comparable] struct {
	indices map[K]int // maps each key to its index in keys
	keys    []K       // maps an index to a key
	g       ex_graph.Graph
}

// New returns a SymbolGraph with a vertex for each key in keys and no edges.
func New[S ~[]K, K comparable](keys S) *SymbolGraph[K] {
	sg := &SymbolGraph[K]{
		indices: make(map[K]int, len(keys)),
		keys:    make([]K, len(keys)),
		g:       ex_graph.New(len(keys)),
	}
	copy(sg.keys, keys)
	for i, k := range sg.keys {
		sg.indices[k] = i
	}

	return sg
}

func (sg *SymbolGraph[K]) Edges() int {
	return sg.g.Edges()
}

func (sg *SymbolGraph[K]) Vertices() int {
	return sg.g.Vertices()
}

func (sg *SymbolGraph[K]) Adjacent(v int) []int {
	return sg.g.Adjacent(v)
}

func (sg *SymbolGraph[K]) AdjacentByKey(key K) []K {
	idx, ok := sg.indices[key]
	if !ok {
		return nil
	}

	adj := sg.Adjacent(idx)
	adjKeys := make([]K, len(adj))
	for _, v := range adj {
		adjKeys = append(adjKeys, sg.keys[v])
	}

	return adjKeys
}

func (sg *SymbolGraph[K]) AddEdge(k1, k2 K) {
	i1, ok := sg.indices[k1]
	if !ok {
		return
	}
	i2, ok := sg.indices[k2]
	if !ok {
		return
	}
	sg.g.AddEdge(i1, i2)
}

func (sg *SymbolGraph[K]) Contains(key K) bool {
	_, ok := sg.indices[key]
	return ok
}

func (sg *SymbolGraph[K]) AddVertex(key K) {
	if sg.Contains(key) {
		return
	}

	sg.keys = append(sg.keys, key)
	sg.indices[key] = len(sg.keys) - 1
	sg.g.AddVertex()
}

func (sg *SymbolGraph[K]) Name(v int) K {
	return sg.keys[v]
}

func (sg *SymbolGraph[K]) Index(key K) int {
	return sg.indices[key]
}
