package q12_copy_stack

import (
	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

func CopySliceStack[E any](s *stack.SliceStack[E]) *stack.SliceStack[E] {
	elems := make([]E, 0, s.Len())
	s.Each(func(elem E) {
		elems = append(elems, elem)
	})

	cp := stack.NewSliceStack[E]()
	for i := s.Len() - 1; i >= 0; i-- {
		cp.Push(elems[i])
	}

	return cp
}

func CopyListStack[E any](s *stack.ListStack[E]) *stack.ListStack[E] {
	elems := make([]E, 0, s.Len())
	s.Each(func(elem E) {
		elems = append(elems, elem)
	})

	cp := stack.NewListStack[E]()
	for i := s.Len() - 1; i >= 0; i-- {
		cp.Push(elems[i])
	}

	return cp
}
