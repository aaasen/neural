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
	scores := make([]float64, len(inputs))

	for i := 0; i < len(inputs); i++ {
		scores[i] = layer.neurons[i].Score(inputs)
	}

	return scores
}
