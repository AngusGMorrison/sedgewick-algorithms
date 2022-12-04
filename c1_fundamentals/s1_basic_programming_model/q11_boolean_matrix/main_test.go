package main

import (
	"bytes"
	"testing"
)

func Test_printMatrix(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		m    [][]bool
		want string
	}{
		{
			name: "nil matrix",
			m:    nil,
			want: "",
		},
		{
			name: "empty matrix",
			m:    [][]bool{},
			want: "",
		},
		{
			name: "empty first row",
			m: [][]bool{
				{},
			},
			want: "",
		},
		{
			name: "empty first row",
			m: [][]bool{
				{true},
			},
			want: "\t1\n1\t*\n\n",
		},
		{
			name: "3x3",
			m: [][]bool{
				{true, true, true},
				{false, true, false},
				{true, true, false},
			},
			want: "\t1\t2\t3\n1\t*\t*\t*\n2\t \t*\t \n3\t*\t*\t \n\n",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}
			printMatrix(buf, tc.m)

			if got := buf.String(); got != tc.want {
				t.Errorf("%s: want\n%s\ngot\n%s\n", tc.name, tc.want, got)
			}
		})
	}
}
