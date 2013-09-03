package neural

import ()

type Layer struct {
	neurons []*Neuron
}

func NewLayer(neurons []*Neuron) *Layer {
	return &Layer{
		neurons,
	}
}

func RandomLayer(numInputs, numOutputs int) *Layer {
	neurons := make([]*Neuron, numOutputs)

	for i := range neurons {
		neurons[i] = RandomNeuron(numInputs)
	}

	return NewLayer(neurons)
}

func (layer *Layer) Score(inputs []float64) []float64 {
	scores := make([]float64, len(layer.neurons))

	for i := range layer.neurons {
		scores[i] = layer.neurons[i].Score(inputs)
	}

	return normal(scores)
}
