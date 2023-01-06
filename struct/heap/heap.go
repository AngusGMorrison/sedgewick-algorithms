package heap

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
			break
		}

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

type Prioritizable[E any] interface {
	HasPriority(other E) bool
}

// IndexedEntry associates a heap entry with its index in the heap's underlying slice. This
// facilitates random deletions from the heap in constant time using (*IndexedHeap).Remove, since we
// don't have to search from the root to locate the order.
type IndexedEntry[E Prioritizable[E]] struct {
	index int
	entry E
}

func (i *IndexedEntry[E]) Index() int {
	return i.index
}

func (i *IndexedEntry[E]) Entry() E {
	return i.entry
}

// hasPriority satisfies prioritizable.
func (a *IndexedEntry[E]) HasPriority(b *IndexedEntry[E]) bool {
	return a.entry.HasPriority(b.entry)
}

type IndexedHeap[E Prioritizable[E]] struct {
	data []*IndexedEntry[E]
}

func NewIndexedHeap[E Prioritizable[E]]() *IndexedHeap[E] {
	return &IndexedHeap[E]{}
}

func (h *IndexedHeap[E]) Size() int {
	return len(h.data)
}

func (h *IndexedHeap[E]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *IndexedHeap[E]) Push(entry E) *IndexedEntry[E] {
	indexedEntry := &IndexedEntry[E]{
		index: len(h.data),
		entry: entry,
	}
	h.data = append(h.data, indexedEntry)
	h.swim(len(h.data) - 1)

	return indexedEntry
}

func (h *IndexedHeap[E]) Pop() (E, bool) {
	return h.removeIth(0)
}

func (h *IndexedHeap[E]) Peek() (E, bool) {
	if h.IsEmpty() {
		return *new(E), false
	}

	return h.data[0].entry, true
}

func (h *IndexedHeap[E]) Remove(i int) (E, bool) {
	return h.removeIth(i)
}

func (h *IndexedHeap[E]) removeIth(i int) (E, bool) {
	if i >= len(h.data) {
		return *new(E), false
	}

	removed := h.data[i]
	h.data[i] = nil // avoid loitering
	h.swap(i, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.sink(i)

	return removed.entry, true
}

func (h *IndexedHeap[E]) swim(i int) {
	for ; i > 0 && h.hasPriority(i, (i-1)/2); i = (i - 1) / 2 {
		h.swap(i, (i-1)/2)
	}
}

func (h *IndexedHeap[E]) sink(i int) {
	for 2*i+1 < len(h.data) {
		j := 2*i + 1
		if j < len(h.data)-1 && h.hasPriority(j+1, j) { // select i's highest-priority child
			j++
		}

		if !h.hasPriority(j, i) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

func (h *IndexedHeap[E]) hasPriority(i, j int) bool {
	return h.data[i].HasPriority(h.data[j])
}

func (h *IndexedHeap[E]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].index = i
	h.data[j].index = j
}
