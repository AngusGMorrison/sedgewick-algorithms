package queue

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"
)

func Test_Len(t *testing.T) {
	t.Parallel()

	wantLen := 13
	q := &Queue[int]{len: wantLen}
	if got := q.Len(); got != wantLen {
		t.Errorf("want %d, got %d", wantLen, got)
	}
}

func Test_Queue_Enqueue(t *testing.T) {
	t.Parallel()

	// Premade node values for convenience. Copy before use.
	node1 := list.Node[int]{Data: 1}
	node2 := list.Node[int]{Data: 2}
	node3 := list.Node[int]{Data: 3}

	t.Run("queue is empty", func(t *testing.T) {
		t.Parallel()

		q := &Queue[int]{}
		wantLast := node1
		wantLast.Next = &wantLast
		wantQ := &Queue[int]{
			len:  1,
			last: &wantLast,
		}

		q.Enqueue(node1.Data)
		if !reflect.DeepEqual(q, wantQ) {
			t.Errorf("\nwant queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})

	t.Run("queue has one element", func(t *testing.T) {
		t.Parallel()

		initialLast := node1
		initialLast.Next = &initialLast
		q := &Queue[int]{
			len:  1,
			last: &initialLast,
		}

		wantLast := node2
		wantFirst := node1
		wantFirst.Next = &wantLast
		wantLast.Next = &wantFirst
		wantQ := &Queue[int]{
			len:  2,
			last: &wantLast,
		}

		q.Enqueue(node2.Data)
		if !reflect.DeepEqual(q, wantQ) {
			fmt.Println("equal")
			t.Errorf("\nwant queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})

	t.Run("queue has more than one element", func(t *testing.T) {
		t.Parallel()

		initialFirst := node1
		initialLast := node2
		initialLast.Next = &initialFirst
		initialFirst.Next = &initialLast
		q := &Queue[int]{
			len:  2,
			last: &initialLast,
		}

		wantLast := node3
		wantSecond := node2
		wantSecond.Next = &wantLast
		wantFirst := node1
		wantFirst.Next = &wantSecond
		wantLast.Next = &wantFirst
		wantQ := &Queue[int]{
			len:  3,
			last: &wantLast,
		}

		q.Enqueue(node3.Data)
		if !reflect.DeepEqual(q, wantQ) {
			fmt.Println("equal")
			t.Errorf("\nwant queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})
}

func Test_Queue_Dequeue(t *testing.T) {
	t.Parallel()

	// Premade node values for convenience. Copy before use.
	node1 := list.Node[int]{Data: 1}
	node2 := list.Node[int]{Data: 2}
	node3 := list.Node[int]{Data: 3}

	t.Run("queue is empty", func(t *testing.T) {
		t.Parallel()

		q := &Queue[int]{}
		wantQ := &Queue[int]{}

		gotElem, gotErr := q.Dequeue()
		if gotElem != 0 {
			t.Errorf("want elem 0, got %d", gotElem)
		}
		if !errors.Is(gotErr, ErrQueueEmpty) {
			t.Errorf("want ErrQueueEmpty, got %v", gotErr)
		}
		if !reflect.DeepEqual(q, wantQ) {
			t.Errorf("want Queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})

	t.Run("queue has one element", func(t *testing.T) {
		t.Parallel()

		initialLast := node1
		initialLast.Next = &initialLast
		q := &Queue[int]{
			len:  1,
			last: &initialLast,
		}

		wantElem := node1.Data
		wantQ := &Queue[int]{}

		gotElem, gotErr := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotErr != nil {
			t.Errorf("want nil error, got %v", gotErr)
		}
		if !reflect.DeepEqual(q, wantQ) {
			t.Errorf("want Queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})

	t.Run("queue has two elements", func(t *testing.T) {
		t.Parallel()

		initialLast := node2
		initialFirst := node1
		initialFirst.Next = &initialLast
		initialLast.Next = &initialFirst
		q := &Queue[int]{
			len:  2,
			last: &initialLast,
		}

		wantElem := node1.Data
		wantLast := node2
		wantLast.Next = &wantLast
		wantQ := &Queue[int]{
			len:  1,
			last: &wantLast,
		}

		gotElem, gotErr := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotErr != nil {
			t.Errorf("want nil error, got %v", gotErr)
		}
		if !reflect.DeepEqual(q, wantQ) {
			t.Errorf("want Queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})

	t.Run("queue has more than two elements", func(t *testing.T) {
		t.Parallel()

		initialLast := node3
		initialSecond := node2
		initialSecond.Next = &initialLast
		initialFirst := node1
		initialFirst.Next = &initialSecond
		initialLast.Next = &initialFirst
		q := &Queue[int]{
			len:  3,
			last: &initialLast,
		}

		wantElem := node1.Data
		wantLast := node3
		wantFirst := node2
		wantFirst.Next = &wantLast
		wantLast.Next = &wantFirst
		wantQ := &Queue[int]{
			len:  2,
			last: &wantLast,
		}

		gotElem, gotErr := q.Dequeue()
		if gotElem != wantElem {
			t.Errorf("want elem %d, got %d", wantElem, gotElem)
		}
		if gotErr != nil {
			t.Errorf("want nil error, got %v", gotErr)
		}
		if !reflect.DeepEqual(q, wantQ) {
			t.Errorf("want Queue\n\t%s\ngot\n\t%s", wantQ, q)
		}
	})
}

func Test_Queue_Each(t *testing.T) {
	t.Parallel()

	t.Run("queue is empty", func(t *testing.T) {
		t.Parallel()

		q := &Queue[int]{}
		rec := &callRecorder[int]{}

		q.Each(rec.recordEach)
		if rec.calledWith != nil {
			t.Errorf("want rec.calledWith to be nil, got %v", rec.calledWith)
		}
	})

	t.Run("queue is not empty", func(t *testing.T) {
		t.Parallel()

		last := &list.Node[int]{Data: 2}
		first := &list.Node[int]{
			Data: 1,
			Next: last,
		}
		last.Next = first
		q := &Queue[int]{
			len:  2,
			last: last,
		}
		rec := &callRecorder[int]{}
		wantCalledWith := []int{first.Data, last.Data}

		q.Each(rec.recordEach)
		if !reflect.DeepEqual(rec.calledWith, wantCalledWith) {
			t.Errorf("\nwant rec.calledWith to eq\n\t%v\ngot\n\t%v", wantCalledWith, rec.calledWith)
		}
	})
}

type callRecorder[E comparable] struct {
	calledWith []E
}

func (cr *callRecorder[E]) recordEach(elem E) {
	cr.calledWith = append(cr.calledWith, elem)
}
