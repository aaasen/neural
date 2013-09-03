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

	return scores
}

func (layer *Layer) Back(inputs []float64) []float64 {
	numWeights := len(layer.neurons[0].weights)
	scores := make2DSliceFloat64(numWeights, len(layer.neurons))

	for i, neuron := range layer.neurons {
		results := neuron.Back(inputs[i])

		for j, result := range results {
			scores[j][i] = result
		}
	}

	sumScores := make([]float64, numWeights)

	for i, scoreSet := range scores {
		sumScores[i] = sum(scoreSet)
	}

	return sumScores
}
