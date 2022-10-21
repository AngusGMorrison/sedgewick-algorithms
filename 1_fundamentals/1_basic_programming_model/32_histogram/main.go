package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

const (
	imageSize    = 1024
	plotPadding  = 100
	barPadding   = 1
	axisPadding  = 2
	labelPadding = 24
	fontFace     = "/Library/Fonts/Arial Unicode.ttf"
	output       = "32_histogram/histogram.png"
	labelPoints  = 16
	titlePoints  = 20
)

var (
	barColor        = color.Black
	borderColor     = color.White
	backgroundColor = color.White
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	n := flag.Uint("n", 1, "The number of buckets comprising the histogram.")
	l := flag.Float64("l", 0, "The lower bound of the open interval to be plotted.")
	r := flag.Float64("r", 1, "The upper bound of the open interval to be plotted.")
	flag.Parse()

	h, err := newHistogram(*n, *l, *r)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return err
		}

		h.add(f)
	}

	p := newPlotter(h)
	return p.plot()
}

// histogram wraps a slice of n "buckets". At each index is the count of elements in data
// that fall into that bucket.
type histogram struct {
	data         []float64
	nBuckets     uint
	buckets      []int
	bSize        float64
	lower, upper float64
}

var (
	errNoBuckets    = errors.New("at least one bucket is required")
	errInvalidRange = errors.New("range upper bound must be > lower bound")
)

func newHistogram(n uint, lower, upper float64) (histogram, error) {
	if n == 0 {
		return histogram{}, errNoBuckets
	}
	if lower >= upper {
		return histogram{}, errInvalidRange
	}

	h := histogram{
		nBuckets: n,
		buckets:  make([]int, n),
		lower:    lower,
		upper:    upper,
	}

	return h, nil
}

func (h *histogram) bucketSize() float64 {
	if h.bSize != 0 {
		return h.bSize
	}

	h.bSize = (h.upper - h.lower) / float64(h.nBuckets)

	return h.bSize
}

func (h *histogram) bucketFor(v float64) int {
	if v <= h.lower || v >= h.upper {
		return -1
	}

	delta := v - h.lower
	return int(delta / h.bucketSize())
}

func (h *histogram) add(data ...float64) {
	for _, v := range data {
		bucket := h.bucketFor(v)
		if bucket == -1 {
			continue
		}

		h.buckets[bucket]++
	}
}

func (h *histogram) max() int {
	max := math.MinInt
	for _, count := range h.buckets {
		if count > max {
			max = count
		}
	}

	return max
}

func (h *histogram) labelForBucket(n int) string {
	return fmt.Sprintf("%.2f", h.lower+float64(n)*h.bucketSize())
}

type palette struct {
	axis, background, bar, border, text color.Color
}

type plotter struct {
	histogram    histogram
	ctx          *gg.Context
	imageSize    float64
	plotPadding  float64
	axisPadding  float64
	barPadding   float64
	labelPadding float64
	fontFace     string
	labelPoints  float64
	titlePoints  float64
	output       string
	palette      *palette
}

func newPlotter(h histogram) *plotter {
	return &plotter{
		histogram:    h,
		imageSize:    imageSize,
		plotPadding:  plotPadding,
		axisPadding:  axisPadding,
		barPadding:   barPadding,
		labelPadding: labelPadding,
		fontFace:     fontFace,
		labelPoints:  labelPoints,
		titlePoints:  titlePoints,
		output:       output,
		palette: &palette{
			axis:       color.Black,
			background: color.White,
			bar:        color.Black,
			border:     color.White,
			text:       color.Black,
		},
	}
}

func (p *plotter) plot() error {
	p.ctx = gg.NewContext(int(p.imageSize), int(p.imageSize))

	labelFont, err := gg.LoadFontFace(p.fontFace, p.labelPoints)
	if err != nil {
		return err
	}

	titleFont, err := gg.LoadFontFace(p.fontFace, p.titlePoints)
	if err != nil {
		return err
	}

	p.drawBackground()
	p.drawAxes()
	p.drawBars()
	p.drawLabels(labelFont)
	p.drawTitles(titleFont)

	return p.save()
}

func (p *plotter) leftMargin() float64 {
	return p.plotPadding
}

func (p *plotter) rightMargin() float64 {
	return p.imageSize - p.plotPadding
}

func (p *plotter) topMargin() float64 {
	return p.plotPadding
}

func (p *plotter) bottomMargin() float64 {
	return p.imageSize - p.plotPadding
}

