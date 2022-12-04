package collinearity

import "testing"

func Test_CountCollinearN2(t *testing.T) {
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

			if got := CountCollinearN2(tc.input); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_binomialCoefficientGenerator_nCr(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		n, r int
		want int
	}{
		{
			name: "n = 0",
			n:    0,
			r:    0,
			want: 0,
		},
		{
			name: "n < 0",
			n:    -1,
			r:    -1,
			want: 0,
		},
		{
			name: "r < 0",
			n:    1,
			r:    -1,
			want: 0,
		},
		{
			name: "n < r",
			n:    1,
			r:    2,
			want: 0,
		},
		{
			name: "r = 0",
			n:    1,
			r:    0,
			want: 1,
		},
		{
			name: "n = 1, r = 1",
			n:    1,
			r:    1,
			want: 1,
		},
		{
			name: "n = 2, r = 1",
			n:    2,
			r:    1,
			want: 2,
		},
		{
			name: "n = 2, r = 2",
			n:    2,
			r:    2,
			want: 1,
		},
		{
			name: "n = 3, r = 1",
			n:    3,
			r:    1,
			want: 3,
		},
		{
			name: "n = 3, r = 2",
			n:    3,
			r:    2,
			want: 3,
		},
		{
			name: "n = 3, r = 3",
			n:    3,
			r:    3,
			want: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bcg := newBinomialCoefficientGenerator()
			if got := bcg.nCr(tc.n, tc.r); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
