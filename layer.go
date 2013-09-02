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

func (layer *Layer) Score(inputs []float64) []float64 {
	scores := make([]float64, len(layer.neurons))

	for i := range layer.neurons {
		scores[i] = layer.neurons[i].Score(inputs)
	}

	return scores
}
