package neural

import (
	// "fmt"
	"math"
)

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
		if i == len(layers)-1 {
			layers[i] = RandomLayer(numInputs, numOutputs)
		} else {
			layers[i] = RandomLayer(numInputs, numInputs)
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

func (net *Net) Train(trainingSet *TrainingSet) {
	/*
		pseudocode from Wikipedia (https://en.wikipedia.org/wiki/Backpropagation#Algorithm)

		 initialize the weights in the network (often small random values)
		    do
		       for each example e in the training set
		          O = neural-net-output(network, e)  // forward pass
		          T = teacher output for e
		          compute error (T - O) at the output units
		          compute delta_wh for all weights from hidden layer to output layer  // backward pass
		          compute delta_wi for all weights from input layer to hidden layer   // backward pass continued
		          update the weights in the network
		    until all examples classified correctly or stopping criterion satisfied
		    return the network
	*/

	rate := 0.04

	for _, example := range trainingSet.examples {
		output := net.Score(example.input)
		err := calcError(example.output, output)
		weightDelta := err * rate

		for _, neuron := range net.layers[len(net.layers)-1].neurons {
			for j, _ := range neuron.weights {
				neuron.weights[j] += weightDelta
			}
		}

		// fmt.Println(weightDelta)

	}
}

func (net *Net) backPropogate(deltas []float64) []float64 {
	// 	if len(net.layers) == 1 {

	// 	} else if len(net.layers) > 1 {

	// 	} else {
	// 		return []float64{0}
	// 	}

	return nil
}

func calcError(expected, result []float64) float64 {
	errorSum := 0.0

	for i := range expected {
		errorSum += math.Pow(expected[i]-result[i], 2)
	}

	return errorSum
}
