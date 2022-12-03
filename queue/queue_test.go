package queue

import (
	"reflect"
	"testing"
)

func Test_SliceQueue_Len(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		slice []int
	}{
		{
			name:  "queue is empty",
			slice: nil,
		},
		{
			name:  "queue is not empty",
			slice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			q := &SliceQueue[int]{slice: tc.slice}
			if got := q.Len(); got != len(tc.slice) {
				t.Errorf("want %d, got %d", len(tc.slice), got)
			}
		})
	}
}

func Test_SliceQueue_Enqueue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		slice []int
		elems []int
	}{
		{
			name:  "queue is empty",
			slice: nil,
			elems: []int{1},
		},
		{
			name:  "queue is populated",
			slice: []int{1},
			elems: []int{2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			q := &SliceQueue[int]{slice: tc.slice}
			for _, e := range tc.elems {
				q.Enqueue(e)
			}

			wantSlice := append(tc.slice, tc.elems...)
			if !reflect.DeepEqual(wantSlice, q.slice) {
				t.Errorf("want queue\n\t%v,\ngot\n\t%v", wantSlice, q.slice)
			}
		})
	}
}

func Test_SliceQueue_Dequeue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		slice     []int
		wantElem  int
		wantOK    bool
		wantSlice []int
	}{
		{
			name:      "queue is empty",
			slice:     nil,
			wantElem:  0,
			wantOK:    false,
			wantSlice: nil,
		},
		{
			name:      "queue has one element",
			slice:     []int{1},
			wantElem:  1,
			wantOK:    true,
			wantSlice: []int{},
		},
		{
			name:      "queue has multiple elements",
			slice:     []int{1, 2},
			wantElem:  1,
			wantOK:    true,
			wantSlice: []int{2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			q := &SliceQueue[int]{slice: tc.slice}
			gotElem, gotOK := q.Dequeue()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if gotOK != tc.wantOK {
				t.Errorf("want OK %t, got %t", tc.wantOK, gotOK)
			}
			if !reflect.DeepEqual(q.slice, tc.wantSlice) {
				t.Errorf("want queue\n\t%v,\ngot\n\t%v", tc.wantSlice, q.slice)
			}
		})
	}
}

func Test_SliceQueue_Each(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		slice []int
	}{
		{
			name:  "empty queue",
			slice: nil,
		},
		{
			name:  "populated queue",
			slice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			q := &SliceQueue[int]{slice: tc.slice}
			rec := &callRecorder{}
			q.Each(rec.recordEach)
			if !reflect.DeepEqual(rec.calledWith, tc.slice) {
				t.Errorf(
					"want callRecorder to be called with\n\t%v,\ngot\n\t%v",
					rec.calledWith, tc.slice,
				)
			}
		})
	}
}

type callRecorder struct {
	calledWith []int
}

func (c *callRecorder) recordEach(elem int) {
	c.calledWith = append(c.calledWith, elem)
}

func Test_ListQueue_Len(t *testing.T) {
	t.Parallel()

	wantLen := 1
	q := &ListQueue[int]{len: wantLen}
	if q.Len() != wantLen {
		t.Errorf("want len %d, got %d", wantLen, q.Len())
	}
}

func Test_ListQueue_Enqueue(t *testing.T) {
	t.Parallel()

	t.Run("queue is empty", func(t *testing.T) {
		t.Parallel()

		var (
			q         = &ListQueue[int]{}
			elem      = 1
			wantFirst = &node[int]{data: elem}
			wantLast  = wantFirst
			wantLen   = 1
		)

		q.Enqueue(elem)
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})

	t.Run("queue is populated", func(t *testing.T) {
		t.Parallel()

		var (
			qNode = &node[int]{data: 1}
			q     = &ListQueue[int]{
				len:   1,
				first: qNode,
				last:  qNode,
			}
			elem      = 2
			wantFirst = qNode
			wantLast  = &node[int]{data: elem}
			wantLen   = 2
		)

		q.Enqueue(elem)
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})
}

func Test_ListQueue_Dequeue(t *testing.T) {
	t.Parallel()

	t.Run("queue is empty", func(t *testing.T) {
		t.Parallel()

		var (
			q         = &ListQueue[int]{}
			wantElem  = 0
			wantOK    = false
			wantFirst = (*node[int])(nil)
			wantLast  = wantFirst
			wantLen   = 0
		)

		gotElem, gotOK := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotOK != wantOK {
			t.Errorf("want OK %t, got %t", wantOK, gotOK)
		}
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})

	t.Run("queue has one element", func(t *testing.T) {
		t.Parallel()

		var (
			qNode = &node[int]{data: 1}
			q     = &ListQueue[int]{
				len:   1,
				first: qNode,
				last:  qNode,
			}
			wantElem  = 1
			wantOK    = true
			wantFirst = (*node[int])(nil)
			wantLast  = wantFirst
			wantLen   = 0
		)

		gotElem, gotOK := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotOK != wantOK {
			t.Errorf("want OK %t, got %t", wantOK, gotOK)
		}
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})

	t.Run("queue has two elements", func(t *testing.T) {
		t.Parallel()

		var (
			qLast  = &node[int]{data: 2}
			qFirst = &node[int]{
				data: 1,
				next: qLast,
			}
			q = &ListQueue[int]{
				len:   2,
				first: qFirst,
				last:  qLast,
			}
			wantElem  = 1
			wantOK    = true
			wantFirst = qLast
			wantLast  = qLast
			wantLen   = 1
		)

		gotElem, gotOK := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotOK != wantOK {
			t.Errorf("want OK %t, got %t", wantOK, gotOK)
		}
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})

	t.Run("queue has more than two elements", func(t *testing.T) {
		t.Parallel()

		var (
			qLast = &node[int]{data: 3}
			qMid  = &node[int]{
				data: 2,
				next: qLast,
			}
			qFirst = &node[int]{
				data: 1,
				next: qMid,
			}
			q = &ListQueue[int]{
				len:   3,
				first: qFirst,
				last:  qLast,
			}
			wantElem  = 1
			wantOK    = true
			wantFirst = qMid
			wantLast  = qLast
			wantLen   = 2
		)

		gotElem, gotOK := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotOK != wantOK {
			t.Errorf("want OK %t, got %t", wantOK, gotOK)
		}
		if !reflect.DeepEqual(q.first, wantFirst) {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", wantFirst, q.first)
		}
		if !reflect.DeepEqual(q.last, wantLast) {
			t.Errorf("want last node\n\t%+v,\ngot\n\t%+v", wantLast, q.last)
		}
		if q.len != wantLen {
			t.Errorf("want len %d, got %d", wantLen, q.len)
		}
	})
}

func Test_ListQueue_Each(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		elems          []int
		wantCalledWith []int
	}{
		{
			name:  "queue is empty",
			elems: nil,
		},
		{
			name:  "queue is populated",
			elems: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := &callRecorder{}
			q := &ListQueue[int]{}
			for _, elem := range tc.elems {
				q.Enqueue(elem)
			}

			q.Each(rec.recordEach)
			if !reflect.DeepEqual(rec.calledWith, tc.elems) {
				t.Errorf(
					"want call recorder to be called with\n\t%v,\ngot\n\t%v",
					tc.elems, rec.calledWith,
				)
			}
		})
	}
}
