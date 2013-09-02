package main

import (
	"fmt"
	"github.com/aaasen/neural"
)

func main() {
	zero, _ := neural.NewNeuron([]float64{0, 0, 0})
	identity, _ := neural.NewNeuron([]float64{1, 1, 1})
	neuron, _ := neural.NewNeuron([]float64{1, 2, 3})

	layer := neural.NewLayer([]*neural.Neuron{
		zero,
		identity,
		neuron,
	})

	fmt.Println(layer.Score([]float64{1, 2, 3}))
}
