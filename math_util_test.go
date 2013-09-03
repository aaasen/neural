package neural

import (
	"testing"
)

func TestDotProduct(t *testing.T) {
	expected := float64(2)
	result := dotProduct([]float64{1, 2, 3}, []float64{-1, 0, 1})

	if expected != result {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestSum(t *testing.T) {
	expected := float64(6)
	result := sum([]float64{1, 2, 3})

	if expected != result {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestNormal(t *testing.T) {
	expected := []float64{0, 0.5, 1}
	result := normal([]float64{1, 2, 3})

	if !equals(expected, result) {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestMax(t *testing.T) {
	max := max([]float64{-1, 0, 1})

	if max != 1 {
		t.Errorf("expected %v, got %v\n", 1, max)
	}
}

func TestMin(t *testing.T) {
	min := min([]float64{-1, 0, 1})

	if min != -1 {
		t.Errorf("expected %v, got %v\n", 1, min)
	}
}
