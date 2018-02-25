package smartsifter

func newContinuousAlgorism(discount, alpha float64, mixtureNum, dim int) ContinuousAlgorism {
	return NoOpContinuousAlgorism{}
}

type ContinuousAlgorism interface {
	input([]float64, bool) float64
}

type NoOpContinuousAlgorism struct {
}

func (c NoOpContinuousAlgorism) input([]float64, bool) float64 {
	return 1.0
}

func newCategoricalAlgorism(discount, beta float64, cellNum int) CategoricalAlgorism {
	if cellNum == 0 {
		return NoOpCategoricalAlgorism{}
	}
	return newSDLE(discount, beta, cellNum)
}

type CategoricalAlgorism interface {
	input([]int, bool) float64
}

type NoOpCategoricalAlgorism struct {
}

func (c NoOpCategoricalAlgorism) input([]int, bool) float64 {
	return 1.0
}
