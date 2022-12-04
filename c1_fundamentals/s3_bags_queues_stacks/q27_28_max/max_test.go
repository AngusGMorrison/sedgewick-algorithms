package q27_28_max

import (
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func Test_max(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		list *list.Node[int]
		want int
	}{
		{
			name: "list is nil",
			list: nil,
			want: 0,
		},
		{
			name: "list is non-nil",
			list: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 3,
					Next: &list.Node[int]{
						Data: 1,
					},
				},
			},
			want: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := max(tc.list); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_maxRecursive(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		list *list.Node[int]
		want int
	}{
		{
			name: "list is nil",
			list: nil,
			want: 0,
		},
		{
			name: "list is non-nil",
			list: &list.Node[int]{
				Data: 2,
				Next: &list.Node[int]{
					Data: 3,
					Next: &list.Node[int]{
						Data: 1,
					},
				},
			},
			want: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := maxRecursive(tc.list); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
