package stack

import (
	"reflect"
	"testing"
)

func Test_SliceStack_Len(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		slice []int
	}{
		{
			name:  "stack is empty",
			slice: nil,
		},
		{
			name:  "stack is populated",
			slice: []int{1},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stack := &SliceStack[int]{slice: tc.slice}
			if gotLen := stack.Len(); gotLen != len(tc.slice) {
				t.Errorf("want len %d, got %d", len(tc.slice), gotLen)
			}
		})
	}
}

func Test_SliceStack_Push(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		slice []int
		elems []int
	}{
		{
			name:  "push to empty stack",
			slice: nil,
			elems: []int{1},
		},
		{
			name:  "push to populated stack",
			slice: []int{1},
			elems: []int{2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stack := NewSliceStack[int]()
			for _, elem := range tc.elems {
				stack.Push(elem)
			}

			for i, wantElem := range tc.elems {
				if gotElem := stack.slice[i]; gotElem != wantElem {
					t.Errorf("want slice position %d to have elem %d, got %d", i, wantElem, gotElem)
				}
			}
		})
	}
}

func Test_SliceStack_Pop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		slice     []int
		wantElem  int
		wantOK    bool
		wantSlice []int
	}{
		{
			name:      "stack is empty",
			slice:     nil,
			wantElem:  0,
			wantOK:    false,
			wantSlice: nil,
		},
		{
			name:      "stack has one element",
			slice:     []int{1},
			wantElem:  1,
			wantOK:    true,
			wantSlice: []int{},
		},
		{
			name:      "stack has multiple elements",
			slice:     []int{1, 2},
			wantElem:  2,
			wantOK:    true,
			wantSlice: []int{1},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stack := &SliceStack[int]{slice: tc.slice}
			gotElem, gotOK := stack.Pop()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if gotOK != tc.wantOK {
				t.Errorf("want ok %t, got %t", tc.wantOK, gotOK)
			}
			if !reflect.DeepEqual(stack.slice, tc.wantSlice) {
				t.Errorf("want slice\n\t%v,\ngot\n\t%v", tc.wantSlice, stack.slice)
			}
		})
	}
}

func Test_SliceStack_Peek(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		slice    []int
		wantElem int
		wantOK   bool
	}{
		{
			name:     "stack is empty",
			slice:    []int{},
			wantElem: 0,
			wantOK:   false,
		},
		{
			name:     "stack is populated",
			slice:    []int{1, 2},
			wantElem: 2,
			wantOK:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialSlice := make([]int, len(tc.slice))
			copy(initialSlice, tc.slice)
			stack := &SliceStack[int]{slice: tc.slice}

			gotElem, gotOK := stack.Peek()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if gotOK != tc.wantOK {
				t.Errorf("want OK %t, got %t", tc.wantOK, gotOK)
			}
			if !reflect.DeepEqual(stack.slice, initialSlice) {
				t.Errorf("want stack\n\t%v,\ngot\n\t%v", initialSlice, stack.slice)
			}
		})
	}
}

func Test_SliceStack_Each(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		slice          []int
		wantCalledWith []int
	}{
		{
			name:           "stack is empty",
			slice:          nil,
			wantCalledWith: nil,
		},
		{
			name:           "stack is populated",
			slice:          []int{1, 2},
			wantCalledWith: []int{2, 1},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := &callRecorder{}
			stack := &SliceStack[int]{slice: tc.slice}

			stack.Each(rec.recordEach)
			if !reflect.DeepEqual(tc.wantCalledWith, rec.calledWith) {
				t.Errorf(
					"want (*SliceStack).Each to have called the recorder with\n\t%v (len %d),\ngot\n\t%v (len %d)",
					tc.wantCalledWith, len(tc.wantCalledWith), rec.calledWith, len(rec.calledWith),
				)
			}
		})
	}
}

func Test_ListStack_Len(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		nodes []*node[int]
	}{
		{
			name:  "stack is empty",
			nodes: nil,
		},
		{
			name: "stack is populated",
			nodes: []*node[int]{
				{data: 1},
				{data: 2},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for i := 0; i < len(tc.nodes)-1; i++ {
				tc.nodes[i].next = tc.nodes[i+1]
			}

			var first *node[int]
			if len(tc.nodes) > 0 {
				first = tc.nodes[0]
			}

			stack := &ListStack[int]{
				len:   len(tc.nodes),
				first: first,
			}
			if gotLen := stack.Len(); gotLen != len(tc.nodes) {
				t.Errorf("want len %d, got %d", len(tc.nodes), gotLen)
			}
		})
	}
}

