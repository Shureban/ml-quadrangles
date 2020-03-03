package math

import (
	"github.com/shureban/quadrangle/structs/coordinate"
	"gonum.org/v1/gonum/floats"
	"math"
)

// Высчитывает длинну отрезка AB
func VectorLength(vector coordinate.Vector) float64 {
	X := math.Pow(vector.B.X - vector.A.X, 2)
	Y := math.Pow(vector.B.Y - vector.A.Y, 2)

	return math.Sqrt(X + Y)
}

// Вычисляет точку пересечения 2-х векторов
// {-Inf NaN} - если вектора параллельны друг другу или точка пересечения действительно в 0:0
func VectorsCrossPoint(v1, v2 coordinate.Vector) coordinate.Coordinate {
	x1, x2, x3, x4 := v1.A.X, v1.B.X, v2.A.X, v2.B.X
	y1, y2, y3, y4 := v1.A.Y, v1.B.Y, v2.A.Y, v2.B.Y
	denominator    := (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)

	if denominator == 0 {
		return coordinate.Coordinate{X: 0, Y: 0}
	}

	return coordinate.Coordinate{
		X: floats.Round(((x1*y2 - y1*x2) * (x3 - x4) - (x1 - x2) * (x3*y4 - y3*x4)) / denominator, 2),
		Y: floats.Round(((x1*y2 - y1*x2) * (y3 - y4) - (y1 - y2) * (x3*y4 - y3*x4)) / denominator, 2),
	}
}

// Проверка параллельности векторов
func IsParallel(v1, v2 coordinate.Vector) bool {
	x1, x2, x3, x4 := v1.A.X, v1.B.X, v2.A.X, v2.B.X
	y1, y2, y3, y4 := v1.A.Y, v1.B.Y, v2.A.Y, v2.B.Y

	return ((x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)) == 0
}

// Высчитывает градусы угла
func AngleDegree(sideA float64, sideB float64, sideC float64) float64 {
	numerator   := math.Pow(sideA, 2) + math.Pow(sideB, 2) - math.Pow(sideC, 2)
	denominator := 2 * sideA * sideB

	return math.Acos(numerator / denominator) * 180 / math.Pi
}
