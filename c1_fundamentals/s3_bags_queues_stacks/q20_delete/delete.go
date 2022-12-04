package q20_delete

type node[E any] struct {
	data E
	next *node[E]
}

// Delete the kth element of the list, if it exists. The list is indexed from k == 1.
func delete[E any](n *node[E], k uint) *node[E] {
	if n == nil || k < 1 {
		return n
	}

	if k == 1 {
		return n.next
	}

	cur := n
	for i := uint(0); i < k-2 && cur.next != nil; i++ {
		cur = cur.next
	}
	if cur.next == nil {
		return n // no kth element
	}
	cur.next = cur.next.next
	return n
}
