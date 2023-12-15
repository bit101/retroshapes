// Package retroshapes draws retro shapes
package retroshapes

import (
	"github.com/bit101/bitlib/geom"
	cairo "github.com/bit101/blcairo"
)

// Blob represents a blob structure.
type Blob struct {
	path geom.PointList
}

// NewBlob creates a blob path.
func NewBlob(x, y, size, rotation, rand float64) *Blob {
	s := size / 1
	path := geom.NewPointList()
	path.AddXY(0, 0)
	path.AddXY(-s, s)
	path.AddXY(0, -s)
	path.AddXY(s, s)
	path.Randomize(rand, rand)
	path.Rotate(rotation)
	path.Translate(x, y)
	path.Randomize(rand/2, rand/2)
	return &Blob{path}
}

// Fill fills a blob shape using the given context.
func (b *Blob) Fill(context *cairo.Context) {
	context.FillMultiLoop(b.path)
}

// Stroke strokes a blob shape using the given context.
func (b *Blob) Stroke(context *cairo.Context) {
	context.StrokeMultiLoop(b.path)
}

// Randomize randomizes the path that makes up a blob.
func (b *Blob) Randomize(rand float64) {
	b.path.Randomize(rand, rand)
}
