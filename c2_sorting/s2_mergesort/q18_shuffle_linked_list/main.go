package main

import (
	"fmt"
	"math/rand"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func main() {
	src := rand.NewSource(12345)
	shuffler := NewShuffler[int](src)

	// Empty l
	l := (*list.Double[int])(nil)
	fmt.Printf("input list: %s\n", l)
	shuffler.ShuffleList(l)
	fmt.Printf("shuffled: %s\n", l)

	// Ordered list
	l = list.NewList[int]()
	for i := 0; i < 100; i++ {
		l.Append(i)
	}
	fmt.Printf("input list: %s\n", l)
	shuffler.ShuffleList(l)
	fmt.Printf("shuffled: %s\n", l)
}

type Shuffler[E comparable] struct {
	rng *rand.Rand
}

func NewShuffler[E comparable](src rand.Source) *Shuffler[E] {
	return &Shuffler[E]{rng: rand.New(src)}
}

// Assumes the length of the list is known and a we're allowed to use a doubly-linked list.
func (s *Shuffler[E]) ShuffleList(l *list.Double[E]) *list.Double[E] {
	if l == nil || l.Len <= 1 {
		return l
	}

	// Pick initial element.
	idx := s.rng.Intn(l.Len)
	head, left, right, err := SplitAtIndex(l, idx)
	if err != nil {
		panic(err) // should never happen
	}

	if left.Len == 0 && right.Len == 0 {
		return head
	}

	var next, second *list.Double[E]
	if s.rng.Intn(2) == 0 {
		if left.Len > 0 {
			next = s.ShuffleList(left)
			if right.Len > 0 {
				second = s.ShuffleList(right)
			}
		} else {
			next = s.ShuffleList(right)
		}
	} else {
		if right.Len > 0 {
			next = s.ShuffleList(right)
			if left.Len > 0 {
				second = s.ShuffleList(left)
			}
		} else {
			next = s.ShuffleList(left)
		}
	}

	if next != nil && next.Len > 0 {
		head.Last.Next = next.First
		next.First.Prev = head.Last
		head.Last = next.Last
		head.Len += next.Len
	}
	if second != nil && second.Len > 0 {
		head.Last.Next = second.First
		second.First.Prev = head.Last
		head.Last = second.Last
		head.Len += second.Len
	}

	fmt.Printf("%+v\n", head)
	return head
}

// SplitAtIndex splits a doubly linked list at index i, returning the node at i, nodes to the left
// of i and nodes to the right of i as three separate linked lists.
func SplitAtIndex[E comparable](l *list.Double[E], i int) (atIndex, left, right *list.Double[E], err error) {
	initialLen := l.Len
	initialLast := l.Last

	node, err := l.AtIndex(i) // delete head
	if err != nil {
		return nil, nil, nil, err
	}

	left = l
	left.Len = i
	left.Last = node.Prev
	if node.Prev == nil {
		left.First = nil
	} else {
		node.Prev.Next = nil
		node.Prev = nil
	}

	right = &list.Double[E]{
		Len:   initialLen - (i + 1),
		First: node.Next,
	}
	if node.Next != nil {
		right.Last = initialLast
		node.Next.Prev = nil
		node.Next = nil
	}

	atIndex = &list.Double[E]{
		Len:   1,
		First: node,
		Last:  node,
	}

	return atIndex, left, right, nil
}
