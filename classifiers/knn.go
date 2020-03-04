package classifiers

import (
	"github.com/pa-m/sklearn/neighbors"
	"github.com/shureban/quadrangle/datasets"
	"github.com/shureban/quadrangle/lib/converter"
	"github.com/shureban/quadrangle/lib/math"
	"github.com/shureban/quadrangle/structs/coordinate"
	"github.com/shureban/quadrangle/structs/geometry"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

const (
	AnalyseParamsCount = 21
	NumberOfNeighbors  = 1
)

var KNNClassifier *neighbors.KNeighborsClassifier

// При инициализации компонента
// Создаем модель классификатора
// И обучаем его на подготовленной выборке данных
func init() {
	matrix  := mat.NewDense(len(datasets.QuadrangleData), AnalyseParamsCount, nil)
	classes := mat.NewDense(len(datasets.QuadrangleData), 1, nil)

	for key, list := range datasets.QuadrangleData {
		matrix.SetRow(key, list[1:])
		classes.SetRow(key, list[0:1])
	}

	KNNClassifier = neighbors.NewKNeighborsClassifier(NumberOfNeighbors, "distance")
	KNNClassifier.Fit(matrix, classes)
}

// Возвращает наиболее вероятный класс
func predict(data []float64) int {
	y := mat.NewDense(1, 1, nil)
	X := mat.NewDense(1, AnalyseParamsCount, data)

	KNNClassifier.Predict(X, y)

	return int(y.RawRowView(0)[0])
}

// Возвращает наиболее вероятный класс четырехугольника
func PredictQuadrangle(q *geometry.Quadrangle) int {
	data := quadrangleAnalyseDataset(q)

	return predict(data)
}

// Возвращает набор параметров для анализа четырехугольника
func quadrangleAnalyseDataset(q *geometry.Quadrangle) []float64 {
	return []float64{
		converter.BoolToFloat64(math.IsParallel(coordinate.Vector{A: q.A, B: q.B}, coordinate.Vector{A: q.C, B: q.D})),
		converter.BoolToFloat64(math.IsParallel(coordinate.Vector{A: q.C, B: q.B}, coordinate.Vector{A: q.A, B: q.D})),
		converter.BoolToFloat64(q.LenAB == q.LenBC),
		converter.BoolToFloat64(q.LenCD == q.LenAD),
		converter.BoolToFloat64(q.LenAB == q.LenCD),
		converter.BoolToFloat64(q.LenBC == q.LenAD),
		converter.BoolToFloat64(floats.Round(q.DegreeABC, 4) == 90),
		converter.BoolToFloat64(floats.Round(q.DegreeBCD, 4) == 90),
		converter.BoolToFloat64(floats.Round(q.DegreeCDA, 4) == 90),
		converter.BoolToFloat64(floats.Round(q.DegreeDAB, 4) == 90),
		converter.BoolToFloat64(floats.Round(q.DegreeABC, 4) > 0),
		converter.BoolToFloat64(floats.Round(q.DegreeBCD, 4) > 0),
		converter.BoolToFloat64(floats.Round(q.DegreeCDA, 4) > 0),
		converter.BoolToFloat64(floats.Round(q.DegreeDAB, 4) > 0),
		converter.BoolToFloat64(q.DiagonalAC == q.DiagonalBD),
		converter.BoolToFloat64(q.TriangleABC.Hypotenuse == q.TriangleCDA.Hypotenuse),
		converter.BoolToFloat64(q.TriangleABC.LegAB == q.TriangleCDA.LegAB),
		converter.BoolToFloat64(floats.Round(q.TriangleABC.DegreeBCA, 4) == 45),
		converter.BoolToFloat64(floats.Round(q.TriangleABC.DegreeCAB, 4) == 45),
		converter.BoolToFloat64(floats.Round(q.TriangleCDA.DegreeBCA, 4) == 45),
		converter.BoolToFloat64(floats.Round(q.TriangleCDA.DegreeCAB, 4) == 45),
	}
}
