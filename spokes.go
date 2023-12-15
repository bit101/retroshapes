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

// NewSpokes creates a new Spokes object.
// func NewSpokes(x, y, minRadius, maxRadius, endRadius, rotation float64, numSpokes int) *Spokes {
// 	points := geom.NewPointList()
// 	ns := float64(numSpokes)
// 	for i := 0.0; i < ns; i++ {
// 		radius := random.FloatRange(minRadius, maxRadius)
// 		angle := i / ns * math.Pi
// 		points.AddXY(math.Cos(angle)*radius, math.Sin(angle)*radius)
// 		radius = random.FloatRange(minRadius, maxRadius)
// 		points.AddXY(-math.Cos(angle)*radius, -math.Sin(angle)*radius)
// 	}
// 	points.Rotate(rotation)
// 	points.Translate(x, y)
// 	return &Spokes{points, endRadius}
// }

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
