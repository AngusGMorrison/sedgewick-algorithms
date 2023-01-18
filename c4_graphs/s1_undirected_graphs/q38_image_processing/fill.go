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
	source := c.pixelIndex(x, y)
	originalColor := c.buf[source]
	if color == originalColor {
		return
	}

	worklist := queue.NewSliceQueue[uint]()
	worklist.Enqueue(source)

	for pixelIdx, ok := worklist.Dequeue(); ok; pixelIdx, ok = worklist.Dequeue() {
		c.buf[pixelIdx] = color
		neighbors := c.neighboringIndices(pixelIdx)
		fmt.Println(neighbors)
		for _, neighbor := range neighbors {
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
	inTopRow := c.pixelIsInTopRow(i)
	inBottomRow := c.pixelIsInBottomRow(i)
	inLeftCol := c.pixelIsInLeftCol(i)
	inRightCol := c.pixelIsInRightCol(i)

	if inTopRow && inBottomRow && inLeftCol && inRightCol {
		return nil
	}
	if inTopRow && inBottomRow && inLeftCol {
		return []uint{i + 1}
	}
	if inTopRow && inBottomRow && inRightCol {
		return []uint{i - 1}
	}
	if inTopRow && inLeftCol && inRightCol {
		return []uint{i + c.width}
	}
	if inBottomRow && inLeftCol && inRightCol {
		return []uint{i - c.width}
	}
	if inTopRow && inBottomRow {
		return []uint{i - 1, i + 1}
	}
	if inLeftCol && inRightCol {
		return []uint{i - c.width, i + c.width}
	}
	if inTopRow && inLeftCol {
		return []uint{i + 1, i + c.width}
	}
	if inTopRow && inRightCol {
		return []uint{i - 1, i + c.width}
	}
	if inBottomRow && inLeftCol {
		return []uint{i + 1, i - c.width}
	}
	if inBottomRow && inRightCol {
		return []uint{i - 1, i - c.width}
	}
	if inTopRow {
		return []uint{i - 1, i + 1, i + c.width}
	}
	if inBottomRow {
		return []uint{i - 1, i + 1, i - c.width}
	}
	if inLeftCol {
		return []uint{i + 1, i - c.width, i + c.width}
	}
	if inRightCol {
		return []uint{i - 1, i - c.width, i + c.width}
	}
	return []uint{i - 1, i + 1, i - c.width, i + c.width}
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

	if c.width-1 == 0 {
		return true
	}
	if i == 0 {
		return false
	}

	return (i % c.width) == c.width-1
}
