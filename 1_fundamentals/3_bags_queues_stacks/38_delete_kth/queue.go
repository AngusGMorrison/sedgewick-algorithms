package queue

import (
	"errors"
	"fmt"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"
)

type SliceQueue[E any] struct {
	s []E
}

func (q *SliceQueue[E]) IsEmpty() bool {
	return len(q.s) == 0
}

func (q *SliceQueue[E]) Insert(elem E) {
	q.s = append(q.s, elem)
}

var ErrOutOfBounds = errors.New("index out of bounds")

func (q *SliceQueue[E]) Delete(k int) (E, error) {
	if k >= len(q.s) {
		return *new(E), fmt.Errorf("delete element at index %d: %w", k, ErrOutOfBounds)
	}
	elem := q.s[k]
	q.s = append(q.s[:k], q.s[k+1:]...)

	return elem, nil
}

type ListQueue[E comparable] struct {
	len   int
	first *list.Node[E]
	last  *list.Node[E]
}

func (q *ListQueue[E]) IsEmpty() bool {
	return q.len == 0
}

func (q *ListQueue[E]) Insert(elem E) {
	oldLast := q.last
	q.last = &list.Node[E]{Data: elem}
	if q.len == 0 {
		q.first = q.last
	} else {
		oldLast.Next = q.last
	}
	q.len++
}

func (q *ListQueue[E]) Delete(k int) (E, error) {
	if k >= q.len {
		return *new(E), fmt.Errorf("delete element at index %d: %w", k, ErrOutOfBounds)
	}

	var elem E
	if k == 0 {
		elem = q.first.Data
		q.first = q.first.Next
		if q.len <= 2 {
			q.last = q.first
		}
	} else {
		// find parent
		parent := q.first
		for i := 1; i < k; i++ {
			parent = parent.Next
		}
		elem = parent.Next.Data
		parent.Next = parent.Next.Next
		if parent.Next == nil {
			q.last = parent
		}
	}
	q.len--

	return elem, nil
}

func (l *ListQueue[E]) Each(f func(elem E)) {
	for cur := l.first; cur != nil; cur = cur.Next {
		f(cur.Data)
	}
}

func (l *ListQueue[E]) String() string {
	var builder strings.Builder
	var counter int
	l.Each(func(elem E) {
		if counter != 0 {
			builder.WriteString(" -> ")
		}
		_, _ = fmt.Fprintf(&builder, "%v", elem)
		counter++
	})
	_, _ = fmt.Fprintf(&builder, " (len %d)", l.len)
	return builder.String()
}
