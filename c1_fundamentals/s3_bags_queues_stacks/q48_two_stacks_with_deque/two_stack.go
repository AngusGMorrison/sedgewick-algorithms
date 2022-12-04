package q48_two_stacks_with_deque

import "github.com/angusgmorrison/sedgewick_algorithms/struct/deque"

// TwoStack represents two stacks backed by a single deque.
type TwoStack[D comparable] struct {
	len1, len2 int
	deque      *deque.ListDeque[D]
}

func (ts *TwoStack[D]) LenStack1() int {
	return ts.len1
}

func (ts *TwoStack[D]) PushStack1(data D) {
	ts.deque.PushLeft(data)
	ts.len1++
}

func (ts *TwoStack[D]) PushStack2(data D) {
	ts.deque.PushRight(data)
	ts.len2++
}

func (ts *TwoStack[D]) PopStack1() (D, bool) {
	if ts.len1 == 0 {
		return *new(D), false
	}

	data, ok := ts.deque.PopLeft()
	if !ok {
		panic("(*TwoStack).deque was unexpectedly empty")
	}
	ts.len1--

	return data, true
}

func (ts *TwoStack[D]) PopStack2() (D, bool) {
	if ts.len2 == 0 {
		return *new(D), false
	}

	data, ok := ts.deque.PopRight()
	if !ok {
		panic("(*TwoStack).deque was unexpectedly empty")
	}
	ts.len2--

	return data, true
}
