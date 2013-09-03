package neural

import (
	"testing"
)

var zeroNeuron = NewNeuron([]float64{0, 0, 0})
var identityNeuron = NewNeuron([]float64{1, 1, 1})
var aNeuron = NewNeuron([]float64{1, 2, 3})
var bNeuron = NewNeuron([]float64{-1, 0, 1})
var cNeuron = NewNeuron([]float64{3, 4})
var dNeuron = NewNeuron([]float64{-1, 1})
var eNeuron = NewNeuron([]float64{1, 2})

var aLayer = NewLayer([]*Neuron{
	zeroNeuron,
	identityNeuron,
	aNeuron,
})

var bLayer = NewLayer([]*Neuron{
	aNeuron,
	bNeuron,
})

var cLayer = NewLayer([]*Neuron{
	cNeuron,
	dNeuron,
	eNeuron,
})

var aNet = NewNet([]*Layer{
	aLayer,
	bLayer,
})

var bNet = NewNet([]*Layer{
	cLayer,
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

func TestBackNet(t *testing.T) {
	expected := []float64{26, 74}
	result := bNet.Back([]float64{1, 2})[0]

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

func TestBackLayer(t *testing.T) {
	expected := []float64{-1, 2, 5}
	result := bLayer.Back([]float64{1, 2})

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

func TestBackNeuron(t *testing.T) {
	expected := []float64{2, 4, 6}
	result := aNeuron.Back(2)

	if !equals(expected, result) {
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
