package tree

import (
	"github.com/shureban/quadrangle/structs/coordinate"
)

type Root struct {
	Coordinates coordinate.List
	Nodes       NodeList
	MaxDeep     int
}

// Возвращает объект нового дерева
func NewTree(maxDeep int) *Root {
	return &Root{
		Coordinates: coordinate.List{},
		Nodes:       NodeList{},
		MaxDeep:     maxDeep,
	}
}

// Заполняет дерево элементами
func (r *Root) Feet(list coordinate.List) *Root {
	r.Coordinates = list

	// Если элементов меньше MaxDeep
	// Значит не получится создать ни одной фигуры состоящей из MaxDeep координат
	// Возвращаем пустое дерево
	if len(list) < r.MaxDeep {
		return r
	}

	// Начинаем разворачивать список
	// Начальная глубина 1
	r.Nodes = NewNodeList(r.Coordinates, 1, r.MaxDeep)

	return r
}

// Все листья сворачиваются в множество списков координат
func (r *Root) Collapse() (nlist []coordinate.List) {
	for _, node := range r.Nodes {
		nlist = append(nlist, node.Collapse()...)
	}

	return nlist
}
