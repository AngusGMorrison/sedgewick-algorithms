package prime

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_buildRelativePrimeMatrix(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		size uint
		want [][]bool
	}{
		{
			size: 0,
			want: nil,
		},
		{
			size: 1,
			want: [][]bool{
				{false},
			},
		},
		{
			size: 2,
			want: [][]bool{
				{false, true},
				{true, true},
			},
		},
		{
			size: 3,
			want: [][]bool{
				{false, true, false},
				{true, true, true},
				{false, true, false},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("buildRelativePrimeMatrix(%d)", tc.size)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := buildRelativePrimeMatrix(tc.size); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%s: want:\n\t%v;\ngot:\n\t%v\n", name, tc.want, got)
			}
		})
	}
}
