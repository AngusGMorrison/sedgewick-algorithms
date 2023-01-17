package ex_symbol_graph

import "github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s1_undirected_graphs/ex_graph"

type SymbolGraph struct {
	indices map[string]int // maps each string key to its index in keys
	keys    []string       // maps an index to a string key
	g       ex_graph.Graph
}

// New returns a SymbolGraph with a vertex for each key in keys and no edges.
func New(keys []string) *SymbolGraph {
	sg := &SymbolGraph{
		indices: make(map[string]int, len(keys)),
		keys:    make([]string, len(keys)),
		g:       ex_graph.New(len(keys)),
	}
	copy(sg.keys, keys)
	for i, k := range sg.keys {
		sg.indices[k] = i
	}

	return sg
}

func (sg *SymbolGraph) Edges() int {
	return sg.g.Edges()
}

func (sg *SymbolGraph) Vertices() int {
	return sg.g.Vertices()
}

func (sg *SymbolGraph) Adjacent(v int) []int {
	return sg.g.Adjacent(v)
}

func (sg *SymbolGraph) AdjacentByKey(key string) []string {
	idx, ok := sg.indices[key]
	if !ok {
		return nil
	}

	adj := sg.Adjacent(idx)
	adjKeys := make([]string, len(adj))
	for _, v := range adj {
		adjKeys = append(adjKeys, sg.keys[v])
	}

	return adjKeys
}

func (sg *SymbolGraph) AddEdge(k1, k2 string) {
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

func (sg *SymbolGraph) Contains(key string) bool {
	_, ok := sg.indices[key]
	return ok
}

func (sg *SymbolGraph) AddVertex(key string) {
	if sg.Contains(key) {
		return
	}

	sg.keys = append(sg.keys, key)
	sg.indices[key] = len(sg.keys) - 1
	sg.g.AddVertex()
}

func (sg *SymbolGraph) Name(v int) string {
	return sg.keys[v]
}

func (sg *SymbolGraph) Index(key string) int {
	return sg.indices[key]
}
