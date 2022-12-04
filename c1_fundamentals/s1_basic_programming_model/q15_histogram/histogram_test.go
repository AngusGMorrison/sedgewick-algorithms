package q15_histogram

import (
	"reflect"
	"testing"
)

func Test_histogram(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		a    []int
		m    int
		want []int
	}{
		{
			name: "a is nil",
			a:    nil,
			m:    1,
			want: nil,
		},
		{
			name: "a is empty",
			a:    []int{},
			m:    1,
			want: nil,
		},
		{
			name: "m is 0",
			a:    []int{1},
			m:    0,
			want: nil,
		},
		{
			name: "a contains a number greater than m",
			a:    []int{4},
			m:    3,
			want: nil,
		},
		{
			name: "a contains a number equal to m",
			a:    []int{3},
			m:    3,
			want: nil,
		},
		{
			name: "a contains only one number",
			a:    []int{0, 0, 0},
			m:    1,
			want: []int{3},
		},
		{
			name: "a several numbers",
			a:    []int{4, 4, 3, 2, 0, 4, 2, 5},
			m:    6,
			want: []int{1, 0, 2, 1, 3, 1},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := histogram(tc.a, tc.m); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("histogram(%+v, %d): want %+v, got %+v", tc.a, tc.m, tc.want, got)
			}
		})
	}
}
