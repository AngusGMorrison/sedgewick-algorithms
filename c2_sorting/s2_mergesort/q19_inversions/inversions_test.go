package q19_inversions

import "testing"

func Test_Inversions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "input is empty",
			input: nil,
			want:  0,
		},
		{
			name:  "len(input) == 1",
			input: []int{1},
			want:  0,
		},
		{
			name:  "len(input) == 2, no inversion",
			input: []int{1, 2},
			want:  0,
		},
		{
			name:  "len(input) == 2, has inversion",
			input: []int{2, 1},
			want:  1,
		},
		{
			name:  "len(input) == 3, no inversions",
			input: []int{1, 2, 3},
			want:  0,
		},
		{
			name:  "len(input) == 3, 1 inversion",
			input: []int{2, 1, 3},
			want:  1,
		},
		{
			name:  "len(input) == 3, 2 inversions",
			input: []int{3, 1, 2},
			want:  2,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Inversions(tc.input)
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
