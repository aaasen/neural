package neural

import (
	"log"
	"math"
)

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	return sigmoid(x) * (1 - sigmoid(x))
}

func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Fatalf("dot product only works on vectors of equal length\n")
		return 0
	}

	products := make([]float64, len(a))

	for i := 0; i < len(a); i++ {
		products[i] = a[i] * b[i]
	}

	return sum(products)
}

func sum(slice []float64) float64 {
	sum := float64(0)

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}

func normal(slice []float64) []float64 {
	max := max(slice)
	min := min(slice)

	for i, v := range slice {
		slice[i] = (v - min) / (max - min)
	}

	return slice
}

func max(slice []float64) float64 {
	max := math.SmallestNonzeroFloat64

	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func min(slice []float64) float64 {
	min := math.MaxFloat64

	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
}
