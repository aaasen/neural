package neural

import ()

type TrainingSet struct {
	examples []Example
}

func NewTrainingSet(examples []Example) *TrainingSet {
	return &TrainingSet{examples}
}
