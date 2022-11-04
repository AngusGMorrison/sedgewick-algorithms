package list

import (
	"errors"
	"fmt"
	"strings"
)

// Node is a node in a singly linked list.
type Node[E comparable] struct {
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

var (
	ErrEmpty    = errors.New("list was empty")
	ErrNotFound = errors.New("element was not present in the list")
)

type DoubleList[E comparable] struct {
	len   int
	first *doubleNode[E]
	last  *doubleNode[E]
}

func NewList[E comparable]() *DoubleList[E] {
	return &DoubleList[E]{}
}

func (l *DoubleList[E]) Len() int {
	return l.len
}

func (l *DoubleList[E]) Prepend(elem E) {
	oldFirst := l.first
	l.first = &doubleNode[E]{
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

func (l *DoubleList[E]) Append(elem E) {
	oldLast := l.last
	l.last = &doubleNode[E]{
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

func (l *DoubleList[E]) Shift() (E, error) {
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

func (l *DoubleList[E]) Pop() (E, error) {
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

func (l *DoubleList[E]) Each(f func(elem E)) {
	for cur := l.first; cur != nil; cur = cur.next {
		f(cur.data)
	}
}

func (l *DoubleList[E]) String() string {
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

func (l *DoubleList[E]) InsertAfter(target, elem E) error {
	n, err := l.find(target)
	if err != nil {
		return err
	}

	oldNext := n.next
	n.next = &doubleNode[E]{
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

func (l *DoubleList[E]) InsertBefore(target, elem E) error {
	n, err := l.find(target)
	if err != nil {
		return err
	}

	oldPrev := n.prev
	n.prev = &doubleNode[E]{
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

func (l *DoubleList[E]) Remove(target E) error {
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

func (l *DoubleList[E]) find(elem E) (*doubleNode[E], error) {
	for cur := l.first; cur != nil; cur = cur.next {
		if cur.data == elem {
			return cur, nil
		}
	}

	return nil, fmt.Errorf("find node with data %v: %w", elem, ErrNotFound)
}

// doubleNode is a node in a doubly-linked list.
type doubleNode[E comparable] struct {
	data E
	next *doubleNode[E]
	prev *doubleNode[E]
}
