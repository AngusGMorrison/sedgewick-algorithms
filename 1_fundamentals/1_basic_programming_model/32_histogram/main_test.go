package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"testing"
)

func Test_newHistogram(t *testing.T) {
	t.Parallel()

	t.Run("instantiation succeeds", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			n            uint
			lower, upper float64
			want         histogram
		}{
			{
				n:     1,
				lower: 1,
				upper: 2,
				want: histogram{
					nBuckets: 1,
					lower:    1,
					upper:    2,
					buckets:  make([]int, 1),
				},
			},
			{
				n:     3,
				lower: 1,
				upper: 2,
				want: histogram{
					nBuckets: 3,
					lower:    1,
					upper:    2,
					buckets:  make([]int, 3),
				},
			},
		}

		for _, tc := range testCases {
			tc := tc
			name := fmt.Sprintf("newHistogram(%d, %.3f, %.3f)", tc.n, tc.lower, tc.upper)

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, err := newHistogram(tc.n, tc.lower, tc.upper)
				if err != nil {
					t.Fatalf("%s returned unexpected error: %v", name, err)
				}

				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("%s:\nwant:\n\t%+v,\ngot:\n\t%+v", name, tc.want, got)
				}
			})
		}
	})

	t.Run("instantiation fails", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name         string
			n            uint
			lower, upper float64
			want         error
		}{
			{
				name:  "n == 0",
				n:     0,
				lower: 1,
				upper: 2,
				want:  errNoBuckets,
			},
			{
				name:  "lower > upper",
				n:     1,
				lower: 2,
				upper: 1,
				want:  errInvalidRange,
			},
			{
				name:  "lower == upper",
				n:     1,
				lower: 1,
				upper: 1,
				want:  errInvalidRange,
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotH, gotErr := newHistogram(tc.n, tc.lower, tc.upper)
				if !errors.Is(gotErr, tc.want) {
					t.Errorf("%s: want error %q, got %q", tc.name, tc.want, gotErr)
				}

				if !reflect.DeepEqual(gotH, histogram{}) {
					t.Errorf("%s: want empty histogram, got %+v", tc.name, gotH)
				}
			})
		}
	})
}

func Test_histogram_bucketSize(t *testing.T) {
	t.Parallel()

	lower := 11.9
	upper := 62.3
	nBuckets := uint(5)
	wantBucketSize := 10.08
	h, err := newHistogram(nBuckets, lower, upper)
	if err != nil {
		t.Fatalf("newHistogram returned unexpected error: %v", err)
	}

	if gotBucketSize := h.bucketSize(); !approxEqual(gotBucketSize, wantBucketSize) {
		t.Errorf("want bucket size %.6f, got %.6f", wantBucketSize, gotBucketSize)
	}
	if !approxEqual(h.bSize, wantBucketSize) {
		t.Errorf("bucket size wasn't cached on the histogram: want %.6f, got %.6f", wantBucketSize, h.bSize)
	}
}

func Test_histogram_bucketFor(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		lower, upper float64
		nBuckets     uint
		v            float64
		wantBucket   int
	}{
		{
			name:       "input < lower bound",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          -2,
			wantBucket: -1,
		},
		{
			name:       "input == lower bound",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          -1,
			wantBucket: -1,
		},
		{
			name:       "input > upper bound",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          2,
			wantBucket: -1,
		},
		{
			name:       "input == upper bound",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          1,
			wantBucket: -1,
		},
		{
			name:       "single bucket, positive input",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          0.5,
			wantBucket: 0,
		},
		{
			name:       "single bucket, negative input",
			lower:      -1,
			upper:      1,
			nBuckets:   1,
			v:          -0.5,
			wantBucket: 0,
		},
		{
			name:       "three buckets, input in first",
			lower:      -1,
			upper:      1,
			nBuckets:   3,
			v:          -(2.0 / 3.0),
			wantBucket: 0,
		},
		{
			name:       "three buckets, input on boundary of first and second",
			lower:      -1,
			upper:      1,
			nBuckets:   3,
			v:          -(1.0 / 3.0),
			wantBucket: 1,
		},
		{
			name:       "three buckets, input in second",
			lower:      -1,
			upper:      1,
			nBuckets:   3,
			v:          0,
			wantBucket: 1,
		},
		{
			name:       "three buckets, input on boundary of second and third",
			lower:      -1,
			upper:      1,
			nBuckets:   3,
			v:          1.0 / 3.0,
			wantBucket: 2,
		},
		{
			name:       "three buckets, input in third",
			lower:      -1,
			upper:      1,
			nBuckets:   3,
			v:          2.0 / 3.0,
			wantBucket: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h, err := newHistogram(tc.nBuckets, tc.lower, tc.upper)
			if err != nil {
				t.Fatalf("newHistogram returned unexpected error: %v", err)
			}

			gotBucket := h.bucketFor(tc.v)
			if gotBucket != tc.wantBucket {
				t.Errorf("h.bucketFor(%.3f): want %d, got %d", tc.v, tc.wantBucket, gotBucket)
			}
		})
	}
}

func Test_histogram_add(t *testing.T) {
	t.Parallel()

	nBuckets := uint(10)
	lower := -5.0
	upper := 5.0
	data := []float64{
		-6.0, -5.0, // out of bounds
		-4.9, -4.6, -4.2, // bucket 0
		-4.0, -3.1, // 1
		-2.8,                         // 2
		-2.0, -1.9, -1.7, -1.6, -1.2, // 3
		-0.4, -0.3, // 4
		0.0, 0.5, 0.6, // 5
		1.1, 1.4, // 6
		2.3,           // 7
		3.0, 3.6, 3.8, // 8
		4.7,      // 9
		5.0, 6.0, // out of bounds
	}
	wantBuckets := []int{3, 2, 1, 5, 2, 3, 2, 1, 3, 1}
	h, err := newHistogram(nBuckets, lower, upper)
	if err != nil {
		t.Fatalf("newHistogram returned unexpected error: %v", err)
	}

	h.add(data...)
	if !reflect.DeepEqual(h.buckets, wantBuckets) {
		t.Errorf("\nwant:\n\t%v,\ngot:\n\t%v", wantBuckets, h.buckets)
	}
}

const float64EqualityThreshold = 1e-7

func approxEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) <= float64EqualityThreshold
}
