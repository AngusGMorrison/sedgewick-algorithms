package list

type node[E any] struct {
	data E
	next *node[E]
}

func removeLast[E any](n *node[E]) *node[E] {
	if n == nil || n.next == nil {
		return nil
	}

	var cur *node[E]
	for cur = n; cur.next.next != nil; cur = cur.next {
	}
	cur.next = nil
	return n
}
