package neural

func make2DSliceFloat64(w, h int) [][]float64 {
	slice := make([][]float64, w)

	for i := range slice {
		slice[i] = make([]float64, h)
	}

	return slice
}
