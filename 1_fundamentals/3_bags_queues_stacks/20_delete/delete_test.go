package list

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func (n *node[E]) String() string {
	if n == nil {
		return "nil"
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%v", n.data))
	for cur := n.next; cur != nil; cur = cur.next {
		_, _ = fmt.Fprintf(&builder, " -> %v", cur.data)
	}

	return builder.String()
}

func Test_delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input *node[int]
		k     uint
		want  *node[int]
	}{
		{
			name:  "list is nil",
			input: nil,
			k:     1,
			want:  nil,
		},
		{
			name: "k == 0",
			input: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
			k: 0,
			want: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
		},
		{
			name: "k > len(list)",
			input: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
			k: 3,
			want: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
		},
		{
			name:  "list has one element, k == 1",
			input: &node[int]{data: 1},
			k:     1,
			want:  nil,
		},
		{
			name: "list has two elements, k == 1",
			input: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
			k: 1,
			want: &node[int]{
				data: 2,
			},
		},
		{
			name: "list has two elements, k == 2",
			input: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
				},
			},
			k: 2,
			want: &node[int]{
				data: 1,
			},
		},
		{
			name: "deletion of a middle element",
			input: &node[int]{
				data: 1,
				next: &node[int]{
					data: 2,
					next: &node[int]{
						data: 3,
					},
				},
			},
			k: 2,
			want: &node[int]{
				data: 1,
				next: &node[int]{
					data: 3,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := delete[int](tc.input, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want list\n\t%s,\ngot\n\t%s", tc.want, got)
			}
		})
	}
}
