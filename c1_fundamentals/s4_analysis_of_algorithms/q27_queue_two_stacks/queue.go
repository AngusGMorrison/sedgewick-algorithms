package q27_queue_two_stacks

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

type queue[E any] struct {
	enq *stack[E] // the stack that data is enqueued to
	deq *stack[E] // the stack that data is dequeued from
}

func newQueue[E any]() *queue[E] {
	return &queue[E]{
		enq: &stack[E]{},
		deq: &stack[E]{},
	}
}

// enqueue pushes the element onto the end of the queue, which is the top of the enq stack. If the
// previous operation was dequeue, the front of the queue is at the top of deq, so all elements in
// deq must be popped and pushed onto enq before enqueuing the new element.
func (q *queue[E]) enqueue(elem E) {
	for priorElem, ok := q.deq.pop(); ok; priorElem, ok = q.deq.pop() {
		q.enq.push(priorElem)
	}
	q.enq.push(elem)
}

// dequeue pops the element from the front of the queue, which is the top of the deq stack. If the
// previous operation was enqueue, the front of the queue is at the bottom of enq, so all elements
// in enq must be popped and pushed onto deq before popping the top element.
func (q *queue[E]) dequeue() (E, bool) {
	for priorElem, ok := q.enq.pop(); ok; priorElem, ok = q.enq.pop() {
		q.deq.push(priorElem)
	}

	return q.deq.pop()
}
