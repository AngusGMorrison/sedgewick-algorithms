package q9_binary_string

import (
	"fmt"
	"testing"
)

func Test_toLittleEndianString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "01"},
		{3, "11"},
		{4, "001"},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("toLittleEndianString(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := toLittleEndianString(tc.n); got != tc.want {
				t.Errorf("%s: want %s, got %s", name, tc.want, got)
			}
		})
	}
}

func Test_toBigEndianString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("toBigEndianString(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := toBigEndianString(tc.n); got != tc.want {
				t.Errorf("%s: want %s, got %s", name, tc.want, got)
			}
		})
	}
}
