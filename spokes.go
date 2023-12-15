// Package retroshapes draws retro shapes
package retroshapes

import (
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
)

// Spoke represents a single spoke.
type Spoke struct {
	radius    float64
	angle     float64
	endRadius float64
}

// Spokes represents a shape with spokes extending from the center.
type Spokes struct {
	spokes      []*Spoke
	x           float64
	y           float64
	innerRadius float64
}

// NewRegularSpokes creates a new spokes object with uniform sized spokes.
func NewRegularSpokes(x, y float64, numSpokes int, spokeLength, endRadius, innerRadius float64) *Spokes {
	ns := float64(numSpokes)
	spokes := []*Spoke{}
	for i := 0.0; i < ns; i++ {
		spokes = append(spokes, &Spoke{spokeLength, i / ns * blmath.Tau, endRadius})
	}
	return &Spokes{spokes, x, y, innerRadius}
}

// NewRandomSpokes creates a new spokes object with random sized spokes.
func NewRandomSpokes(x, y float64, numSpokes int, minSpokeLength, maxSpokeLength, angleVariance, endRadius, innerRadius float64) *Spokes {
	ns := float64(numSpokes)
	spokes := []*Spoke{}
	for i := 0.0; i < ns; i++ {
		spokeLength := random.FloatRange(minSpokeLength, maxSpokeLength)

		angleVar := blmath.Tau / ns / 2 * random.FloatRange(-angleVariance, angleVariance)
		spokes = append(spokes, &Spoke{spokeLength, i/ns*blmath.Tau + angleVar, endRadius})
	}
	return &Spokes{spokes, x, y, innerRadius}
}

// Draw draws the spokes shape.
func (s *Spokes) Draw(context *cairo.Context, fill bool) {
	context.Save()
	context.Translate(s.x, s.y)
	for _, spoke := range s.spokes {
		context.Save()
		context.Rotate(spoke.angle)
		context.MoveTo(s.innerRadius, 0)
		context.LineTo(spoke.radius-spoke.endRadius, 0)
		context.Stroke()
		if fill {
			context.FillCircle(spoke.radius, 0, spoke.endRadius)
		} else {
			context.StrokeCircle(spoke.radius, 0, spoke.endRadius)
		}
		context.Restore()
	}
	if fill {
		context.FillCircle(0, 0, s.innerRadius)
	} else {
		context.StrokeCircle(0, 0, s.innerRadius)
	}
	context.Restore()
}

// FillCircles draws the spokes shape.
func (s *Spokes) FillCircles(context *cairo.Context, randCenter, randSpokes float64) {
	context.Save()
	context.Translate(s.x, s.y)
	for _, spoke := range s.spokes {
		context.Save()
		context.Rotate(spoke.angle)
		context.FillCircle(spoke.radius+random.FloatRange(-randSpokes, randSpokes), random.FloatRange(-randSpokes, randSpokes), spoke.endRadius)
		context.Restore()
	}
	context.FillCircle(random.FloatRange(-randCenter, randCenter), random.FloatRange(-randCenter, randCenter), s.innerRadius)
	context.Restore()
}

// RandomizeLengths randomizes the lengths of each spoke.
func (s *Spokes) RandomizeLengths(rand float64) {
	for _, spoke := range s.spokes {
		spoke.radius += random.FloatRange(-rand, rand)
	}
}

// RandomizeAngles randomizes the angles between spokes.
// The rand parameter should be between 0 and 1, zero being no randomization and one being maximum.
// Larger numbers can be used, but it gets chaotic.
func (s *Spokes) RandomizeAngles(rand float64) {
	ns := float64(len(s.spokes))
	for _, spoke := range s.spokes {
		angleVar := blmath.Tau / ns / 2 * random.FloatRange(-rand, rand)
		spoke.angle += angleVar
	}
}