func Test_ListStack_Push(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		stack *ListStack[int]
		data  []int
	}{
		{
			name: "push single node",
			stack: &ListStack[int]{
				len:   0,
				first: nil,
			},
			data: []int{1},
		},
		{
			name: "push multiple nodes",
			stack: &ListStack[int]{
				len:   0,
				first: nil,
			},
			data: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialLen := tc.stack.len
			for _, elem := range tc.data {
				tc.stack.Push(elem)
			}

			wantLen := initialLen + len(tc.data)
			if tc.stack.len != wantLen {
				t.Errorf("want len %d, got %d", wantLen, tc.stack.len)
			}

			for i, cur := 1, tc.stack.first; cur != nil; i, cur = i+1, cur.next {
				wantData := tc.data[len(tc.data)-i]
				if cur.data != wantData {
					t.Errorf(
						"want node at position %d to have data %d, got %d",
						i, wantData, cur.data,
					)
				}
			}
		})
	}
}

func Test_ListStack_Pop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		stack     *ListStack[int]
		wantData  int
		wantOK    bool
		wantFirst *node[int]
		wantLen   int
	}{
		{
			name:      "empty stack",
			stack:     &ListStack[int]{},
			wantData:  0,
			wantOK:    false,
			wantFirst: nil,
			wantLen:   0,
		},
		{
			name: "one-node stack",
			stack: &ListStack[int]{
				len:   1,
				first: &node[int]{data: 1},
			},
			wantData:  1,
			wantOK:    true,
			wantFirst: nil,
			wantLen:   0,
		},
		{
			name: "multiple-node stack",
			stack: &ListStack[int]{
				len: 2,
				first: &node[int]{
					data: 1,
					next: &node[int]{
						data: 2,
					},
				},
			},
			wantData: 1,
			wantOK:   true,
			wantFirst: &node[int]{
				data: 2,
			},
			wantLen: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotData, gotOK := tc.stack.Pop()
			if gotData != tc.wantData {
				t.Errorf("want data %d, got %d", tc.wantData, gotData)
			}
			if gotOK != tc.wantOK {
				t.Errorf("want OK %t, got %t", tc.wantOK, gotOK)
			}
			if !reflect.DeepEqual(tc.stack.first, tc.wantFirst) {
				t.Errorf("want first node %+v, got %+v", tc.wantFirst, tc.stack.first)
			}
			if tc.stack.len != tc.wantLen {
				t.Errorf("want stack len %d, got %d", tc.wantLen, tc.stack.len)
			}
		})
	}
}

func Test_ListStack_Peek(t *testing.T) {
	t.Parallel()

	t.Run("stack is empty", func(t *testing.T) {
		t.Parallel()

		stack := &ListStack[int]{}
		gotElem, gotOK := stack.Peek()
		if gotElem != 0 {
			t.Errorf("want elem 0, got %d", gotElem)
		}
		if gotOK {
			t.Errorf("want OK false, got true")
		}
		if stack.first != nil {
			t.Errorf("want nil first element, got %+v", stack.first)
		}
		if stack.len != 0 {
			t.Errorf("want len 0, got %d", stack.len)
		}
	})

	t.Run("stack is populated", func(t *testing.T) {
		t.Parallel()

		second := &node[int]{data: 2}
		first := &node[int]{
			data: 1,
			next: second,
		}
		stack := &ListStack[int]{
			len:   2,
			first: &(*first),
		}

		gotElem, gotOK := stack.Peek()
		if gotElem != first.data {
			t.Errorf("want elem %d, got %d", first.data, gotElem)
		}
		if !gotOK {
			t.Errorf("want OK true, got false")
		}
		if !reflect.DeepEqual(stack.first, first) {
			t.Errorf("want first node %+v, got %+v", first, stack.first)
		}
		if stack.len != 2 {
			t.Errorf("want len 2, got %d", stack.len)
		}
	})
}

func Test_ListStack_Each(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		stack          *ListStack[int]
		wantCalledWith []int
	}{
		{
			name:           "empty stack",
			stack:          &ListStack[int]{},
			wantCalledWith: nil,
		},
		{
			name: "populated stack",
			stack: &ListStack[int]{
				len: 2,
				first: &node[int]{
					data: 1,
					next: &node[int]{
						data: 2,
					},
				},
			},
			wantCalledWith: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := &callRecorder{}
			tc.stack.Each(rec.recordEach)
			if !reflect.DeepEqual(tc.wantCalledWith, rec.calledWith) {
				t.Errorf(
					"want call recorder to be called with\n\t%v,\ngot\n\t%v",
					tc.wantCalledWith, rec.calledWith,
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
