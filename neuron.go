package neural

import (
	"errors"
)

type Neuron struct {
	weights []float64
}

func NewNeuron(weights []float64) (*Neuron, error) {
	return &Neuron{
		weights,
	}, nil
}

func (neuron *Neuron) Score(inputs []float64) (float64, error) {
	if len(neuron.weights) != len(inputs) {
		return 0, errors.New("wrong number of inputs for neuron")
	}

	return dotProduct(inputs, neuron.weights)
}

func dotProduct(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("dot product only works on vectors of equal length")
	}

	products := make([]float64, len(a))

	for i := 0; i < len(a); i++ {
		products[i] = a[i] * b[i]
	}

	return sum(products), nil
}

func sum(slice []float64) float64 {
	sum := float64(0)

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}
