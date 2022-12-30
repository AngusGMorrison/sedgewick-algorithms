package q4_dedup

import (
	"reflect"
	"testing"
)

func Test_Dedup(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "input contains duplicates",
			in:   []string{"non", "hi", "seven", "hi", "seven", "seven"},
			want: []string{"hi", "non", "seven"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Dedup(tc.in)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.want, got)
			}

		})
	}
}
