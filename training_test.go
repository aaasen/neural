package neural

import (
	"testing"
)

func TestTrainXOR(t *testing.T) {
	trainingSet := NewTrainingSet([]*Example{
		NewExample([]float64{0, 0}, []float64{0}),
		NewExample([]float64{0, 1}, []float64{1}),
		NewExample([]float64{1, 0}, []float64{1}),
		NewExample([]float64{1, 1}, []float64{0}),
	})

	net := RandomNet(2, 2, 1)
	net.Train(trainingSet)

	for _, example := range trainingSet.examples {
		netOutput := net.Score(example.input)

		if !equals(example.output, netOutput) {
			t.Errorf("input: %v\nexpected: %v\nreceived: %v\n",
				example.input,
				example.output,
				netOutput,
			)
		}
	}
}
