package q38_image_processing

import (
	"fmt"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

// Color represents an RGBA color.
type Color uint32

type colorBuffer []Color

func (cb colorBuffer) String() string {
	var builder strings.Builder
	builder.WriteString("[ ")
	for _, color := range cb {
		fmt.Fprintf(&builder, "%X ", color)
	}
	builder.WriteRune(']')
	return builder.String()
}

type Canvas struct {
	buf    colorBuffer
	width  uint
	height uint
}

func NewCanvas(w, h uint) *Canvas {
	return &Canvas{
		buf:    make([]Color, w*h),
		width:  w,
		height: h,
	}
}

// Fill fills a contiguous region having the color of pixel (x, y) with the given color.
func (c *Canvas) Fill(x, y uint, color Color) {
	// Perform a breadth-first search from the source to find all contiguous pixels of the same
	// color.
	source := c.pixelIndex(x, y)
	originalColor := c.buf[source]
	if color == originalColor {
		return
	}

	worklist := queue.NewSliceQueue[uint]()
	worklist.Enqueue(source)

	for pixelIdx, ok := worklist.Dequeue(); ok; pixelIdx, ok = worklist.Dequeue() {
		c.buf[pixelIdx] = color

		for _, neighbor := range c.neighboringIndices(pixelIdx) {
			if c.buf[neighbor] == originalColor {
				worklist.Enqueue(neighbor)
			}
		}
	}
}

func (c *Canvas) pixelIndex(x, y uint) uint {
	return c.width*y + x
}

func (c *Canvas) neighboringIndices(i uint) []uint {
	var neighbors []uint
	if !c.pixelIsInTopRow(i) {
		neighbors = append(neighbors, i-c.width)
	}
	if !c.pixelIsInBottomRow(i) {
		neighbors = append(neighbors, i+c.width)
	}
	if !c.pixelIsInLeftCol(i) {
		neighbors = append(neighbors, i-1)
	}
	if !c.pixelIsInRightCol(i) {
		neighbors = append(neighbors, i+1)
	}
	return neighbors
}

func (c *Canvas) pixelIsInTopRow(i uint) bool {
	return i < c.width
}

func (c *Canvas) pixelIsInBottomRow(i uint) bool {
	return i >= (c.height-1)*c.width && i < c.height*c.width
}

func (c *Canvas) pixelIsInLeftCol(i uint) bool {
	return i%c.width == 0
}

func (c *Canvas) pixelIsInRightCol(i uint) bool {
	return (i % c.width) == c.width-1
}
