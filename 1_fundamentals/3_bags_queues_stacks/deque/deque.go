package deque

import "github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"

type SliceDeque[D any] struct {
	data []D
}

func (sd *SliceDeque[D]) IsEmpty() bool {
	return len(sd.data) == 0
}

func (sd *SliceDeque[D]) Len() int {
	return len(sd.data)
}

func (sd *SliceDeque[D]) PushLeft(data D) {
	// To avoid reallocating the whole array...
	sd.data = append(sd.data, *new(D)) // extend the slice by one empty element
	copy(sd.data[1:], sd.data)         // copy the existing data one position forward, overwriting the empty element
	sd.data[0] = data                  // insert the new data at the beginning
}

func (sd *SliceDeque[D]) PushRight(data D) {
	sd.data = append(sd.data, data)
}

func (sd *SliceDeque[D]) PopLeft() (D, bool) {
	if len(sd.data) == 0 {
		return *new(D), false
	}

	data := sd.data[0]
	sd.data = sd.data[1:]

	return data, true
}

func (sd *SliceDeque[D]) PopRight() (D, bool) {
	if len(sd.data) == 0 {
		return *new(D), false
	}

	data := sd.data[len(sd.data)-1]
	sd.data = sd.data[:len(sd.data)-1]

	return data, true
}

type ListDeque[D comparable] struct {
	list *list.DoubleList[D]
}

func NewListDeque[D comparable]() *ListDeque[D] {
	return &ListDeque[D]{
		list: &list.DoubleList[D]{},
	}
}

func (ld *ListDeque[D]) IsEmpty() bool {
	return ld.list.Len() == 0
}

func (ld *ListDeque[D]) Len() int {
	return ld.list.Len()
}

func (ld *ListDeque[D]) PushLeft(data D) {
	ld.list.Prepend(data)
}

func (ld *ListDeque[D]) PushRight(data D) {
	ld.list.Append(data)
}

func (ld *ListDeque[D]) PopLeft() (D, bool) {
	data, err := ld.list.Shift()
	if err != nil {
		return *new(D), false
	}

	return data, true
}

func (ld *ListDeque[D]) PopRight() (D, bool) {
	data, err := ld.list.Pop()
	if err != nil {
		return *new(D), false
	}

	return data, true
}
