package ex_heap

// MaxHeap implements a maximum priority queue, where the node at the root has greatest priority.
type MaxHeap[E any] struct {
	data []E
	less func(a, b E) bool
}

func NewMaxPQ[E any](less func(a, b E) bool) *MaxHeap[E] {
	return &MaxHeap[E]{
		data: []E{*new(E)}, // pq is initialized with a dummy value at the zero-index so we can index from 1
		less: less,
	}
}

func (mh *MaxHeap[E]) Len() int {
	return len(mh.data) - 1 // pq is indexed from 1
}

func (mh *MaxHeap[E]) IsEmpty() bool {
	return mh.Len() == 0
}

func (mh *MaxHeap[E]) Push(elem E) {
	mh.data = append(mh.data, elem)
	mh.swim()
}

// swim moves the last element in the queue up the heap until it reaches its correct position.
func (mh *MaxHeap[E]) swim() {
	for k := mh.Len(); k > 1 && mh.less(mh.data[k/2], mh.data[k]); k /= 2 { // while the current element is less than its parent...
		mh.swap(k, k/2)
	}
}

func (mh *MaxHeap[E]) Pop() (E, bool) {
	if mh.IsEmpty() {
		return *new(E), false
	}

	max := mh.data[1]
	len := mh.Len()
	mh.swap(1, len)
	mh.data[len] = *new(E) // zero the value to avoid loitering
	mh.data = mh.data[:len]
	mh.sink()
	return max, true
}

func (mh *MaxHeap[E]) sink() {
	len := mh.Len()
	// While k is smaller than its children, swap k with its greatest child.
	for k := 1; 2*k <= len; { // while k still has children...
		j := 2 * k                                        // initialize j to the first child of k
		if j < len && mh.less(mh.data[j], mh.data[j+1]) { // if there is another child and it is greater than the current value at j, select that child
			j++
		}

		if !mh.less(mh.data[k], mh.data[j]) { // k is greater than or equal to its greatest child, so it is in the correct position
			break
		}

		mh.swap(k, j) // k is less than its greatest child, so swap it into position
		k = j
	}
}

func (mh *MaxHeap[E]) swap(i, j int) {
	mh.data[i], mh.data[j] = mh.data[j], mh.data[i]
}
