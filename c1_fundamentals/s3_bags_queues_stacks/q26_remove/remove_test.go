package q26_remove

import (
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func Test_remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		list *list.Node[int]
		key  int
		want *list.Node[int]
	}{
		{
			name: "list is nil",
			list: nil,
			key:  1,
			want: nil,
		},
		{
			name: "list is comprised of only matching keys",
			list: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 1,
					Next: &list.Node[int]{
						Data: 1,
					},
				},
			},
			key:  1,
			want: nil,
		},
		{
			name: "list is prefixed by matching keys",
			list: &list.Node[int]{
				Data: 1,
				Next: &list.Node[int]{
					Data: 1,
					Next: &list.Node[int]{
						Data: 2,
					},
				},
			},
			key: 1,
			want: &list.Node[int]{
				Data: 2,
			},
		},
		{
			name: "list ends with matching keys",
			list: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 1,
					Next: &list.Node[int]{
						Data: 1,
					},
				},
			},
			key: 1,
			want: &list.Node[int]{
				Data: 2,
			},
		},
		{
			name: "list contains matching keys",
			list: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 1,
					Next: &list.Node[int]{
						Data: 2,
					},
				},
			},
			key: 1,
			want: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 2,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := remove[int](tc.list, tc.key); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nwant\n\t%s,\ngot\n\t%s", tc.want, got)
			}
		})
	}
}
