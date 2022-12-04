package q17_farthest_pair

import (
	"reflect"
	"testing"
)

func Test_FarthestPair(t *testing.T) {
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
			name:     "len(input) == 1",
			input:    []float64{1},
			wantPair: nil,
		},
		{
			name:     "len(input) == 2",
			input:    []float64{2, 1},
			wantPair: []float64{1, 2},
		},
		{
			name:     "len(input) > 2",
			input:    []float64{2, 1, 4, 9, 3},
			wantPair: []float64{1, 9},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotPair := FarthestPair(tc.input); !reflect.DeepEqual(gotPair, tc.wantPair) {
				t.Errorf("/nwant\n\t%v\ngot\n\t%v", tc.wantPair, gotPair)
			}
		})
	}
}
