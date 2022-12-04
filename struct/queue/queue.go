package queue

// SliceQueue implements a queue backed by a slice, where elements are appended to the end of the
// slice and removed from the start.
type SliceQueue[E any] struct {
	slice []E
}

func NewSliceQueue[E any]() *SliceQueue[E] {
	return &SliceQueue[E]{}
}

func (sq *SliceQueue[E]) Len() int {
	return len(sq.slice)
}

func (sq *SliceQueue[E]) Enqueue(elem E) {
	sq.slice = append(sq.slice, elem)
}

func (sq *SliceQueue[E]) Dequeue() (E, bool) {
	if sq.Len() == 0 {
		return *new(E), false
	}

	elem := sq.slice[0]
	sq.slice = sq.slice[1:]
	return elem, true
}

func (sq *SliceQueue[E]) Each(f func(elem E)) {
	for _, elem := range sq.slice {
		f(elem)
	}
}

// ListQueue implements a queue backed by a linked list, where items are enqueued at the end of the
// list and dequeued from the head of the list.
type ListQueue[E any] struct {
	len         int
	first, last *node[E]
}

func (lq *ListQueue[E]) Len() int {
	return lq.len
}

func (lq *ListQueue[E]) Enqueue(elem E) {
	oldLast := lq.last
	lq.last = &node[E]{data: elem}
	// If the list is empty, we know that oldLast is nil, so we don't have to update is next
	// pointer.
	if lq.Len() == 0 {
		lq.first = lq.last
	} else {
		oldLast.next = lq.last
	}
	lq.len++
}

func (lq *ListQueue[E]) Dequeue() (E, bool) {
	if lq.Len() == 0 {
		return *new(E), false
	}

	oldFirst := lq.first
	lq.first = oldFirst.next
	if lq.first == nil {
		lq.last = nil
	}
	lq.len--

	return oldFirst.data, true
}

func (lq *ListQueue[E]) Each(f func(elem E)) {
	for cur := lq.first; cur != nil; cur = cur.next {
		f(cur.data)
	}
}

type node[E any] struct {
	data E
	next *node[E]
}
