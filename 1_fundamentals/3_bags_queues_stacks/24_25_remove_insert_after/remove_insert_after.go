package after

import "github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"

func removeAfter[E comparable](n *list.Node[E]) {
	if n == nil || n.Next == nil {
		return
	}

	n.Next = n.Next.Next
}

func insertAfter[E comparable](n1, n2 *list.Node[E]) {
	if n1 == nil {
		return
	}
	if n2 == nil {
		n1.Next = nil
		return
	}

	n2.Next = n1.Next
	n1.Next = n2
}
