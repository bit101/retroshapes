// Package retroshapes draws retro shapes
package retroshapes

import (
	"math"

	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
)

// Atom is an atom shape.
type Atom struct {
	x, y           float64
	nucleusRadius  float64
	ringRadius     float64
	ringRatio      float64
	electronRadius float64
	numElectrons   int
	rotation       float64
}

// NewAtom returns a new Atom object.
func NewAtom(x, y, nucleusRadius, ringRadius, ringRatio, electronRadius float64, numElectrons int) *Atom {
	return &Atom{x, y, nucleusRadius, ringRadius, ringRatio, electronRadius, numElectrons, 0.0}
}

// Draw draws this atom to the context.
func (a *Atom) Draw(context *cairo.Context) {
	count := float64(a.numElectrons)
	context.Save()
	context.Translate(a.x, a.y)
	context.FillCircle(0, 0, a.nucleusRadius)
	for i := 0.0; i < count; i++ {
		context.Save()
		context.Rotate(i/count*math.Pi + a.rotation)
		context.StrokeEllipse(0, 0, a.ringRadius, a.ringRadius*a.ringRatio)
		if a.electronRadius > 0 {
			angle := random.Angle()
			context.FillCircle(math.Cos(angle)*a.ringRadius, math.Sin(angle)*a.ringRadius*a.ringRatio, a.electronRadius)
		}
		context.Restore()
	}

	context.Restore()
}

// Rotate rotates the atom.
func (a *Atom) Rotate(angle float64) {
	a.rotation = angle
}
