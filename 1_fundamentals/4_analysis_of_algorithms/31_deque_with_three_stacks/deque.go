package deque

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

type deque[E any] struct {
	left, mid, right *stack[E]
}

func (d *deque[E]) len() int {
	return d.left.len() + d.mid.len() + d.right.len()
}

func (d *deque[E]) pushLeft(elem E) {
	d.left.push(elem)
}

func (d *deque[E]) pushRight(elem E) {
	d.right.push(elem)
}

func (d *deque[E]) popLeft() (E, bool) {
	if d.len() == 0 {
		return *new(E), false
	}
	if d.left.len() == 0 {
		d.rebalance(d.left, d.right)
	}

	return d.left.pop()
}

func (d *deque[E]) popRight() (E, bool) {
	if d.len() == 0 {
		return *new(E), false
	}
	if d.right.len() == 0 {
		d.rebalance(d.right, d.left)
	}

	return d.right.pop()
}

// rebalance takes one empty stack and one full stack and rebalances them so that each is half full.
func (d *deque[E]) rebalance(empty, full *stack[E]) {
	half := full.len() / 2
	for i := 0; i < half; i++ {
		elem, _ := full.pop()
		d.mid.push(elem)
	}
	for elem, ok := full.pop(); ok; elem, ok = full.pop() {
		empty.push(elem)
	}
	for elem, ok := d.mid.pop(); ok; elem, ok = d.mid.pop() {
		full.push(elem)
	}
}
