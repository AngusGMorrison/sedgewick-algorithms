package q29_equal_keys

import "testing"

func Test_countSmallerThan(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		key  int
		a    []int
		want int
	}{
		{
			name: "slice is nil",
			key:  5,
			a:    nil,
			want: -1,
		},
		{
			name: "slice is empty",
			key:  5,
			a:    []int{},
			want: -1,
		},
		{
			name: "key is not duplicated",
			key:  3,
			a:    []int{1, 2, 3, 4, 5},
			want: 2,
		},
		{
			name: "key is duplicated",
			key:  3,
			a:    []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 5},
			want: 2,
		},
		{
			name: "key is the first element",
			key:  1,
			a:    []int{1, 1, 1, 1, 1, 2, 3, 4, 5},
			want: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := countSmallerThan(tc.key, tc.a); got != tc.want {
				t.Errorf("want: %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_countGreaterThan(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		key  int
		a    []int
		want int
	}{
		{
			name: "slice is nil",
			key:  5,
			a:    nil,
			want: -1,
		},
		{
			name: "slice is empty",
			key:  5,
			a:    []int{},
			want: -1,
		},
		{
			name: "key is not duplicated",
			key:  3,
			a:    []int{1, 2, 3, 4, 5},
			want: 2,
		},
		{
			name: "key is duplicated",
			key:  3,
			a:    []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 5},
			want: 2,
		},
		{
			name: "key is the last element",
			key:  5,
			a:    []int{1, 2, 3, 4, 5, 5, 5, 5, 5},
			want: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := countGreaterThan(tc.key, tc.a); got != tc.want {
				t.Errorf("want: %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_count(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		key  int
		a    []int
		want int
	}{
		{
			name: "slice is nil",
			key:  5,
			a:    nil,
			want: 0,
		},
		{
			name: "slice is empty",
			key:  5,
			a:    []int{},
			want: 0,
		},
		{
			name: "single instance of key",
			key:  3,
			a:    []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "multiple instances of key",
			key:  3,
			a:    []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 5},
			want: 6,
		},
		{
			name: "key is the last element",
			key:  5,
			a:    []int{1, 2, 3, 4, 5, 5},
			want: 2,
		},
		{
			name: "key is the first element",
			key:  1,
			a:    []int{1, 1, 2, 3, 4, 5},
			want: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := count(tc.key, tc.a); got != tc.want {
				t.Errorf("want: %d, got %d", tc.want, got)
			}
		})
	}
}
