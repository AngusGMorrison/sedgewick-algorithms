package steque

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

// enqueue pushes the element onto the bottom of the steque, which is the top of the enq stack. If
// the previous operation was dequeue, the front of the queue is at the top of deq, so all elements
// in deq must be popped and pushed onto enq before enqueuing the new element.
func (s *steque[E]) enqueue(elem E) {
	for priorElem, ok := s.top.pop(); ok; priorElem, ok = s.top.pop() {
		s.bottom.push(priorElem)
	}
	s.bottom.push(elem)
}

// push pushes the element onto the top of steque, which is the top of the deq stack. If the
// previous operation was enqueue, the back of the queue is at the top of enq, so all elements in
// enq must be popped and pushed onto deq before pushing the new element.
func (s *steque[E]) push(elem E) {
	for priorElem, ok := s.bottom.pop(); ok; priorElem, ok = s.bottom.pop() {
		s.top.push(priorElem)
	}
	s.top.push(elem)
}

func (s *steque[E]) pop() (E, bool) {
	for priorElem, ok := s.bottom.pop(); ok; priorElem, ok = s.bottom.pop() {
		s.top.push(priorElem)
	}
	return s.top.pop()
}
