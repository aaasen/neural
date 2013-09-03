package neural

import (
	"testing"
)

func TestRandomNeuron(t *testing.T) {
	neuron := RandomNeuron(3)

	neuron.Score([]float64{1, 2, 3})
}

func TestRandomLayer(t *testing.T) {
	layer := RandomLayer(3, 4)
	scores := layer.Score([]float64{1, 2, 3})

	if len(scores) != 4 {
		t.Errorf("expected 4 outputs but received %v\n", len(scores))
	}
}

func TestRandomNet(t *testing.T) {
	net := RandomNet(3, 4, 4)
	scores := net.Score([]float64{1, 2, 3})

	if len(scores) != 4 {
		t.Errorf("expected 4 outputs but received %v\n", len(scores))
	}
}
