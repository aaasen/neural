package neural

import (
	"testing"
)

var zeroNeuron = NewNeuron([]float64{0, 0, 0})
var identityNeuron = NewNeuron([]float64{1, 1, 1})
var aNeuron = NewNeuron([]float64{1, 2, 3})
var bNeuron = NewNeuron([]float64{-1, 0, 1})

var aLayer = NewLayer([]*Neuron{
	zeroNeuron,
	identityNeuron,
	aNeuron,
})

var bLayer = NewLayer([]*Neuron{
	aNeuron,
	bNeuron,
})

var aNet = NewNet([]*Layer{
	aLayer,
	bLayer,
})

func TestNet(t *testing.T) {
	expected := []float64{54, 14}
	result := aNet.Score([]float64{1, 2, 3})

	if !equals(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestLayer(t *testing.T) {
	expected := []float64{0, 6, 14}
	result := aLayer.Score([]float64{1, 2, 3})

	if !equals(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNeuron(t *testing.T) {
	expected := float64(2)
	result := aNeuron.Score([]float64{-1, 0, 1})

	if expected != result {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

func equals(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
