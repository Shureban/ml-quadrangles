package main

import (
	"encoding/json"
	"fmt"
	"github.com/shureban/quadrangle/classifiers"
	"github.com/shureban/quadrangle/lib/uniq"
	. "github.com/shureban/quadrangle/structs/coordinate"
	. "github.com/shureban/quadrangle/structs/geometry"
	. "github.com/shureban/quadrangle/structs/tree"
	"io/ioutil"
	"log"
)

const (
	MaxTreeDeep = 4
)

func main() {
	coordinates := parseCoordinatesFile()
	tree        := NewTree(MaxTreeDeep).Feet(coordinates)
	quadrangles := quadranglesFromCoordinates(tree.Collapse())
	result      := map[string]int{"Squares": 0, "Rhombuses": 0, "Rectangles": 0, "Parallelograms": 0}

	for _, q := range quadrangles {
		switch classifiers.PredictQuadrangle(q) {
		case 1:
			result["Rectangles"]++
		case 2:
			result["Squares"]++
		case 3:
			result["Parallelograms"]++
		case 4:
			result["Rhombuses"]++
		}
	}
	fmt.Printf(`
 Figure         | Count  
----------------+--------
 Squares        | %d     
 Rhombuses      | %d     
 Rectangles     | %d     
 Parallelograms | %d     
`, result["Squares"], result["Rhombuses"], result["Rectangles"], result["Parallelograms"])
}

// Из списка равных списков точек
// Создается список уникальных четырехугольников
func quadranglesFromCoordinates(lists []List) (quadrangles []*Quadrangle) {
	for _, coordinates := range lists {
		quadrangle := NewQuadrangle(coordinates[0], coordinates[1], coordinates[2], coordinates[3])
		quadrangles = append(quadrangles, quadrangle)
	}

	return uniq.Quadrangles(quadrangles)
}

// Возвращает данные о координатах из файла coordinates.json
func parseCoordinatesFile() (list List) {
	file, err := ioutil.ReadFile("coordinates.json")

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(file, &list); err != nil {
		log.Fatal(err)
	}

	return list
}
