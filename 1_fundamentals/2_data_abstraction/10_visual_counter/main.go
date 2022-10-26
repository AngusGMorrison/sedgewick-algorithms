package main

import (
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	nOps     = 100
	maxOps   = 80
	maxCount = 2
	plotFile = "points.png"
	plotSize = 4 * vg.Inch
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	c := newCounter(maxOps, maxCount)
	for i := 0; i < nOps; i++ {
		if rng.Intn(2) == 1 {
			c.inc()
			continue
		}

		c.dec()
	}

	if err := c.plotCount(); err != nil {
		log.Fatalln(err)
	}
}

type visualCounter struct {
	count, maxCount int
	ops, maxOps     int
	points          plotter.XYs
}

func newCounter(maxOps, maxCount int) *visualCounter {
	return &visualCounter{
		maxOps:   maxOps,
		maxCount: maxCount,
	}
}

func (c *visualCounter) inc() {
	defer c.capturePoint()
	defer c.incOps()

	if c.maxOpsReached() || c.maxCountReached() {
		return
	}

	c.count++
}

func (c *visualCounter) dec() {
	defer c.capturePoint()
	defer c.incOps()

	if c.maxOpsReached() || c.minCountReached() {
		return
	}

	c.count--
}

func (c *visualCounter) incOps() {
	c.ops++
}

func (c *visualCounter) maxOpsReached() bool {
	return c.ops >= c.maxOps
}

func (c *visualCounter) maxCountReached() bool {
	return c.count >= c.maxCount
}

func (c *visualCounter) minCountReached() bool {
	return c.count <= -c.maxCount
}

func (c *visualCounter) capturePoint() {
	c.points = append(c.points, plotter.XY{X: float64(c.ops), Y: float64(c.count)})
}

func (c *visualCounter) plotCount() error {
	p := plot.New()
	p.Title.Text = "(*visualCounter).count"
	p.X.Label.Text = "Op number"
	p.Y.Label.Text = "Count"

	if err := plotutil.AddLinePoints(p, "Count", c.points); err != nil {
		return err
	}

	return p.Save(plotSize, plotSize, plotFile)
}
