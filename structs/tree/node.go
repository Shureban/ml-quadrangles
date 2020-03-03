package tree

import . "github.com/shureban/quadrangle/structs/coordinate"

type NodeList []*Node
type Node struct {
	Deep  int
	Value Coordinate
	Nodes NodeList
}

// Возвращает список листьев, созданный из набора координат
func NewNodeList(list List, deep int, maxDeep int) (nodes NodeList) {
	for key, coordinate := range list {
		slice := list[key+1:]

		// Добавляем [0:key+1] итерируемого списка в конец
		// Чтобы получить все возможные варианты комбинаций координат
		if tail := list[:key]; len(tail) != 0 {
			slice = append(slice, tail...)
		}

		node := NewNode(coordinate, slice, deep, maxDeep)
		nodes = append(nodes, node)
	}

	return nodes
}

// Возвращает объект листа
// Пока не достигнута максимальная глубина
// Рекурсивно обходим список координат и создаем новые листья
func NewNode(value Coordinate, list List, deep int, maxDeep int) *Node {
	node := &Node{
		Deep:  deep,
		Value: value,
		Nodes: NodeList{},
	}

	// Пока не достигли максимальной глубины,
	// Продолжаем рекурсивно разворачивать список координат
	if node.Deep < maxDeep {
		node.Nodes = NewNodeList(list, node.Deep + 1, maxDeep)
	}

	return node
}

// Возвращает значения дочерних элементов в виде множества списков координат
func (n Node) Collapse() (lists []List) {
	// Если нет дочерних узлов
	// Возвращаем значение текущего листа
	if len(n.Nodes) == 0 {
		return []List{{n.Value}}
	}

	for _, node := range n.Nodes {
		lists = append(lists, node.Collapse()...)
	}

	for key, value := range lists {
		lists[key] = append(value, n.Value)
	}

	return lists
}

