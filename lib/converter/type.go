package converter

// Преобразовывает тип bool в тип float64
func BoolToFloat64(value bool) float64 {
	if value == true {
		return 1
	}

	return 0
}
