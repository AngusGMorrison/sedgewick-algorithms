package main

import (
	"bytes"
	"errors"
	"os"
	"reflect"
	"testing"
)

func Test_readTable(t *testing.T) {
	t.Parallel()

	t.Run("read successful", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name    string
			fixture string
			want    table
		}{
			{
				name:    "columns are separated by single spaces",
				fixture: "testdata/ok.txt",
				want: table{
					{name: "Angus", score1: 34, score2: 192},
					{name: "Bianca", score1: 54, score2: 218},
					{name: "Steve", score1: 11, score2: 188},
				},
			},
			{
				name:    "columns are separated by variable amounts of whitespace",
				fixture: "testdata/extra_whitespace.txt",
				want: table{
					{name: "Angus", score1: 34, score2: 192},
					{name: "Bianca", score1: 54, score2: 218},
					{name: "Steve", score1: 11, score2: 188},
				},
			},
			{
				name:    "table is empty",
				fixture: "testdata/empty.txt",
				want:    nil,
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				file, err := os.Open(tc.fixture)
				if err != nil {
					t.Fatal(err)
				}
				defer file.Close()

				got, err := readTable(file)
				if err != nil {
					t.Fatal(err)
				}

				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("readInTable(%s): want %+v, got %+v", tc.fixture, tc.want, got)
				}
			})
		}
	})

	t.Run("read failure", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name    string
			fixture string
			wantErr error
		}{
			{
				name:    "row contains too few fields",
				fixture: "testdata/too_few_fields.txt",
				wantErr: errInsufficientFields,
			},
			{
				name:    "score 1 is not an integer",
				fixture: "testdata/non_int_score_1.txt",
				wantErr: atoiError{score: "Colin"},
			},
			{
				name:    "score 2 is not an integer",
				fixture: "testdata/non_int_score_2.txt",
				wantErr: atoiError{score: "Colin"},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				file, err := os.Open(tc.fixture)
				if err != nil {
					t.Fatal(err)
				}
				defer file.Close()

				got, err := readTable(file)
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("readInTable(%s): want error %+v, got %+v", tc.fixture, tc.wantErr, err)
				}

				if got != nil {
					t.Errorf("readInTable(%s): want table to be nil, got %+v", tc.fixture, got)
				}
			})
		}
	})
}

func Test_table_WriteTo(t *testing.T) {
	t.Parallel()

	t.Run("write successful", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name  string
			table table
			want  string
		}{
			{
				name:  "table is nil",
				table: nil,
				want:  "Name\tScore 1\tScore 2\tRatio\n",
			},
			{
				name:  "table is empty",
				table: table{},
				want:  "Name\tScore 1\tScore 2\tRatio\n",
			},
			{
				name: "table is populated",
				table: table{
					{name: "Angus", score1: 34, score2: 192},
					{name: "Bianca", score1: 54, score2: 218},
					{name: "Steve", score1: 11, score2: 188},
				},
				want: "Name\tScore 1\tScore 2\tRatio\nAngus\t34\t192\t0.177\nBianca\t54\t218\t0.248\nSteve\t11\t188\t0.059\n",
			},
			{
				name: "table contains names of great length",
				table: table{
					{name: "AngusHoratioConningsby", score1: 34, score2: 192},
					{name: "Bianca", score1: 54, score2: 218},
					{name: "Steve", score1: 11, score2: 188},
				},
				want: "Name\t\t\t\tScore 1\tScore 2\tRatio\nAngusHoratioConningsby\t34\t192\t0.177\nBianca\t\t\t\t54\t218\t0.248\nSteve\t\t\t\t11\t188\t0.059\n",
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				buf := &bytes.Buffer{}
				if _, err := tc.table.WriteTo(buf); err != nil {
					t.Fatal(err)
				}

				if got := buf.String(); got != tc.want {
					t.Errorf("table.WriteTo(&bytes.Buffer{}):\nwant\n\t%q,\ngot\n\t%q", tc.want, got)
				}
			})
		}
	})

	t.Run("write fails", func(t *testing.T) {
		t.Parallel()

		wantErr := errors.New("some error")
		writer := &mockWriter{wantErr: wantErr}
		tab := table{}

		if _, err := tab.WriteTo(writer); !errors.Is(err, wantErr) {
			t.Fatalf("table.WriteTo(&mockWriter{}): want error %q, got %q", wantErr, err)
		}
	})
}

type mockWriter struct {
	wantErr error
}

func (m *mockWriter) Write(b []byte) (int, error) {
	return 0, m.wantErr
}
