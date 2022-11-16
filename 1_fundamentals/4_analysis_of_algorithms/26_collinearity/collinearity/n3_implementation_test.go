package collinearity

import "testing"

func Test_countCollinearN3(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []Point
		want  int
	}{
		{
			name: "input contains one collinear triple",
			input: []Point{
				{0, 0},
				{1, 1},
				{2, 2},
			},
			want: 1,
		},
		{
			name: "input contains four collinear triples along one line",
			input: []Point{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: 4,
		},
		{
			name: "input contains five collinear triples along two lines",
			input: []Point{
				{0, -2},
				{0, -1},
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: 5,
		},
		{
			name: "input contains six collinear triples along three lines",
			input: []Point{
				{0, -2},
				{0, -1},
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
				{2, 0},
				{3, 0},
			},
			want: 6,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := countCollinearN3(tc.input); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
