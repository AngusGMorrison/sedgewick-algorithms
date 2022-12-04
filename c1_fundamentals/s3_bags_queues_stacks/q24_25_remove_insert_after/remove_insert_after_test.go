package q24_25_remove_insert_after

import (
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func Test_removeAfter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input *list.Node[int]
		want  *list.Node[int]
	}{
		{
			name:  "input is nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "input is the last node in the list",
			input: &list.Node[int]{Data: 1},
			want:  &list.Node[int]{Data: 1},
		},
		{
			name: "input.next is the last node in the list",
			input: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 2,
				},
			},
			want: &list.Node[int]{Data: 1},
		},
		{
			name: "input.next is a middle node",
			input: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 2,
					Next: &list.Node[int]{
						Data: 3,
					},
				},
			},
			want: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 3,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			removeAfter[int](tc.input)
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("want\n\t%s,\ngot\n\t%s", tc.want, tc.input)
			}
		})
	}
}

func Test_insertAfter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		list *list.Node[int]
		node *list.Node[int]
		want *list.Node[int]
	}{
		{
			name: "list is nil",
			list: nil,
			node: &list.Node[int]{Data: 1},
			want: nil,
		},
		{
			name: "node to insert is nil",
			list: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 2,
				},
			},
			node: nil,
			want: &list.Node[int]{Data: 1},
		},
		{
			name: "next node is nil",
			list: &list.Node[int]{Data: 1},
			node: &list.Node[int]{Data: 2},
			want: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 2,
				},
			},
		},
		{
			name: "next node is not nil",
			list: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 2,
				},
			},
			node: &list.Node[int]{Data: 3},
			want: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 3,
					Next: &list.Node[int]{
						Data: 2,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			insertAfter[int](tc.list, tc.node)
			if !reflect.DeepEqual(tc.list, tc.want) {
				t.Errorf("want list\n\t%s,\ngot\n\t%s", tc.want, tc.list)
			}
		})
	}
}
