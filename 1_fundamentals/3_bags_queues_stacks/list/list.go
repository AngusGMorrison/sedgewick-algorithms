package list

import (
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
