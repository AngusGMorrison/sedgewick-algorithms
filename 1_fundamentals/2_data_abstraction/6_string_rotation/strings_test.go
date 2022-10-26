package strings

import (
	"fmt"
	"testing"
)

func Test_isRotation(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		s1, s2 string
		want   bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"a", "aa", false},
		{"TGACGAC", "ACTGACG", true},
		{"TGACGAC", "ACTGACA", false},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("isRotation(%s, %s)", tc.s1, tc.s2)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := isRotation(tc.s1, tc.s2); got != tc.want {
				t.Errorf("%s: want %t, got %t", name, tc.want, got)
			}
		})
	}
}
