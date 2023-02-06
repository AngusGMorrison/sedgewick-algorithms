package ex_bellman_ford

import (
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
)

func Test_BellmanFordSP_HasNegativeCycle(t *testing.T) {
	t.Parallel()

	zeroOne := ex_ewdg.NewDirectedEdge(0, 1, 1)
	oneTwo := ex_ewdg.NewDirectedEdge(1, 2, -5)
	twoThree := ex_ewdg.NewDirectedEdge(2, 3, -3)
	threeOne := ex_ewdg.NewDirectedEdge(3, 1, -5)
	g := ex_ewdg.NewWeightedAdjList(4)
	g.AddEdge(zeroOne)
	g.AddEdge(oneTwo)
	g.AddEdge(twoThree)
	g.AddEdge(threeOne)

	bf := NewBellmanFordSP(g, 0)
	t.Log(bf.HasNegativeCycle())
	t.Log(bf.NegativeCycle())
}
