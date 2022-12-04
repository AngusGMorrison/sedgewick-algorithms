package q29_steque_two_stacks

type stack[E any] struct {
	data []E
}

func (s *stack[E]) len() int {
	return len(s.data)
}

func (s *stack[E]) push(elem E) {
	s.data = append(s.data, elem)
}

func (s *stack[E]) pop() (E, bool) {
	if len(s.data) == 0 {
		return *new(E), false
	}

	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, true
}

type steque[E any] struct {
	bottom, top *stack[E]
}

func newSteque[E any]() *steque[E] {
	return &steque[E]{
		bottom: &stack[E]{},
		top:    &stack[E]{},
	}
}

func (s *steque[E]) enqueue(elem E) {
	s.bottom.push(elem)
}

func (s *steque[E]) push(elem E) {
	s.top.push(elem)
}

func (s *steque[E]) pop() (E, bool) {
	if s.top.len() == 0 {
		for priorElem, ok := s.bottom.pop(); ok; priorElem, ok = s.bottom.pop() {
			s.top.push(priorElem)
		}
	}
	return s.top.pop()
}
