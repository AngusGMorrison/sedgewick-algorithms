package list

import (
	"errors"
	"reflect"
	"testing"
)

func Test_List_Prepend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		elem           int
		newWantList    func() *List[int]
	}{
		{
			name: "list is empty",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			elem: 1,
			newWantList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
		},
		{
			name: "list has one element",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			elem: 2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n2.next = n1
				n1.prev = n2
				return &List[int]{
					len:   2,
					first: n2,
					last:  n1,
				}
			},
		},
		{
			name: "list has more than one element",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			elem: 3,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				n3 := newNode[int](3)
				n3.next = n1
				n1.prev = n3
				return &List[int]{
					len:   3,
					first: n3,
					last:  n2,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			initialList.Prepend(tc.elem)
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("want list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_Append(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		elem           int
		newWantList    func() *List[int]
	}{
		{
			name: "list is empty",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			elem: 1,
			newWantList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
		},
		{
			name: "list has one element",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			elem: 2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
		},
		{
			name: "list has more than one element",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			elem: 3,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				n3 := newNode[int](3)
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			initialList.Append(tc.elem)
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("want list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_Shift(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		wantElem       int
		wantErr        error
		newWantList    func() *List[int]
	}{
		{
			name: "list is empty",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			wantElem: 0,
			wantErr:  ErrEmpty,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
		},
		{
			name: "list has one element",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
		},
		{
			name: "list has two elements",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *List[int] {
				n := newNode[int](2)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
		},
		{
			name: "list has more than two elements",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *List[int] {
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   2,
					first: n2,
					last:  n3,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			gotElem, gotErr := initialList.Shift()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("want ErrEmpty, got %v", gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("want list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_Pop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		wantElem       int
		wantErr        error
		newWantList    func() *List[int]
	}{
		{
			name: "list is empty",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			wantElem: 0,
			wantErr:  ErrEmpty,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
		},
		{
			name: "list has one element",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
		},
		{
			name: "list has two elements",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			wantElem: 2,
			wantErr:  nil,
			newWantList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
		},
		{
			name: "list has more than two elements",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantElem: 3,
			wantErr:  nil,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			gotElem, gotErr := initialList.Pop()
			if gotElem != tc.wantElem {
				t.Errorf("want elem %d, got %d", tc.wantElem, gotElem)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("want ErrEmpty, got %v", gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("want list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_InsertBefore(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		target         int
		elem           int
		newWantList    func() *List[int]
		wantErr        error
	}{
		{
			name: "target is not present in the list",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			target: 2,
			elem:   1,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
			wantErr: ErrNotFound,
		},
		{
			name: "target is the only element in the list",
			newInitialList: func() *List[int] {
				n := newNode[int](2)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			target: 2,
			elem:   1,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the first element in the list",
			newInitialList: func() *List[int] {
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   2,
					first: n2,
					last:  n3,
				}
			},
			target: 2,
			elem:   1,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the last element in the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n3 := newNode[int](3)
				n1.next = n3
				n3.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n3,
				}
			},
			target: 3,
			elem:   2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is in the middle of the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n3 := newNode[int](3)
				n4 := newNode[int](4)
				n1.next = n3
				n3.prev = n1
				n3.next = n4
				n4.prev = n3
				return &List[int]{
					len:   3,
					first: n1,
					last:  n4,
				}
			},
			target: 3,
			elem:   2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n4 := newNode[int](4)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				n3.next = n4
				n4.prev = n3
				return &List[int]{
					len:   4,
					first: n1,
					last:  n4,
				}
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			if gotErr := initialList.InsertBefore(tc.target, tc.elem); !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("want ErrNotFound, got %v", gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("\nwant list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_InsertAfter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		target         int
		elem           int
		newWantList    func() *List[int]
		wantErr        error
	}{
		{
			name: "target is not present in the list",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			target: 1,
			elem:   2,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
			wantErr: ErrNotFound,
		},
		{
			name: "target is the only element in the list",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			target: 1,
			elem:   2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the first element in the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n3 := newNode[int](3)
				n1.next = n3
				n3.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n3,
				}
			},
			target: 1,
			elem:   2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the last element in the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			target: 2,
			elem:   3,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is in the middle of the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n4 := newNode[int](4)
				n1.next = n2
				n2.prev = n1
				n2.next = n4
				n4.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n4,
				}
			},
			target: 2,
			elem:   3,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n4 := newNode[int](4)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				n3.next = n4
				n4.prev = n3
				return &List[int]{
					len:   4,
					first: n1,
					last:  n4,
				}
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			if gotErr := initialList.InsertAfter(tc.target, tc.elem); !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("want ErrNotFound, got %v", gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("\nwant list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func Test_List_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *List[int]
		target         int
		newWantList    func() *List[int]
		wantErr        error
	}{
		{
			name: "target is not present in the list",
			newInitialList: func() *List[int] {
				return NewList[int]()
			},
			target: 1,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
			wantErr: ErrNotFound,
		},
		{
			name: "target is the only element in the list",
			newInitialList: func() *List[int] {
				n := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			target: 1,
			newWantList: func() *List[int] {
				return NewList[int]()
			},
			wantErr: nil,
		},
		{
			name: "target is the first element of a two-node list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			target: 1,
			newWantList: func() *List[int] {
				n2 := newNode[int](2)
				return &List[int]{
					len:   1,
					first: n2,
					last:  n2,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the first element list with more than two nodes",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			target: 1,
			newWantList: func() *List[int] {
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   2,
					first: n2,
					last:  n3,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the last element of a two-node list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			target: 2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				return &List[int]{
					len:   1,
					first: n1,
					last:  n1,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is the last element of a list with more than two nodes",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			target: 3,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n1.next = n2
				n2.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			wantErr: nil,
		},
		{
			name: "target is in the middle of the list",
			newInitialList: func() *List[int] {
				n1 := newNode[int](1)
				n2 := newNode[int](2)
				n3 := newNode[int](3)
				n1.next = n2
				n2.prev = n1
				n2.next = n3
				n3.prev = n2
				return &List[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			target: 2,
			newWantList: func() *List[int] {
				n1 := newNode[int](1)
				n3 := newNode[int](3)
				n1.next = n3
				n3.prev = n1
				return &List[int]{
					len:   2,
					first: n1,
					last:  n3,
				}
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			initialList := tc.newInitialList()
			wantList := tc.newWantList()

			if gotErr := initialList.Remove(tc.target); !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("want ErrNotFound, got %v", gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("\nwant list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}

func newNode[E comparable](elem E) *node[E] {
	return &node[E]{data: elem}
}
