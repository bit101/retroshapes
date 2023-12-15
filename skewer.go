// Package retroshapes draws retro shapes
package retroshapes

import (
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
)

// Skewer is a line with circles on the end.
type Skewer struct {
	x0        float64
	y0        float64
	x1        float64
	y1        float64
	width     float64
	endRadius float64
}

// NewSkewer creates a new skewer object.
func NewSkewer(x0, y0, x1, y1, width, endRadius float64) *Skewer {
	return &Skewer{x0, y0, x1, y1, width, endRadius}
}

// Draw draws the skewer on the given context.
func (s *Skewer) Draw(context *cairo.Context) {
	context.Save()
	context.SetLineWidth(s.width)
	context.MoveTo(s.x0, s.y0)
	context.LineTo(s.x1, s.y1)
	context.Stroke()
	context.FillCircle(s.x0, s.y0, s.endRadius)
	context.FillCircle(s.x1, s.y1, s.endRadius)
	context.Restore()
}

// Randomize randomized the position of the skewer.
func (s *Skewer) Randomize(rand float64) {
	s.x0 += random.FloatRange(-rand, rand)
	s.y0 += random.FloatRange(-rand, rand)
	s.x1 += random.FloatRange(-rand, rand)
	s.y1 += random.FloatRange(-rand, rand)
}
