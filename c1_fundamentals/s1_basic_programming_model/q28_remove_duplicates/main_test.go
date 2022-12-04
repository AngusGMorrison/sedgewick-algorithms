package main

import (
	"reflect"
	"testing"
)

func Test_dedup(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		in, want []int
	}{
		{
			name: "slice is nil",
			in:   nil,
			want: nil,
		},
		{
			name: "slice is empty",
			in:   []int{},
			want: []int{},
		},
		{
			name: "slice has one element",
			in:   []int{1},
			want: []int{1},
		},
		{
			name: "slice has no duplicates",
			in:   []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "slice has duplicates",
			in:   []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 4, 5, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := dedup(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want:\n\t%v\ngot\n\t%v", tc.want, got)
			}
		})
	}
}
