package ln

import (
	"fmt"
	"math"
	"testing"
)

func Test_lnFactorial(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n    int
		want float64
	}{
		{0, 0},
		{1, 0},
		{2, 0.69314},
		{3, 1.79175},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("lnFactorial(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := lnFactorial(tc.n); !approxEqual(t, got, tc.want) {
				t.Errorf("%s: want %.5f, got %.5f", name, tc.want, got)
			}
		})

	}
}

// Rounding error prevents us from comparing floats for equality directly. Instead, we set some
// threshold and treat floats that differ by less than this threshold as equal.
const float64EqualityThreshhold = 1e-5

func approxEqual(t *testing.T, f1, f2 float64) bool {
	t.Helper()

	return math.Abs(f1-f2) <= float64EqualityThreshhold
}
