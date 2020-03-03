package geometry

import (
	"github.com/shureban/quadrangle/lib/math"
	"github.com/shureban/quadrangle/structs/coordinate"
)

type Quadrangle struct {
	A coordinate.Coordinate // Координата точки A
	B coordinate.Coordinate // Координата точки B
	C coordinate.Coordinate // Координата точки C
	D coordinate.Coordinate // Координата точки D

	DiagonalBD float64 // Диагональ BD
	DiagonalAC float64 // Диагональ AC

	DiagonalCrossPoint coordinate.Coordinate // Координата пересечения диагоналей

	LenAB float64 // Длина стороны AB
	LenBC float64 // Длина стороны BC
	LenCD float64 // Длина стороны CD
	LenAD float64 // Длина стороны AD

	DegreeABC float64 // Градус угла ABC
	DegreeBCD float64 // Градус угла BCD
	DegreeCDA float64 // Градус угла CDA
	DegreeDAB float64 // Градус угла DAB

	TriangleABC *Triangle // Вписанный треугольник ABC
	TriangleCDA *Triangle // Вписанный треугольник CDA
}

/*
   B _________________ C
	|                 |
	|        O        |
   A|_________________|D
 */
func NewQuadrangle(A, B, C, D coordinate.Coordinate) *Quadrangle {
	q := &Quadrangle{A: A, B: B, C: C, D: D}

	q.DiagonalAC = math.VectorLength(coordinate.Vector{A: q.A, B: q.C})
	q.DiagonalBD = math.VectorLength(coordinate.Vector{A: q.B, B: q.D})

	q.DiagonalCrossPoint = math.VectorsCrossPoint(coordinate.Vector{A: q.B, B: q.D}, coordinate.Vector{A: q.A, B: q.C})

	q.LenAB = math.VectorLength(coordinate.Vector{A: q.A, B: q.B})
	q.LenBC = math.VectorLength(coordinate.Vector{A: q.B, B: q.C})
	q.LenCD = math.VectorLength(coordinate.Vector{A: q.C, B: q.D})
	q.LenAD = math.VectorLength(coordinate.Vector{A: q.A, B: q.D})

	q.DegreeABC = math.AngleDegree(q.LenAB, q.LenBC, q.DiagonalAC)
	q.DegreeBCD = math.AngleDegree(q.LenBC, q.LenCD, q.DiagonalBD)
	q.DegreeCDA = math.AngleDegree(q.LenCD, q.LenAD, q.DiagonalAC)
	q.DegreeDAB = math.AngleDegree(q.LenAD, q.LenAB, q.DiagonalBD)

	q.TriangleABC = NewTriangle(q.A, q.B, q.C)
	q.TriangleCDA = NewTriangle(q.C, q.D, q.A)

	return q
}
