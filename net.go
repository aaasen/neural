package neural

import ()

type Net struct {
	layers []*Layer
}

func NewNet(layers []*Layer) *Net {
	return &Net{
		layers,
	}
}

func RandomNet(numInputs, numLayers, numOutputs int) *Net {
	layers := make([]*Layer, numLayers)

	for i := range layers {
		if i == 0 {
			layers[i] = RandomLayer(numInputs, numOutputs)
		} else {
			layers[i] = RandomLayer(numOutputs, numOutputs)
		}
	}

	return NewNet(layers)
}

func (net *Net) Score(inputs []float64) []float64 {
	if len(net.layers) == 1 {
		return net.layers[0].Score(inputs)
	} else if len(net.layers) > 1 {
		newNet := NewNet(net.layers[1:])

		return newNet.Score(net.layers[0].Score(inputs))
	} else {
		return []float64{0}
	}
}
