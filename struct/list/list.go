package list

import (
	"errors"
	"fmt"
	"strings"
)

// Node is a node in a singly linked list.
type Node[E any] struct {
	Data E
	Next *Node[E]
}

// Len returns the length of the list from n.
func (n *Node[E]) Len() int {
	var len int
	for cur := n; cur != nil; cur = cur.Next {
		len++
	}
	return len
}

// String returns the string representation of the list.
func (n *Node[E]) String() string {
	if n == nil {
		return "nil"
	}

	var builder strings.Builder
	fmt.Fprintf(&builder, "%v", n.Data)
	for cur := n.Next; cur != nil; cur = cur.Next {
		fmt.Fprintf(&builder, " -> %v", cur.Data)
	}

	return builder.String()
}

// Each performs the given operation for each element of the list.
func (n *Node[E]) Each(f func(elem E)) {
	for cur := n; cur != nil; cur = cur.Next {
		f(n.Data)
	}
}

func (n *Node[E]) AtIndex(i int) (*Node[E], error) {
	var cur *Node[E]
	for j, cur := 0, n; j < i && cur != nil; cur = cur.Next {
	}
	if cur == nil {
		return nil, ErrOutOfBounds
	}
	return cur, nil
}

var (
	ErrEmpty       = errors.New("list was empty")
	ErrNotFound    = errors.New("element was not present in the list")
	ErrOutOfBounds = errors.New("index out of bounds")
)

type Double[E comparable] struct {
	Len   int
	First *DNode[E]
	Last  *DNode[E]
}

func NewList[E comparable]() *Double[E] {
	return &Double[E]{}
}

func (l *Double[E]) Prepend(elem E) {
	oldFirst := l.First
	l.First = &DNode[E]{
		Data: elem,
		Next: oldFirst,
	}
	if l.Len == 0 {
		l.Last = l.First
	} else {
		oldFirst.Prev = l.First
	}
	l.Len++
}

func (l *Double[E]) Append(elem E) {
	oldLast := l.Last
	l.Last = &DNode[E]{
		Data: elem,
		Prev: oldLast,
	}
	if l.Len == 0 {
		l.First = l.Last
	} else {
		oldLast.Next = l.Last
	}
	l.Len++
}

func (l *Double[E]) Shift() (E, error) {
	if l.Len == 0 {
		return *new(E), ErrEmpty
	}

	oldFirst := l.First
	l.First = oldFirst.Next
	if l.Len == 1 {
		l.Last = nil
	} else {
		l.First.Prev = nil
	}
	l.Len--

	return oldFirst.Data, nil
}

func (l *Double[E]) Pop() (E, error) {
	if l.Len == 0 {
		return *new(E), ErrEmpty
	}

	oldLast := l.Last
	l.Last = oldLast.Prev
	if l.Len == 1 {
		l.First = nil
	} else {
		l.Last.Next = nil
	}
	l.Len--

	return oldLast.Data, nil
}

func (l *Double[E]) Each(f func(elem E)) {
	for cur := l.First; cur != nil; cur = cur.Next {
		f(cur.Data)
	}
}

func (l *Double[E]) String() string {
	var builder strings.Builder
	var counter int
	l.Each(func(elem E) {
		if counter != 0 {
			builder.WriteString(" ↔ ")
		}
		_, _ = fmt.Fprintf(&builder, "%v", elem)
		counter++
	})
	_, _ = fmt.Fprintf(&builder, " (len %d)", l.Len)
	return builder.String()
}

func (l *Double[E]) InsertAfter(target, elem E) error {
	n, err := l.Find(target)
	if err != nil {
		return err
	}

	oldNext := n.Next
	n.Next = &DNode[E]{
		Data: elem,
		Next: oldNext,
		Prev: n,
	}
	if oldNext == nil { // n is the end of the list
		l.Last = n.Next
	} else {
		oldNext.Prev = n.Next
	}
	l.Len++

	return nil
}

func (l *Double[E]) InsertBefore(target, elem E) error {
	n, err := l.Find(target)
	if err != nil {
		return err
	}

	oldPrev := n.Prev
	n.Prev = &DNode[E]{
		Data: elem,
		Next: n,
		Prev: oldPrev,
	}
	if oldPrev == nil { // n is the start of the list
		l.First = n.Prev
	} else {
		oldPrev.Next = n.Prev
	}
	l.Len++

	return nil
}

func (l *Double[E]) Remove(target E) error {
	n, err := l.Find(target)
	if err != nil {
		return err
	}

	if n.Prev == nil { // n is the start of the list
		l.First = n.Next
	} else {
		n.Prev.Next = n.Next
	}
	if n.Next == nil { // n is the end of the list
		l.Last = n.Prev
	} else {
		n.Next.Prev = n.Prev
	}
	l.Len--

	return nil
}

func (l *Double[E]) Find(elem E) (*DNode[E], error) {
	for cur := l.First; cur != nil; cur = cur.Next {
		if cur.Data == elem {
			return cur, nil
		}
	}

	return nil, fmt.Errorf("find node with data %v: %w", elem, ErrNotFound)
}

func (l *Double[E]) AtIndex(i int) (*DNode[E], error) {
	var cur *DNode[E]
	for j, cur := 0, l.First; j < i && cur != nil; cur = cur.Next {
	}
	if cur == nil {
		return nil, ErrOutOfBounds
	}
	return cur, nil
}

// DNode is a node in a doubly-linked list.
type DNode[E comparable] struct {
	Data E
	Next *DNode[E]
	Prev *DNode[E]
}
