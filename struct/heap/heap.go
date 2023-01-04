package heap

import "fmt"

// Heap implements a priority queue, where the node at the root has greatest priority.
type Heap[E any] struct {
	data       []E
	priorityFn func(a, b E) bool
}

func NewHeap[E any](hasPriority func(a, b E) bool) *Heap[E] {
	return &Heap[E]{
		priorityFn: hasPriority,
	}
}

func NewHeapFromSlice[S ~[]E, E any](s S, hasPriority func(a, b E) bool) *Heap[E] {
	h := NewHeap(hasPriority)
	for _, e := range s {
		h.Push(e)
	}

	return h
}

func (mh *Heap[E]) Len() int {
	return len(mh.data)
}

func (mh *Heap[E]) IsEmpty() bool {
	return len(mh.data) == 0
}

func (h *Heap[E]) Push(elem E) {
	h.data = append(h.data, elem)
	h.swim()
}

// swim moves the last element in the queue up the heap until it reaches its correct position.
func (h *Heap[E]) swim() {
	for i := len(h.data) - 1; i > 0 && h.hasPriority(i, (i-1)/2); i = (i - 1) / 2 { // while the current element is less than its parent...
		h.swap(i, (i-1)/2)
	}
}

func (h *Heap[E]) Pop() (E, bool) {
	if h.IsEmpty() {
		return *new(E), false
	}

	head := h.data[0]
	h.data[0] = *new(E) // zero the value to avoid loitering
	h.swap(0, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.sink()

	return head, true
}

func (h *Heap[E]) sink() {
	// While i is has lower priority than its children, swap i with its greatest child.
	for i := 0; 2*i+1 < len(h.data); { // while i still has children...
		j := 2*i + 1                                    // initialize j to the first child of i
		if j < len(h.data)-1 && h.hasPriority(j+1, j) { // if there is another child and it has greater priority than the current value at j, select that child
			j++
		}

		if !h.hasPriority(j, i) { // i is has priority over or is equal to its greatest child, so it is in the correct position
			fmt.Printf("%v is greater than or equal to %v\n", h.data[j], h.data[i])
			break
		}
		fmt.Printf("%v is less than %v\n", h.data[j], h.data[i])

		h.swap(i, j) // i has lower priority than its highest-priority child, so swap it into position
		i = j
	}
}

func (h *Heap[E]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[E]) hasPriority(i, j int) bool {
	return h.priorityFn(h.data[i], h.data[j])
}
