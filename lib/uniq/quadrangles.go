package uniq

import (
	"fmt"
	. "github.com/shureban/quadrangle/structs/geometry"
	"strings"
)

// Возвращает список уникальных четырехугольников из набора
func Quadrangles(list []*Quadrangle) (result []*Quadrangle) {
	uniqMap := map[string]*Quadrangle{}

	for _, q := range list {
		fKey   := fmt.Sprint(q.A, q.B, q.C, q.D)
		rKey   := fmt.Sprint(q.D, q.C, q.B, q.A)
		exists := false

		for key := range uniqMap {
			if strings.Contains(key, fKey) || strings.Contains(key, rKey) {
				exists = true; break
			}
		}

		if exists == false {
			key := strings.Repeat(fKey + " ", 2)
			uniqMap[key] = q
		}
	}

	for _, q := range uniqMap {
		result = append(result, q)
	}

	return result
}
