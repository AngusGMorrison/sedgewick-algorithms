package q27_28_max

import (
	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
	"golang.org/x/exp/constraints"
)

func max[E constraints.Ordered](n *list.Node[E]) E {
	if n == nil {
		return *new(E)
	}

	max := n.Data
	for cur := n.Next; cur != nil; cur = cur.Next {
		if cur.Data > max {
			max = cur.Data
		}
	}
	return max
}

func maxRecursive[E constraints.Ordered](n *list.Node[E]) E {
	var max E
	var recurse func(n *list.Node[E])
	recurse = func(n *list.Node[E]) {
		if n == nil {
			return
		}

		if n.Data > max {
			max = n.Data
		}

		recurse(n.Next)
	}

	recurse(n)
	return max
}
