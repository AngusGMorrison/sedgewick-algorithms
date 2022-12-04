package q39_ring_buffer

import (
	"errors"
	"reflect"
	"testing"
)

func Test_Ring_Publish(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		ring     *Ring[int]
		data     int
		wantRing *Ring[int]
		wantErr  error
	}{
		{
			name: "buffer is empty and writer does not wrap",
			ring: &Ring[int]{
				len:    0,
				cap:    3,
				buffer: make([]int, 3),
				reader: 0,
				writer: 0,
			},
			data: 1,
			wantRing: &Ring[int]{
				len:    1,
				cap:    3,
				buffer: []int{1, 0, 0},
				reader: 0,
				writer: 1,
			},
			wantErr: nil,
		},
		{
			name: "buffer is not empty and writer does not wrap",
			ring: &Ring[int]{
				len:    1,
				cap:    3,
				buffer: []int{1, 0, 0},
				reader: 0,
				writer: 1,
			},
			data: 2,
			wantRing: &Ring[int]{
				len:    2,
				cap:    3,
				buffer: []int{1, 2, 0},
				reader: 0,
				writer: 2,
			},
			wantErr: nil,
		},
		{
			name: "buffer is not empty and writer wraps",
			ring: &Ring[int]{
				len:    2,
				cap:    3,
				buffer: []int{1, 2, 0},
				reader: 0,
				writer: 2,
			},
			data: 3,
			wantRing: &Ring[int]{
				len:    3,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			wantErr: nil,
		},
		{
			name: "buffer is full",
			ring: &Ring[int]{
				len:    3,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			data: 4,
			wantRing: &Ring[int]{
				len:    3,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			wantErr: ErrFull,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotErr := tc.ring.Publish(tc.data); !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("\nwant error\n\t%v\ngot\n\t%v", tc.wantErr, gotErr)
			}
			if !reflect.DeepEqual(tc.ring, tc.wantRing) {
				t.Errorf("\nwant ring\n\t%+v\ngot\n\t%+v", tc.wantRing, tc.ring)
			}
		})
	}
}

func Test_Ring_Consume(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		ring     *Ring[int]
		wantData int
		wantErr  error
		wantRing *Ring[int]
	}{
		{
			name: "buffer is full and reader does not wrap",
			ring: &Ring[int]{
				len:    3,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			wantData: 1,
			wantErr:  nil,
			wantRing: &Ring[int]{
				len:    2,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 1,
				writer: 0,
			},
		},
		{
			name: "buffer is full and reader wraps",
			ring: &Ring[int]{
				len:    3,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 2,
				writer: 2,
			},
			wantData: 3,
			wantErr:  nil,
			wantRing: &Ring[int]{
				len:    2,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 2,
			},
		},
		{
			name: "buffer is part-full and reader does not wrap",
			ring: &Ring[int]{
				len:    1,
				cap:    3,
				buffer: []int{1, 0, 0},
				reader: 0,
				writer: 1,
			},
			wantData: 1,
			wantErr:  nil,
			wantRing: &Ring[int]{
				len:    0,
				cap:    3,
				buffer: []int{1, 0, 0},
				reader: 1,
				writer: 1,
			},
		},
		{
			name: "buffer is part-full and reader wraps",
			ring: &Ring[int]{
				len:    1,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 2,
				writer: 0,
			},
			wantData: 3,
			wantRing: &Ring[int]{
				len:    0,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			wantErr: nil,
		},
		{
			name: "buffer is empty",
			ring: &Ring[int]{
				len:    0,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
			wantData: 0,
			wantErr:  ErrEmpty,
			wantRing: &Ring[int]{
				len:    0,
				cap:    3,
				buffer: []int{1, 2, 3},
				reader: 0,
				writer: 0,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotData, gotErr := tc.ring.Consume()
			if gotData != tc.wantData {
				t.Errorf("\nwant data\n\t%v\ngot\n\t%v", tc.wantData, gotData)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("\nwant error\n\t%v\ngot\n\t%v", tc.wantErr, gotErr)
			}
			if !reflect.DeepEqual(tc.ring, tc.wantRing) {
				t.Errorf("\nwant ring\n\t%+v\ngot\n\t%+v", tc.wantRing, tc.ring)
			}
		})
	}
}
