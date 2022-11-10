package closestpair

import (
	"reflect"
	"testing"
)

func Test_ClosestPair(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []float64
		wantPair []float64
	}{
		{
			name:     "slice is empty",
			input:    nil,
			wantPair: nil,
		},
		{
			name:     "len(a) == 1",
			input:    []float64{1},
			wantPair: nil,
		},
		{
			name:     "len(a) == 2",
			input:    []float64{1, 2},
			wantPair: []float64{1, 2},
		},
		{
			name:     "len(a) > 2, ∃ distinct closest pair",
			input:    []float64{1, 3, 4, 6},
			wantPair: []float64{3, 4},
		},
		{
			name:     "len(a) > 2, ∄ distinct closest pair",
			input:    []float64{1, 4, 5, 6, 7},
			wantPair: []float64{4, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotPair := ClosestPair(tc.input); !reflect.DeepEqual(gotPair, tc.wantPair) {
				t.Errorf("\nwant\n\t%v\ngot\n\t%v", tc.wantPair, gotPair)
			}
		})
	}
}
