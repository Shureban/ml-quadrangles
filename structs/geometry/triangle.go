package geometry

import (
	"github.com/shureban/quadrangle/lib/math"
	"github.com/shureban/quadrangle/structs/coordinate"
)

type Triangle struct {
	A coordinate.Coordinate
	B coordinate.Coordinate
	C coordinate.Coordinate

	LegBC      float64
	LegAB      float64
	Hypotenuse float64

	DegreeABC float64
	DegreeBCA float64
	DegreeCAB float64
}

/*
	C               C
	|\             /\
	| \           /  \
	|  \         /    \
	|   \       /      \
   B|____\A   B/________\A
 */
func NewTriangle(A, B, C coordinate.Coordinate) *Triangle {
	t := &Triangle{A: A, B: B, C: C}

	t.LegAB = math.VectorLength(coordinate.Vector{A: A, B: B})
	t.LegBC = math.VectorLength(coordinate.Vector{A: C, B: B})
	t.Hypotenuse = math.VectorLength(coordinate.Vector{A: A, B: C})

	t.DegreeABC = math.AngleDegree(t.LegAB, t.LegBC, t.Hypotenuse)
	t.DegreeBCA = math.AngleDegree(t.LegBC, t.Hypotenuse, t.LegAB)
	t.DegreeCAB = math.AngleDegree(t.Hypotenuse, t.LegAB, t.LegBC)

	return t
}
