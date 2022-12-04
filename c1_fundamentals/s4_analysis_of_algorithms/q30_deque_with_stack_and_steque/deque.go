package q30_deque_with_stack_and_steque

type stack[E any] struct {
	data []E
}

func (s *stack[E]) len() int {
	return len(s.data)
}

func (s *stack[E]) push(elem E) {
	s.data = append(s.data, elem)
}

func (s *stack[E]) pop() (elem E, ok bool) {
	if s.len() == 0 {
		return
	}

	elem, s.data = s.data[s.len()-1], s.data[:s.len()-1]
	return elem, true
}

type steque[E any] struct {
	data []E
}

func (s *steque[E]) len() int {
	return len(s.data)
}

func (s *steque[E]) enqueue(elem E) {
	s.data = append([]E{elem}, s.data...)
}

func (s *steque[E]) push(elem E) {
	s.data = append(s.data, elem)
}

func (s *steque[E]) pop() (elem E, ok bool) {
	if s.len() == 0 {
		return
	}

	elem, s.data = s.data[s.len()-1], s.data[:s.len()-1]
	return elem, true
}

type deque[E any] struct {
	stack  *stack[E]
	steque *steque[E]
}

func newDeque[E any]() *deque[E] {
	return &deque[E]{
		stack:  &stack[E]{},
		steque: &steque[E]{},
	}
}

func (d *deque[E]) len() int {
	return d.stack.len() + d.steque.len()
}

func (d *deque[E]) pushLeft(elem E) {
	d.steque.enqueue(elem)
	for higherElem, ok := d.stack.pop(); ok; higherElem, ok = d.stack.pop() {
		d.steque.push(higherElem)
	}
}

func (d *deque[E]) popLeft() (E, bool) {
	for lowerElem, ok := d.steque.pop(); ok; lowerElem, ok = d.steque.pop() {
		d.stack.push(lowerElem)
	}

	return d.stack.pop()
}

func (d *deque[E]) pushRight(elem E) {
	for lowerElem, ok := d.stack.pop(); ok; lowerElem, ok = d.stack.pop() {
		d.steque.push(lowerElem)
	}
	d.steque.push(elem)
}

func (d *deque[E]) popRight() (E, bool) {
	for lowerElem, ok := d.stack.pop(); ok; lowerElem, ok = d.stack.pop() {
		d.steque.push(lowerElem)
	}
	return d.steque.pop()
}
