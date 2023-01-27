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

type SymbolHeap[K comparable, V Prioritizable[V]] struct {
	data    []V
	indices map[K]int // maps a key K to its index in data
	keys    []K       // maps an index in data to a key
}

func NewSymbolHeap[K comparable, V Prioritizable[V]]() *SymbolHeap[K, V] {
	return &SymbolHeap[K, V]{
		indices: make(map[K]int),
	}
}

func (sh *SymbolHeap[K, V]) Size() int {
	return len(sh.data)
}

func (sh *SymbolHeap[K, V]) IsEmpty() bool {
	return sh.Size() == 0
}

func (sh *SymbolHeap[K, V]) Contains(key K) bool {
	_, ok := sh.indices[key]
	return ok
}

// Push pushes val onto the heap and returns its index. If a value corresponding to key already
// exists on the heap, the value is updated.
func (sh *SymbolHeap[K, V]) Push(key K, val V) {
	if sh.Contains(key) {
		sh.Update(key, val)
		return
	}

	sh.keys = append(sh.keys, key)
	sh.data = append(sh.data, val)
	lastIdx := len(sh.data) - 1
	sh.indices[key] = lastIdx
	sh.swim(lastIdx)
}

func (sh *SymbolHeap[K, V]) Pop() (K, V, bool) {
	if sh.Size() == 0 {
		return *new(K), *new(V), false
	}

	val := sh.data[0]
	key := sh.keys[0]
	sh.delete(0)

	return key, val, true
}

// Delete removes the value associated with key and restores the heap. Deleting a key that is not in
// the heap has no effect.
func (sh *SymbolHeap[K, V]) Delete(key K) (V, bool) {
	if !sh.Contains(key) {
		return *new(V), false
	}

	idx := sh.indices[key]
	val := sh.data[idx]
	sh.delete(idx)

	return val, true
}

// Update updates the current value at key and restores the heap. If no value exists for the key, it
// is inserted using Push.
func (sh *SymbolHeap[K, V]) Update(key K, val V) {
	if !sh.Contains(key) {
		sh.Push(key, val)
		return
	}

	idx := sh.indices[key]
	sh.data[idx] = val
	sh.swim(idx) // if the new val has priority over the old one, it will rise
	sh.sink(idx) // otherwise it will sink
}

// delete removes the key and value corresponding to the given heap index and restores the heap.
func (sh *SymbolHeap[K, V]) delete(idx int) {
	delete(sh.indices, sh.keys[idx])
	last := len(sh.data) - 1
	sh.swap(idx, last)

	// Truncate the slices without loitering.
	sh.data[last], sh.data = *new(V), sh.data[:last]
	sh.keys[last], sh.keys = *new(K), sh.keys[:last]

	// Restore the heap.
	sh.sink(idx)
}

func (sh *SymbolHeap[K, V]) swim(idx int) {
	for ; idx > 0 && sh.hasPriority(idx, (idx-1)/2); idx = (idx - 1) / 2 {
		sh.swap(idx, (idx-1)/2)
	}
}

func (sh *SymbolHeap[K, V]) sink(idx int) {
	for idx*2+1 < len(sh.data) {
		maxPriorityChild := idx*2 + 1
		if maxPriorityChild < len(sh.data)-1 && sh.hasPriority(maxPriorityChild+1, maxPriorityChild) {
			maxPriorityChild++
		}

		if !sh.hasPriority(maxPriorityChild, idx) {
			break
		}

		sh.swap(idx, maxPriorityChild)
		idx = maxPriorityChild
	}
}

func (sh *SymbolHeap[K, V]) hasPriority(i, j int) bool {
	return sh.data[i].HasPriority(sh.data[j])
}

func (sh *SymbolHeap[K, V]) swap(i, j int) {
	iKey, jKey := sh.keys[i], sh.keys[j]
	sh.data[i], sh.data[j] = sh.data[j], sh.data[i]
	sh.keys[i], sh.keys[j] = sh.keys[j], sh.keys[i]
	sh.indices[iKey], sh.indices[jKey] = j, i
}
