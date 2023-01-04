package heap

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
			want:  []int{1},
		},
		{
			name:  "multiple values in heap order",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:  "multiple unordered values",
			input: []int{1, 5, 2, 7, 4, 9, 8, 6, 3, 10},
			want:  []int{10, 9, 8, 5, 6, 2, 7, 1, 3, 4},
		},
		{
			name:  "duplicate keys",
			input: []int{1, 5, 2, 7, 5, 9, 5, 5, 3, 10},
			want:  []int{10, 9, 7, 5, 5, 2, 5, 1, 3, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			pq := NewHeap(func(a, b int) bool {
				return a > b
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
		data     []int
		wantElem int
		wantOK   bool
		wantPQ   []int
	}{
		{
			name:     "empty heap",
			data:     nil,
			wantElem: 0,
			wantOK:   false,
			wantPQ:   nil,
		},
		{
			name:     "single value",
			data:     []int{1},
			wantElem: 1,
			wantOK:   true,
			wantPQ:   []int{},
		},
		{
			name:     "multiple values",
			data:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			wantElem: 10,
			wantOK:   true,
			wantPQ:   []int{9, 7, 8, 3, 6, 5, 4, 1, 2},
		},
		{
			name:     "duplicate keys",
			data:     []int{10, 9, 7, 5, 5, 2, 5, 1, 3, 5},
			wantElem: 10,
			wantOK:   true,
			wantPQ:   []int{9, 5, 7, 5, 5, 2, 5, 1, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			pq := NewHeapFromSlice(
				tc.data,
				func(a, b int) bool {
					return a > b
				},
			)

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
