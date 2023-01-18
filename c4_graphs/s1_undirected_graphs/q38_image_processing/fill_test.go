package q38_image_processing

import (
	"reflect"
	"testing"
)

const (
	white Color = 0xFFFFFFFF
	black Color = 0x00000000
	red   Color = 0xFF0000FF
)

func Test_Canvas_Fill(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		width           uint
		height          uint
		initialColorBuf colorBuffer
		x, y            uint
		color           Color
		wantColorBuf    colorBuffer
	}{
		{
			name:            "1x1",
			width:           1,
			height:          1,
			initialColorBuf: colorBuffer{black},
			x:               0,
			y:               0,
			color:           white,
			wantColorBuf:    colorBuffer{white},
		},
		{
			name:            "2x1, origin left",
			width:           2,
			height:          1,
			initialColorBuf: colorBuffer{black, black},
			x:               0,
			y:               0,
			color:           white,
			wantColorBuf:    colorBuffer{white, white},
		},
		{
			name:            "2x1, origin right",
			width:           2,
			height:          1,
			initialColorBuf: colorBuffer{black, black},
			x:               1,
			y:               0,
			color:           white,
			wantColorBuf:    colorBuffer{white, white},
		},
		{
			name:            "1x2, origin top",
			width:           1,
			height:          2,
			initialColorBuf: colorBuffer{black, black},
			x:               0,
			y:               0,
			color:           white,
			wantColorBuf:    colorBuffer{white, white},
		},
		{
			name:            "1x2, origin bottom",
			width:           1,
			height:          2,
			initialColorBuf: colorBuffer{black, black},
			x:               0,
			y:               1,
			color:           white,
			wantColorBuf:    colorBuffer{white, white},
		},
		{
			name:   "3x3, origin top left, uniform fill",
			width:  3,
			height: 3,
			initialColorBuf: colorBuffer{
				black, black, black,
				black, black, black,
				black, black, black,
			},
			x:     0,
			y:     0,
			color: white,
			wantColorBuf: colorBuffer{
				white, white, white,
				white, white, white,
				white, white, white,
			},
		},
		{
			name:   "3x3, origin centre, uniform fill",
			width:  3,
			height: 3,
			initialColorBuf: colorBuffer{
				black, black, black,
				black, black, black,
				black, black, black,
			},
			x:     1,
			y:     1,
			color: white,
			wantColorBuf: colorBuffer{
				white, white, white,
				white, white, white,
				white, white, white,
			},
		},
		{
			name:   "3x3, origin bottom right, uniform fill",
			width:  3,
			height: 3,
			initialColorBuf: colorBuffer{
				black, black, black,
				black, black, black,
				black, black, black,
			},
			x:     2,
			y:     2,
			color: white,
			wantColorBuf: colorBuffer{
				white, white, white,
				white, white, white,
				white, white, white,
			},
		},
		{
			name:   "3x3, partial fill",
			width:  3,
			height: 3,
			initialColorBuf: colorBuffer{
				black, black, red,
				black, red, red,
				red, red, black,
			},
			x:     1,
			y:     1,
			color: white,
			wantColorBuf: colorBuffer{
				black, black, white,
				black, white, white,
				white, white, black,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			canvas := &Canvas{
				width:  tc.width,
				height: tc.height,
				buf:    tc.initialColorBuf,
			}
			canvas.Fill(tc.x, tc.y, tc.color)
			if !reflect.DeepEqual(tc.wantColorBuf, canvas.buf) {
				t.Errorf("want\n\t%s\ngot\n\t%s", tc.wantColorBuf, canvas.buf)
			}
		})
	}
}
