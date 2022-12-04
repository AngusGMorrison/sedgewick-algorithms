package q31_doubly_linked_list

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmpty    = errors.New("list was empty")
	ErrNotFound = errors.New("element was not present in the list")
)

type List[E comparable] struct {
	len   int
	first *node[E]
	last  *node[E]
}

func NewList[E comparable]() *List[E] {
	return &List[E]{}
}

func (l *List[E]) Len() int {
	return l.len
}

func (l *List[E]) Prepend(elem E) {
	oldFirst := l.first
	l.first = &node[E]{
		data: elem,
		next: oldFirst,
	}
	if l.len == 0 {
		l.last = l.first
	} else {
		oldFirst.prev = l.first
	}
	l.len++
}

func (l *List[E]) Append(elem E) {
	oldLast := l.last
	l.last = &node[E]{
		data: elem,
		prev: oldLast,
	}
	if l.len == 0 {
		l.first = l.last
	} else {
		oldLast.next = l.last
	}
	l.len++
}

func (l *List[E]) Shift() (E, error) {
	if l.len == 0 {
		return *new(E), ErrEmpty
	}

	oldFirst := l.first
	l.first = oldFirst.next
	if l.len == 1 {
		l.last = nil
	} else {
		l.first.prev = nil
	}
	l.len--

	return oldFirst.data, nil
}

func (l *List[E]) Pop() (E, error) {
	if l.len == 0 {
		return *new(E), ErrEmpty
	}

	oldLast := l.last
	l.last = oldLast.prev
	if l.len == 1 {
		l.first = nil
	} else {
		l.last.next = nil
	}
	l.len--

	return oldLast.data, nil
}

func (l *List[E]) Each(f func(elem E)) {
	for cur := l.first; cur != nil; cur = cur.next {
		f(cur.data)
	}
}

func (l *List[E]) String() string {
	var builder strings.Builder
	var counter int
	l.Each(func(elem E) {
		if counter != 0 {
			builder.WriteString(" ↔ ")
		}
		_, _ = fmt.Fprintf(&builder, "%v", elem)
		counter++
	})
	_, _ = fmt.Fprintf(&builder, " (len %d)", l.len)
	return builder.String()
}

// node is a node in a doubly-linked list.
type node[E comparable] struct {
	data E
	next *node[E]
	prev *node[E]
}

func (l *List[E]) InsertAfter(target, elem E) error {
	n, err := l.find(target)
	if err != nil {
		return err
	}

	oldNext := n.next
	n.next = &node[E]{
		data: elem,
		next: oldNext,
		prev: n,
	}
	if oldNext == nil { // n is the end of the list
		l.last = n.next
	} else {
		oldNext.prev = n.next
	}
	l.len++

	return nil
}

func (l *List[E]) InsertBefore(target, elem E) error {
	n, err := l.find(target)
	if err != nil {
		return err
	}

	oldPrev := n.prev
	n.prev = &node[E]{
		data: elem,
		next: n,
		prev: oldPrev,
	}
	if oldPrev == nil { // n is the start of the list
		l.first = n.prev
	} else {
		oldPrev.next = n.prev
	}
	l.len++

	return nil
}

func (l *List[E]) Remove(target E) error {
	n, err := l.find(target)
	if err != nil {
		return err
	}

	if n.prev == nil { // n is the start of the list
		l.first = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil { // n is the end of the list
		l.last = n.prev
	} else {
		n.next.prev = n.prev
	}
	l.len--

	return nil
}

func (l *List[E]) find(elem E) (*node[E], error) {
	for cur := l.first; cur != nil; cur = cur.next {
		if cur.data == elem {
			return cur, nil
		}
	}

	return nil, fmt.Errorf("find node with data %v: %w", elem, ErrNotFound)
}
