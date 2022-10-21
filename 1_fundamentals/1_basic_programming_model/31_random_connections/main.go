package main

import (
	"flag"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	n := flag.Int("n", 0, "The number of points to plot.")
	p := flag.Float64("p", 1, "The probability with which each pair of points should be joined by a line.")
	flag.Parse()

	plot(*n, *p)
}

func plot(n int, p float64) {
	const (
		size   = 1024
		radius = 400
	)

	context := gg.NewContext(size, size)
	// Background.
	context.DrawRectangle(0, 0, size, size)
	context.SetColor(color.White)
	context.Fill()

	// Plot points.
	context.SetColor(color.Black)
	points := make([]gg.Point, 0, n)
	angle := (2 * math.Pi) / float64(n)
	for i := 0; i < n; i++ {
		x := size/2 + math.Cos(angle*float64(i))*radius
		y := size/2 + math.Sin(angle*float64(i))*radius
		points = append(points, gg.Point{X: x, Y: y})
		context.DrawPoint(x, y, 5)
	}
	context.Fill()

	// Plot lines.
	context.SetColor(color.RGBA{R: 150, G: 150, B: 150, A: 255})
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if rand.Float64() <= p {
				context.DrawLine(points[i].X, points[i].Y, points[j].X, points[j].Y)
			}
		}
	}
	context.Stroke()

	context.SavePNG("out.png")
}
