package q30_deque_with_stack_and_steque

import (
	"reflect"
	"testing"
)

func Test_deque(t *testing.T) {
	t.Parallel()

	t.Run("push left", func(t *testing.T) {
		t.Parallel()

		d := newDeque[int]()
		input := []int{1, 2, 3}
		for _, elem := range input {
			d.pushLeft(elem)
		}

		wantStequeData := []int{3, 2, 1}
		var wantStackData []int
		if !reflect.DeepEqual(wantStequeData, d.steque.data) {
			t.Errorf("\nwant steque data\n\t%v\ngot\n\t%v", wantStequeData, d.steque.data)
		}
		if !reflect.DeepEqual(wantStackData, d.stack.data) {
			t.Errorf("\nwant stack data\n\t%v\ngot\n\t%v", wantStackData, d.stack.data)
		}
	})

	t.Run("pop left", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name       string
			stack      []int
			steque     []int
			wantElem   int
			wantOK     bool
			wantStack  []int
			wantSteque []int
		}{
			{
				name:       "when len == 0",
				stack:      nil,
				steque:     nil,
				wantElem:   0,
				wantOK:     false,
				wantStack:  nil,
				wantSteque: nil,
			},
			{
				name:       "when steque contains 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{},
				wantSteque: []int{},
			},
			{
				name:       "when steque contains 1 element and stack is not empty",
				stack:      []int{2},
				steque:     []int{1},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{2},
				wantSteque: []int{},
			},
			{
				name:       "when steque contains > 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1, 2},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{2},
				wantSteque: []int{},
			},
			{
				name:       "when steque contains > 1 element and stack is not empty",
				stack:      []int{3},
				steque:     []int{1, 2},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{3, 2},
				wantSteque: []int{},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				d := &deque[int]{
					stack:  &stack[int]{data: tc.stack},
					steque: &steque[int]{data: tc.steque},
				}
				gotElem, gotOK := d.popLeft()
				if gotElem != tc.wantElem {
					t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
				}
				if gotOK != tc.wantOK {
					t.Errorf("want OK %t, got %t", tc.wantOK, gotOK)
				}
				if !reflect.DeepEqual(tc.wantStack, d.stack.data) {
					t.Errorf("\nwant stack\n\t%v\ngot\n\t%v", tc.wantStack, d.stack.data)
				}
				if !reflect.DeepEqual(tc.wantSteque, d.steque.data) {
					t.Errorf("\nwant steque\n\t%v\ngot\n\t%v", tc.wantSteque, d.steque.data)
				}
			})
		}
	})

	t.Run("push right", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name       string
			stack      []int
			steque     []int
			elem       int
			wantStack  []int
			wantSteque []int
		}{
			{
				name:       "when len == 0",
				stack:      nil,
				steque:     nil,
				elem:       1,
				wantStack:  nil,
				wantSteque: []int{1},
			},
			{
				name:       "when steque contains 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1},
				elem:       2,
				wantStack:  nil,
				wantSteque: []int{1, 2},
			},
			{
				name:       "when steque contains 1 element and stack is not empty",
				stack:      []int{3, 2},
				steque:     []int{1},
				elem:       4,
				wantStack:  []int{},
				wantSteque: []int{1, 2, 3, 4},
			},
			{
				name:       "when steque contains > 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1, 2},
				elem:       3,
				wantStack:  nil,
				wantSteque: []int{1, 2, 3},
			},
			{
				name:       "when steque contains > 1 element and stack is not empty",
				stack:      []int{4, 3},
				steque:     []int{1, 2},
				elem:       5,
				wantStack:  []int{},
				wantSteque: []int{1, 2, 3, 4, 5},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				d := &deque[int]{
					stack:  &stack[int]{data: tc.stack},
					steque: &steque[int]{data: tc.steque},
				}
				d.pushRight(tc.elem)
				if !reflect.DeepEqual(tc.wantStack, d.stack.data) {
					t.Errorf("\nwant stack\n\t%v\ngot\n\t%v", tc.wantStack, d.stack.data)
				}
				if !reflect.DeepEqual(tc.wantSteque, d.steque.data) {
					t.Errorf("\nwant steque\n\t%v\ngot\n\t%v", tc.wantSteque, d.steque.data)
				}
			})
		}
	})

	t.Run("pop right", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name       string
			stack      []int
			steque     []int
			wantElem   int
			wantOK     bool
			wantStack  []int
			wantSteque []int
		}{
			{
				name:       "when len == 0",
				stack:      nil,
				steque:     nil,
				wantElem:   0,
				wantOK:     false,
				wantStack:  nil,
				wantSteque: nil,
			},
			{
				name:       "when steque contains 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1},
				wantElem:   1,
				wantOK:     true,
				wantStack:  nil,
				wantSteque: []int{},
			},
			{
				name:       "when steque contains 1 element and stack is not empty",
				stack:      []int{1, 2},
				steque:     []int{3},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{},
				wantSteque: []int{3, 2},
			},
			{
				name:       "when steque contains > 1 element and stack is empty",
				stack:      nil,
				steque:     []int{1, 2},
				wantElem:   2,
				wantOK:     true,
				wantStack:  nil,
				wantSteque: []int{1},
			},
			{
				name:       "when steque contains > 1 element and stack is not empty",
				stack:      []int{1, 2},
				steque:     []int{4, 3},
				wantElem:   1,
				wantOK:     true,
				wantStack:  []int{},
				wantSteque: []int{4, 3, 2},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				d := &deque[int]{
					stack:  &stack[int]{data: tc.stack},
					steque: &steque[int]{data: tc.steque},
				}
				gotElem, gotOK := d.popRight()
				if gotElem != tc.wantElem {
					t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
				}
				if gotOK != tc.wantOK {
					t.Errorf("want OK %t, got %t", tc.wantOK, gotOK)
				}
				if !reflect.DeepEqual(tc.wantStack, d.stack.data) {
					t.Errorf("\nwant stack\n\t%v\ngot\n\t%v", tc.wantStack, d.stack.data)
				}
				if !reflect.DeepEqual(tc.wantSteque, d.steque.data) {
					t.Errorf("\nwant steque\n\t%v\ngot\n\t%v", tc.wantSteque, d.steque.data)
				}
			})
		}
	})

	t.Run("intermixed operations result in the correct stack and steque", func(t *testing.T) {
		t.Parallel()

		d := newDeque[int]()
		d.pushLeft(1)
		d.pushLeft(2)
		d.pushRight(3)
		d.popLeft()
		d.pushLeft(4)
		d.popRight()
		d.pushRight(5)

		wantStack := []int{}
		wantSteque := []int{4, 1, 5}
		if !reflect.DeepEqual(wantStack, d.stack.data) {
			t.Errorf("\nwant stack\n\t%v\ngot\n\t%v", wantStack, d.stack.data)
		}
		if !reflect.DeepEqual(wantSteque, d.steque.data) {
			t.Errorf("\nwant steque\n\t%v\ngot\n\t%v", wantSteque, d.steque.data)
		}
	})
}
