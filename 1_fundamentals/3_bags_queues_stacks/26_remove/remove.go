package remove

import "github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"

func remove[E comparable](n *list.Node[E], key E) *list.Node[E] {
	if n == nil {
		return nil
	}

	var first *list.Node[E]
	// Fast-forward to the first nonmatching node.
	for first = n; first != nil && first.Data == key; first = first.Next {
	}
	if first == nil {
		return nil
	}

	for cur := first; cur.Next != nil; {
		if cur.Next.Data == key {
			cur.Next = cur.Next.Next
			continue
		}

		cur = cur.Next
	}

	return first
}
