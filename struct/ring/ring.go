package ring

import "errors"

var (
	ErrFull  = errors.New("write to full ring buffer")
	ErrEmpty = errors.New("read from empty ring buffer")
)

type Ring[D any] struct {
	len, cap       int
	reader, writer int
	buffer         []D
}

func (r *Ring[D]) Len() int {
	return r.len
}

func (r *Ring[D]) IsFull() bool {
	return r.len == r.cap
}

func (r *Ring[D]) IsEmpty() bool {
	return r.len == 0
}

func (r *Ring[D]) Publish(data D) error {
	if r.IsFull() {
		return ErrFull
	}

	r.buffer[r.writer] = data
	r.writer = (r.writer + 1) % r.cap
	r.len++

	return nil
}

func (r *Ring[D]) Consume() (D, error) {
	if r.IsEmpty() {
		return *new(D), ErrEmpty
	}

	data := r.buffer[r.reader]
	r.reader = (r.reader + 1) % r.cap
	r.len--

	return data, nil
}

// ConcurrentRing allows blocking, thread-safe publish and consume operations on the ring using Go's
// concurrency primitives.
type ConcurrentRing[D any] struct {
	data chan D
}

func NewConcurrentRing[D any](cap int) *ConcurrentRing[D] {
	return &ConcurrentRing[D]{
		data: make(chan D, cap),
	}
}

func (r *ConcurrentRing[D]) Len() int {
	return len(r.data)
}

func (r *ConcurrentRing[D]) Publish(data D) {
	r.data <- data
}

func (r *ConcurrentRing[D]) Consume() D {
	return <-r.data
}
