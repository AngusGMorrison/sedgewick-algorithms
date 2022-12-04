package q28_stack_queue

type queue[E any] struct {
	data []E
}

func (q *queue[E]) len() int {
	return len(q.data)
}

func (q *queue[E]) enqueue(elem E) {
	q.data = append(q.data, elem)
}

func (q *queue[E]) dequeue() (elem E, ok bool) {
	if len(q.data) == 0 {
		return
	}

	elem, q.data = q.data[0], q.data[1:]
	return elem, true
}

type stack[E any] struct {
	q *queue[E]
}

func newStack[E any]() *stack[E] {
	return &stack[E]{
		q: &queue[E]{},
	}
}

func (s *stack[E]) push(elem E) {
	s.q.enqueue(elem) // push item onto the end of the queue
}

func (s *stack[E]) pop() (E, bool) {
	// Get the item at the end of the queue by cycling existing items to the back of the queue in
	// order.
	for i := 0; i < s.q.len()-1; i++ {
		elem, _ := s.q.dequeue()
		s.q.enqueue(elem)
	}

	return s.q.dequeue() // previously the back of the queue
}
