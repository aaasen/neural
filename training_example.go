package neural

import ()

type Example struct {
	input, output []float64
}

func NewExample(input, output []float64) *Example {
	return &Example{input, output}
}
