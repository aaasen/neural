package main

import (
	"fmt"
	"github.com/aaasen/neural"
)

func main() {
	neuron, err := neural.NewNeuron([]float64{1, 2})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(neuron.Score([]float64{1, 1}))
}
