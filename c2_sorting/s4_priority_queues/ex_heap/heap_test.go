package ex_heap

import (
	"reflect"
	"testing"
)

func Test_MaxHeap_Push(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "single value",
			input: []int{1},
			want:  []int{0, 1},
		},
		{
			name:  "multiple values in heap order",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{0, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:  "multiple unordered values",
			input: []int{1, 5, 2, 7, 4, 9, 8, 6, 3, 10},
			want:  []int{0, 10, 9, 8, 5, 6, 2, 7, 1, 3, 4},
		},
		{
			name:  "duplicate keys",
			input: []int{1, 5, 2, 7, 5, 9, 5, 5, 3, 10},
			want:  []int{0, 10, 9, 7, 5, 5, 2, 5, 1, 3, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			pq := NewMaxPQ(func(a, b int) bool {
				return a < b
			})
			for _, e := range tc.input {
				pq.Push(e)
			}

			if !reflect.DeepEqual(tc.want, pq.data) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.want, pq.data)
			}

		})
	}
}

func Test_MaxHeap_Pop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		initial  []int
		wantElem int
		wantOK   bool
		wantPQ   []int
	}{
		{
			name:     "empty heap",
			initial:  []int{0},
			wantElem: 0,
			wantOK:   false,
			wantPQ:   []int{0},
		},
		{
			name:     "single value",
			initial:  []int{0, 1},
			wantElem: 1,
			wantOK:   true,
			wantPQ:   []int{0},
		},
		{
			name:     "multiple values",
			initial:  []int{0, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			wantElem: 10,
			wantOK:   true,
			wantPQ:   []int{0, 9, 7, 8, 3, 6, 5, 4, 1, 2},
		},
		{
			name:     "duplicate keys",
			initial:  []int{0, 10, 9, 7, 5, 5, 2, 5, 1, 3, 5},
			wantElem: 10,
			wantOK:   true,
			wantPQ:   []int{0, 9, 5, 7, 5, 5, 2, 5, 1, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			pq := MaxHeap[int]{
				data: tc.initial,
				less: func(a, b int) bool {
					return a < b
				},
			}

			gotElem, gotOK := pq.Pop()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if gotOK != tc.wantOK {
				t.Errorf("want elem %t, got %t", tc.wantOK, gotOK)
			}
			if !reflect.DeepEqual(tc.wantPQ, pq.data) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.wantPQ, pq.data)
			}

		})
	}
}
