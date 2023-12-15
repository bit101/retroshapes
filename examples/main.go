// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/retroshapes"
)

func main() {
	// renderFrame := skewerBlobs
	// renderFrame := spokes
	// renderFrame := skeweredSpokes
	renderFrame := atoms

	render.Image(900, 900, "out/out.png", renderFrame, 0.0)
	render.ViewImage("out/out.png")
}

//revive:disable-next-line:unused-parameter
func skewerBlobs(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	for x := width / 4; x < width; x += width / 4 {
		context.SetLineWidth(3)

		s := retroshapes.NewSkewer(x, 50, x, height-50, 3, 5)
		s.Randomize(20)
		s.Draw(context)

		context.SetLineWidth(0.75)

		points := s.GetPoints(6, 100, 100)
		for _, p := range points {
			size := random.FloatRange(50, 100)
			rand := size * 0.2
			rotation := random.FloatRange(-0.4, 0.4)

			blob := retroshapes.NewBlob(p.X, p.Y, size, rotation, rand)
			context.SetSourceGray(random.FloatRange(0.4, 0.9))
			blob.Fill(context)

			blob.Randomize(size * 0.1)
			context.SetSourceBlack()
			blob.Stroke(context)
		}
	}
}

//revive:disable-next-line:unused-parameter
func spokes(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.75)

	for x := 0.0; x <= width; x += 150 {
		for y := 0.0; y <= height; y += 150 {
			numSpokes := random.IntRange(6, 20)
			spokeLength := random.FloatRange(20, 80)
			endRadius := spokeLength * random.FloatRange(0.01, 0.15)
			innerRadius := spokeLength * random.FloatRange(0.0, 0.4)
			fill := random.Boolean()
			spokes := retroshapes.NewRegularSpokes(x, y, numSpokes, spokeLength, endRadius, innerRadius)
			// spokes.RandomizeLengths(5)
			spokes.RandomizeAngles(1)
			// spokes := retroshapes.NewRandomSpokes(x, y, numSpokes, 50, 80, 0.5, endRadius, innerRadius)
			spokes.Draw(context, fill)
		}
	}
}

//revive:disable-next-line:unused-parameter
func skeweredSpokes(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.75)

	s := retroshapes.NewSkewer(50, 850, 850, 50, 1, 4)
	s.Draw(context)

	points := s.GetPoints(8, 100, 100)
	for _, p := range points {
		p.X += random.FloatRange(-10, 10)
		p.Y += random.FloatRange(-10, 10)
		spokeLength := random.FloatRange(30, 70)
		numSpokes := random.IntRange(6, 16)
		spokes := retroshapes.NewRegularSpokes(p.X, p.Y, numSpokes, spokeLength, 5, 20)
		spokes.RandomizeLengths(10)
		context.SetSourceHSV(random.FloatRange(0, 360), 0.5, 1)
		spokes.FillCircles(context, 5, 3)
		context.SetSourceBlack()
		spokes.Draw(context, false)

	}
}

//revive:disable-next-line:unused-parameter
func atoms(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineWidth(0.75)

	for x := 0.0; x <= width; x += 150 {
		for y := 0.0; y <= height; y += 150 {
			nRadius := random.FloatRange(3, 20)
			rRadius := math.Min(70, nRadius*random.FloatRange(4, 6))
			rRatio := random.FloatRange(0.25, 0.6)
			eRadius := nRadius * random.FloatRange(0.1, 0.25)
			numElectrons := random.IntRange(2, 5)
			a := retroshapes.NewAtom(x, y, nRadius, rRadius, rRatio, eRadius, numElectrons)
			a.Rotate(random.Angle())
			a.Draw(context)
		}
	}
}
