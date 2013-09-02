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

func (layer *Layer) Score(inputs []float64) ([]float64, error) {
	scores := make([]float64, len(inputs))

	for i := 0; i < len(inputs); i++ {
		score, err := layer.neurons[i].Score(inputs)

		if err != nil {
			return nil, err
		}

		scores[i] = score
	}

	return scores, nil
}
