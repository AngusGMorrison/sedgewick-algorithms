package stack

import (
	"reflect"
	"testing"
)

func Test_SliceStack_Len(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		slice   []int
		wantLen int
	}{
		{
			name:    "stack is empty",
			slice:   nil,
			wantLen: 0,
		},
		{
			name:    "stack is populated",
			slice:   []int{1},
			wantLen: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stack := &SliceStack[int]{slice: tc.slice}
			if gotLen := stack.Len(); gotLen != tc.wantLen {
				t.Errorf("want len %d, got %d", tc.wantLen, gotLen)
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

func Test_Stack_Each(t *testing.T) {
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
			slice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := &callRecorder{}
			stack := &SliceStack[int]{slice: tc.slice}

			stack.Each(rec.recordEach)
			if !reflect.DeepEqual(tc.slice, rec.calledWith) {
				t.Errorf(
					"want (*SliceStack).Each to have called the recorder with\n\t%v,\ngot\n\t%v",
					tc.slice, rec.calledWith,
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
