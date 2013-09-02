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
