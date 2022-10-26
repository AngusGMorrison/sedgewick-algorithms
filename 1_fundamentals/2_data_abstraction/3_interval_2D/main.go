package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

const (
	scaleMin   = 0
	scaleMax   = 1
	imageSize  = 1024
	outputPath = "output.png"
)

func main() {
	n := flag.Uint("n", 1, "The number of intervals to draw.")
	min := flag.Float64("min", 0, "The minimum width and height of a rectangular interval. Must be <= max.")
	max := flag.Float64("max", 1, "The maximum width and height of a rectangular interval. Must be >= min.")
	flag.Parse()

	if *min < 0 || *max < 0 || *min > *max {
		flag.PrintDefaults()
		os.Exit(1)
	}

	src := rand.NewSource(time.Now().UnixNano())
	randGen := rng{rand.New(src)}
	rectGen := newRectangleGenerator(*min, *max, scaleMin, scaleMax, &randGen)
	rects := rectGen.generateN(int(*n))
	printStats(rects)
	if err := drawRectangles(rects, &randGen); err != nil {
		log.Fatalln(err)
	}
}

type point struct {
	x, y float64
}

func (p point) scale(factor float64) point {
	return point{
		x: p.x * factor,
		y: p.y * factor,
	}
}

func (p point) leftOf(q point) bool {
	return p.x <= q.x
}

func (p point) rightOf(q point) bool {
	return p.x >= q.x
}

func (p point) above(q point) bool {
	return p.y <= q.y
}

func (p point) below(q point) bool {
	return p.y >= q.y
}

type rectangle struct {
	topLeft, bottomRight point
}

func (r rectangle) width() float64 {
	return r.bottomRight.x - r.topLeft.x
}

func (r rectangle) height() float64 {
	return r.bottomRight.y - r.topLeft.y
}

func (r rectangle) scale(factor float64) rectangle {
	return rectangle{
		topLeft:     r.topLeft.scale(factor),
		bottomRight: r.bottomRight.scale(factor),
	}
}

func (r rectangle) intersects(s rectangle) bool {
	return r.verticallyOverlaps(s) && r.horizontallyOverlaps(s)
}

func (r rectangle) verticallyOverlaps(s rectangle) bool {
	return r.bottomRight.below(s.topLeft) && r.topLeft.above(s.bottomRight)
}

func (r rectangle) horizontallyOverlaps(s rectangle) bool {
	return r.topLeft.leftOf(s.bottomRight) && r.bottomRight.rightOf(s.topLeft)
}

func (r rectangle) contains(s rectangle) bool {
	return r.topLeft.above(s.topLeft) &&
		r.topLeft.leftOf(s.topLeft) &&
		r.bottomRight.below(s.bottomRight) &&
		r.bottomRight.rightOf(s.bottomRight)
}

// rng is a wrapper around *rand.Rand that adds convenience functions for generating random floats
// in a range and random RGBA colors.
type rng struct {
	*rand.Rand
}

func (r *rng) float64InRange(min, max float64) float64 {
	return min + r.Float64()*(max-min)
}

func (r *rng) color() color.RGBA {
	return color.RGBA{
		R: uint8(r.Intn(256)),
		G: uint8(r.Intn(256)),
		B: uint8(r.Intn(256)),
		A: 255,
	}
}

// rectangleGenerator randomly generates rectangles whose origins and dimensions are normalized to
// the unit square.
type rectangleGenerator struct {
	normalizedMin float64
	normalizedMax float64
	rng           *rng
}

func newRectangleGenerator(measureMin, measureMax, scaleMin, scaleMax float64, r *rng) *rectangleGenerator {
	normalizer := normalizer{
		measureMin: measureMin,
		measureMax: measureMax,
		scaleMin:   scaleMin,
		scaleMax:   scaleMax,
	}

	return &rectangleGenerator{
		normalizedMin: normalizer.normalize(measureMin),
		normalizedMax: normalizer.normalize(measureMax),
		rng:           r,
	}
}

func (urg *rectangleGenerator) generateN(n int) []rectangle {
	rects := make([]rectangle, n)
	for i := 0; i < n; i++ {
		rects[i] = urg.generate()
	}

	return rects
}

func (urg *rectangleGenerator) generate() rectangle {
	x1, y1 := urg.rng.Float64(), urg.rng.Float64()
	x2 := x1 + urg.rng.float64InRange(urg.normalizedMin, urg.normalizedMax)
	y2 := y1 + urg.rng.float64InRange(urg.normalizedMin, urg.normalizedMax)

	return rectangle{
		topLeft:     point{x1, y1},
		bottomRight: point{x2, y2},
	}
}

type normalizer struct {
	measureMin, measureMax float64
	scaleMin, scaleMax     float64
}

// normalize normalizes a value by calculating the value as a fraction of the full measurement
// range, then scaling it by the target range (which is 1 in the case of the unit measure) and
// offsetting it by the minimum value of the target range (0 for the unit measure).
func (n *normalizer) normalize(v float64) float64 {
	return (v-n.measureMin)/n.measureRange()*n.scaleRange() + n.scaleMin
}

func (n *normalizer) measureRange() float64 {
	return n.measureMax - n.measureMin
}

func (n *normalizer) scaleRange() float64 {
	return n.scaleMax - n.scaleMin
}

func printStats(rects []rectangle) {
	var isxns, contained int
	for i := 0; i < len(rects); i++ {
		for j := i + 1; j < len(rects); j++ {
			if rects[i].contains(rects[j]) {
				contained++
				continue
			}

			if rects[i].intersects(rects[j]) {
				isxns++
			}
		}
	}

	fmt.Printf("Intersections: %d, Contained: %d\n", isxns, contained)
}

func drawRectangles(rects []rectangle, rGen *rng) error {
	ctx := gg.NewContext(imageSize, imageSize)
	ctx.SetColor(color.White)
	ctx.Clear()

	for _, r := range rects {
		scaled := r.scale(imageSize)
		ctx.DrawRectangle(scaled.topLeft.x, scaled.topLeft.y, scaled.width(), scaled.height())
		ctx.SetColor(rGen.color())
		ctx.Stroke()
	}

	return ctx.SavePNG(outputPath)
}
