package list

type node[E comparable] struct {
	data E
	next *node[E]
}

func find[E comparable](head *node[E], key E) bool {
	for cur := head; cur != nil; cur = cur.next {
		if cur.data == key {
			return true
		}
	}

	return false
}
