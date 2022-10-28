package stack

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

func (s *SliceStack[E]) Each(f func(elem E)) {
	for _, elem := range s.slice {
		f(elem)
	}
}

// ListStack implements a generic stack backed by a linked list.
type ListStack[E any] struct {
	len   int
	first *node[E]
}

func NewListStack[E any]() *ListStack[E] {
	return &ListStack[E]{}
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
}

func (ls *ListStack[E]) Pop() (E, bool) {
	if ls.len == 0 {
		return *new(E), false
	}

	elem := ls.first.data
	ls.first = ls.first.next
	ls.len--

	return elem, true
}

func (ls *ListStack[E]) Each(f func(elem E)) {
	for cur := ls.first; cur != nil; cur = cur.next {
		f(cur.data)
	}
}

type node[E any] struct {
	data E
	next *node[E]
}
