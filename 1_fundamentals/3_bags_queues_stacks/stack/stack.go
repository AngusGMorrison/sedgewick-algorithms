package stack

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
