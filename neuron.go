package neural

import (
	"log"
	"math/rand"
)

type Neuron struct {
	weights []float64
}

func NewNeuron(weights []float64) *Neuron {
	return &Neuron{
		weights,
	}
}

func RandomNeuron(numInputs int) *Neuron {
	weights := make([]float64, numInputs)

	for i := range weights {
		weights[i] = rand.Float64()
	}

	return NewNeuron(weights)
}

func (neuron *Neuron) Score(inputs []float64) float64 {
	if len(neuron.weights) != len(inputs) {
		log.Fatalf("wrong number of inputs for neuron: expected %v, got %v\n",
			len(neuron.weights), len(inputs))
		return 0
	}

	return dotProduct(inputs, neuron.weights)
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