func (p *plotter) xAxisLabelYPos() float64 {
	return p.bottomMargin() + p.labelPadding
}

func (p *plotter) yAxisLabelXPos() float64 {
	return p.leftMargin() - p.labelPadding
}

func (p *plotter) xAxisTitleXPos() float64 {
	return p.imageSize / 2
}

func (p *plotter) xAxisTitleYPos() float64 {
	return p.bottomMargin() + p.plotPadding/1.5
}

func (p *plotter) yAxisTitleXPos() float64 {
	return p.plotPadding / 4
}

func (p *plotter) yAxisTitleYPos() float64 {
	return p.imageSize / 2
}

func (p *plotter) plotSize() float64 {
	return p.imageSize - p.plotPadding*2
}

func (p *plotter) totalBarPaddingWidth() float64 {
	return float64(p.histogram.nBuckets)*p.barPadding + p.axisPadding
}

func (p *plotter) nBars() float64 {
	return float64(p.histogram.nBuckets)
}

func (p *plotter) barWidth() float64 {
	return (p.plotSize() - p.totalBarPaddingWidth()) / p.nBars()
}

func (p *plotter) barHeight(nElems int) float64 {
	return float64(nElems) * p.barUnitHeight()
}

func (p *plotter) barUnitHeight() float64 {
	return (p.plotSize() - p.axisPadding) / float64(p.histogram.max())
}

func (p *plotter) barX(barNumber int) float64 {
	return p.leftMargin() + p.axisPadding + float64(barNumber)*(p.barWidth()+p.barPadding)
}

func (p *plotter) barY(nElems int) float64 {
	return p.bottomMargin() - p.axisPadding - p.barHeight(nElems)
}

func (p *plotter) drawBackground() {
	p.ctx.SetColor(p.palette.background)
	p.ctx.Clear()
}

func (p *plotter) drawAxes() {
	p.ctx.SetColor(p.palette.axis)
	p.ctx.DrawLine(p.leftMargin()-10, p.bottomMargin(), p.rightMargin()+10, p.bottomMargin())
	p.ctx.DrawLine(p.leftMargin(), p.bottomMargin()+10, p.leftMargin(), p.topMargin()-10)
	p.ctx.Stroke()
}

func (p *plotter) drawBars() {
	for i, count := range p.histogram.buckets {
		p.ctx.DrawRectangle(p.barX(i), p.barY(count), p.barWidth(), p.barHeight(count))
	}
	p.ctx.SetColor(p.palette.bar)
	p.ctx.FillPreserve()
	p.ctx.SetColor(p.palette.border)
	p.ctx.Stroke()
}

func (p *plotter) drawLabels(fontFace font.Face) {
	p.ctx.SetFontFace(fontFace)
	p.drawXLabels()
	p.drawYLabels()
}

func (p *plotter) drawXLabels() {
	p.ctx.SetColor(p.palette.text)

	for i := range p.histogram.buckets {
		p.ctx.DrawStringAnchored(p.histogram.labelForBucket(i), p.barX(i), p.xAxisLabelYPos(), 0.5, 0.5)
	}
	upperBound := int(p.histogram.nBuckets)
	p.ctx.DrawStringAnchored(p.histogram.labelForBucket(upperBound), p.barX(upperBound), p.xAxisLabelYPos(), 0.5, 0.5)
}

func (p *plotter) drawYLabels() {
	p.ctx.SetColor(p.palette.text)
	for i := 0; i <= p.histogram.max(); i++ {
		label := strconv.Itoa(i)
		p.ctx.DrawStringAnchored(label, p.yAxisLabelXPos(), p.barY(i), 1, 0.5)
	}
}

func (p *plotter) drawTitles(fontFace font.Face) {
	p.ctx.SetFontFace(fontFace)
	p.drawXTitle()
	p.drawYTitle()
}

func (p *plotter) drawXTitle() {
	p.ctx.DrawStringAnchored("Range", p.xAxisTitleXPos(), p.xAxisTitleYPos(), 0.5, 0.5)
}

func (p *plotter) drawYTitle() {
	p.ctx.RotateAbout(-math.Pi/2, p.yAxisTitleXPos(), p.yAxisTitleYPos())
	p.ctx.DrawStringAnchored("Frequency", p.yAxisTitleXPos(), p.yAxisTitleYPos(), 0.5, 0.5)
	p.ctx.RotateAbout(math.Pi/2, p.yAxisTitleXPos(), p.yAxisTitleYPos())
}

func (p *plotter) save() error {
	return p.ctx.SavePNG(p.output)
}
