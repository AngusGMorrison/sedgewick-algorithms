package q29_circular_queue

import (
	"errors"
	"fmt"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

type Queue[E comparable] struct {
	len  int
	last *list.Node[E]
}

func (q *Queue[E]) Len() int {
	return q.len
}

func (q *Queue[E]) Enqueue(elem E) {
	oldLast := q.last
	q.last = &list.Node[E]{Data: elem}
	if q.Len() == 0 {
		q.last.Next = q.last
	} else {
		q.last.Next = oldLast.Next
		oldLast.Next = q.last
	}

	q.len++
}

var ErrQueueEmpty = errors.New("attempted to dequeue empty queue")

func (q *Queue[E]) Dequeue() (E, error) {
	if q.Len() == 0 {
		return *new(E), ErrQueueEmpty
	}

	elem := q.last.Next.Data
	if q.Len() == 1 {
		q.last = nil
	} else {
		q.last.Next = q.last.Next.Next
	}

	q.len--
	return elem, nil
}

func (q *Queue[E]) Each(f func(elem E)) {
	cur := q.last
	for i := 0; i < q.len; i++ { // use the length of the list to avoid dereferencing q.last if nil
		f(cur.Next.Data)
		cur = cur.Next
	}
}

func (q *Queue[E]) String() string {
	var builder strings.Builder
	q.Each(func(elem E) {
		_, _ = fmt.Fprintf(&builder, "-> %v ", elem)
	})
	builder.WriteString("->")

	return builder.String()
}
