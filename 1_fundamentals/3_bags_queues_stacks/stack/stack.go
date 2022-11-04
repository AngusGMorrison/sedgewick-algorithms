package stack

import (
	"errors"
	"fmt"
	"strings"
)

// SliceStack implements a generic stack backed by a slice.
type SliceStack[E any] struct {
	slice []E
}

func NewSliceStack[E any]() *SliceStack[E] {
	return &SliceStack[E]{}
}

func (s *SliceStack[E]) Len() int {
	return len(s.slice)
}

func (s *SliceStack[E]) Push(elem E) {
	s.slice = append(s.slice, elem)
}

func (s *SliceStack[E]) Pop() (E, bool) {
	if len(s.slice) == 0 {
		return *new(E), false
	}

	elem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return elem, true
}

func (s *SliceStack[E]) Peek() (E, bool) {
	if len(s.slice) == 0 {
		return *new(E), false
	}

	return s.slice[len(s.slice)-1], true
}

func (s *SliceStack[E]) Each(f func(elem E)) {
	for i := len(s.slice) - 1; i >= 0; i-- {
		f(s.slice[i])
	}
}

// ListStack implements a generic stack backed by a linked list.
type ListStack[E any] struct {
	len   int
	ops   int // a counter of total stack operations, used to detect stack modification during iteration
	first *node[E]
}

func NewListStack[E any]() *ListStack[E] {
	return &ListStack[E]{}
}

// Q1.3.42
func (ls *ListStack[E]) Clone() *ListStack[E] {
	clone := &ListStack[E]{len: ls.len}
	if ls.len == 0 {
		return clone
	}

	clone.first = &node[E]{data: ls.first.data}
	for origCur, clonePrev := ls.first.next, clone.first; origCur != nil; origCur, clonePrev = origCur.next, clonePrev.next {
		clonePrev.next = &node[E]{
			data: origCur.data,
		}
	}

	return clone
}

func (ls *ListStack[E]) CloneRecursive() *ListStack[E] {
	var cloneFunc func(n *node[E]) *node[E]
	cloneFunc = func(n *node[E]) *node[E] {
		if n == nil {
			return nil
		}

		return &node[E]{
			data: n.data,
			next: cloneFunc(n.next),
		}
	}

	return &ListStack[E]{
		len:   ls.len,
		first: cloneFunc(ls.first),
	}
}

func (ls *ListStack[E]) Len() int {
	return ls.len
}

func (ls *ListStack[E]) Push(elem E) {
	ls.first = &node[E]{
		data: elem,
		next: ls.first,
	}
	ls.len++
	ls.ops++
}

func (ls *ListStack[E]) Pop() (E, bool) {
	if ls.Len() == 0 {
		return *new(E), false
	}

	elem := ls.first.data
	ls.first = ls.first.next
	ls.len--
	ls.ops++

	return elem, true
}

func (ls *ListStack[E]) Peek() (E, bool) {
	if ls.Len() == 0 {
		return *new(E), false
	}

	return ls.first.data, true
}

func (ls *ListStack[E]) Each(f func(elem E)) {
	for cur := ls.first; cur != nil; cur = cur.next {
		f(cur.data)
	}
}

func (ls *ListStack[E]) String() string {
	if ls.len == 0 {
		return "{}"
	}

	var builder strings.Builder
	_, _ = fmt.Fprintf(&builder, "%v", ls.first.data)
	for cur := ls.first.next; cur != nil; cur = cur.next {
		_, _ = fmt.Fprintf(&builder, " -> %v", ls.first.data)
	}

	return builder.String()
}

var ErrConcurrentModification = errors.New("stack was modified during iteration")

type Iterator[E any] interface {
	HasNext() bool
	Next() (E, error)
}

// Q1.3.50
type ListStackIterator[E any] struct {
	stack                   *ListStack[E]
	stackOpsAtInstantiation int
	cur                     *node[E]
}

func (lsi *ListStackIterator[E]) HasNext() bool {
	return lsi.cur != nil
}

func (lsi *ListStackIterator[E]) Next() (E, error) {
	if lsi.stackOpsAtInstantiation != lsi.stack.ops {
		return *new(E), ErrConcurrentModification
	}

	elem := lsi.cur.data
	lsi.cur = lsi.cur.next
	return elem, nil
}

func (ls *ListStack[E]) Iterator() *ListStackIterator[E] {
	return &ListStackIterator[E]{
		stack:                   ls,
		stackOpsAtInstantiation: ls.ops,
		cur:                     ls.first,
	}
}

type node[E any] struct {
	data E
	next *node[E]
}
