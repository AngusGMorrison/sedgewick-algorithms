package q38_delete_kth

import (
	"errors"
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func Test_ListQueue_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		newInitialList func() *ListQueue[int]
		k              int
		wantElem       int
		wantErr        error
		newWantList    func() *ListQueue[int]
	}{
		{
			name: "k is out of bounds",
			newInitialList: func() *ListQueue[int] {
				return &ListQueue[int]{}
			},
			k:        0,
			wantElem: 0,
			wantErr:  ErrOutOfBounds,
			newWantList: func() *ListQueue[int] {
				return &ListQueue[int]{}
			},
		},
		{
			name: "delete only element",
			newInitialList: func() *ListQueue[int] {
				n := &list.Node[int]{Data: 1}
				return &ListQueue[int]{
					len:   1,
					first: n,
					last:  n,
				}
			},
			k:        0,
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				return &ListQueue[int]{}
			},
		},
		{
			name: "delete first element of two-node list",
			newInitialList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n1.Next = n2
				return &ListQueue[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			k:        0,
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				n2 := &list.Node[int]{Data: 2}
				return &ListQueue[int]{
					len:   1,
					first: n2,
					last:  n2,
				}
			},
		},
		{
			name: "delete first element of list with more than two nodes",
			newInitialList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n3 := &list.Node[int]{Data: 3}
				n1.Next = n2
				n2.Next = n3
				return &ListQueue[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			k:        0,
			wantElem: 1,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				n2 := &list.Node[int]{Data: 2}
				n3 := &list.Node[int]{Data: 3}
				n2.Next = n3
				return &ListQueue[int]{
					len:   2,
					first: n2,
					last:  n3,
				}
			},
		},
		{
			name: "delete last element of two-node list",
			newInitialList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n1.Next = n2
				return &ListQueue[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
			k:        1,
			wantElem: 2,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				return &ListQueue[int]{
					len:   1,
					first: n1,
					last:  n1,
				}
			},
		},
		{
			name: "delete last element of list with more than two nodes",
			newInitialList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n3 := &list.Node[int]{Data: 3}
				n1.Next = n2
				n2.Next = n3
				return &ListQueue[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			k:        2,
			wantElem: 3,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n1.Next = n2
				return &ListQueue[int]{
					len:   2,
					first: n1,
					last:  n2,
				}
			},
		},
		{
			name: "delete element in middle of list",
			newInitialList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n2 := &list.Node[int]{Data: 2}
				n3 := &list.Node[int]{Data: 3}
				n1.Next = n2
				n2.Next = n3
				return &ListQueue[int]{
					len:   3,
					first: n1,
					last:  n3,
				}
			},
			k:        1,
			wantElem: 2,
			wantErr:  nil,
			newWantList: func() *ListQueue[int] {
				n1 := &list.Node[int]{Data: 1}
				n3 := &list.Node[int]{Data: 3}
				n1.Next = n3
				return &ListQueue[int]{
					len:   2,
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

			gotElem, gotErr := initialList.Delete(tc.k)
			if gotElem != tc.wantElem {
				t.Errorf("\nwant elem %d, got %d", tc.wantElem, gotElem)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("\nwant err %v, got %v", tc.wantErr, gotErr)
			}
			if !reflect.DeepEqual(initialList, wantList) {
				t.Errorf("\nwant list\n\t%s\ngot\n\t%s", wantList, initialList)
			}
		})
	}
}
