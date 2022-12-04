package main

import (
	"fmt"
	"math/rand"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func main() {
	src := rand.NewSource(12345)
	shuffler := NewListShuffler[int](src)

	// Empty l
	l := (*list.Double[int])(nil)
	fmt.Printf("top-down input list: %s\n", l)
	result := shuffler.BottomUp(l)
	fmt.Printf("top-down shuffled: %s\n", result)

	// Ordered list
	l = list.NewDouble[int]()
	for i := 0; i < 100; i++ {
		l.Append(i)
	}
	fmt.Printf("top-down input list: %s\n", l)
	result = shuffler.BottomUp(l)
	fmt.Printf("top-down shuffled: %s\n", result)

	// Empty l
	l = (*list.Double[int])(nil)
	fmt.Printf("bottom-up input list: %s\n", l)
	result = shuffler.TopDown(l)
	fmt.Printf("bottom-up shuffled: %s\n", result)

	// Ordered list
	l = list.NewDouble[int]()
	for i := 0; i < 100; i++ {
		l.Append(i)
	}
	fmt.Printf("bottom-up input list: %s\n", l)
	result = shuffler.TopDown(l)
	fmt.Printf("bottom-up shuffled: %s\n", result)
}

type ListShuffler[E comparable] struct {
	rng *rand.Rand
}

func NewListShuffler[E comparable](src rand.Source) *ListShuffler[E] {
	return &ListShuffler[E]{rng: rand.New(src)}
}

// BottomUp shuffles the list in bottom-up style by recursively breaking it into halves, then
// building it back up by randomly selecting elements from its two shuffled halves until both halves
// are depleted.
func (ls *ListShuffler[E]) BottomUp(l *list.Double[E]) *list.Double[E] {
	if l == nil || l.Len <= 1 {
		return l
	}

	mid := (l.Len / 2) - 1
	l1, l2, err := splitAtIndex(l, mid)
	if err != nil {
		panic(err) // should never happen
	}

	l1 = ls.BottomUp(l1)
	l2 = ls.BottomUp(l2)

	result := list.NewDouble[E]()
	l1InitialLen, l2InitialLen := l1.Len, l2.Len
	for i, j := 0, 0; i < l1InitialLen || j < l2InitialLen; {
		if ls.takeFromTargetWithLen(l1.Len, l2.Len) {
			elem, _ := l1.Pop()
			result.Append(elem)
			i++
		} else {
			elem, _ := l2.Pop()
			result.Append(elem)
			j++
		}
	}

	return result
}

// TopDown creates progressively more disordered lists of smaller and smaller sizes, then combines
// them into a single, randomly distributed list.
func (ls *ListShuffler[E]) TopDown(l *list.Double[E]) *list.Double[E] {
	if l == nil || l.Len <= 1 {
		return l
	}

	mid := (l.Len / 2) - 1
	initialLen := l.Len
	l1, l2 := list.NewDouble[E](), list.NewDouble[E]()
	for i, j := 0, mid+1; i <= mid || j < initialLen; {
		elem, err := l.Pop()
		if err != nil {
			panic(err) // should never happen
		}

		if ls.takeFromTargetWithLen(mid-i+1, initialLen-(j+1)) {
			l1.Append(elem)
			i++
		} else {
			l2.Append(elem)
			j++
		}
	}

	l1 = ls.TopDown(l1)
	l2 = ls.TopDown(l2)
	return l1.Join(l2)
}

// takeFromTargetWithLen returns true with probability equal to targetLen as a percentage of the
// total elements remaining in two lists. This allows us to write conditionals such that the
// probability of selecting any element is uniformly distributed.
func (ls *ListShuffler[E]) takeFromTargetWithLen(targetLen, otherLen int) bool {
	totalLen := float64(targetLen + otherLen)
	return ls.rng.Float64() < float64(targetLen)/totalLen
}

// splitAtIndex splits a doubly linked list at index i into two lists. The first starts at l.First
// and ends at i. The second starts at i+1 and ends at l.Last.
func splitAtIndex[E comparable](l *list.Double[E], i int) (*list.Double[E], *list.Double[E], error) {
	initialLen := l.Len

	ithNode, err := l.AtIndex(i)
	if err != nil {
		return nil, nil, err
	}

	l1 := &list.Double[E]{
		Len:   i + 1,
		First: l.First,
		Last:  ithNode,
	}

	len2 := initialLen - (i + 1)
	l2 := &list.Double[E]{
		Len:   len2,
		First: ithNode.Next,
	}

	if ithNode.Next != nil {
		l2.Last = l.Last
		ithNode.Next.Prev = nil
		ithNode.Next = nil
	}

	return l1, l2, nil
}
