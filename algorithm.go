package smartsifter

func newContinuousAlgorithm(discount, alpha float64, mixtureNum, dim int) ContinuousAlgorithm {
	if mixtureNum == 0 {
		return NoOpContinuousAlgorithm{}
	}
	return newSDEM(discount, alpha, mixtureNum, dim)
}

type ContinuousAlgorithm interface {
	input([]float64, bool) float64
}

type NoOpContinuousAlgorithm struct {
}

func (c NoOpContinuousAlgorithm) input([]float64, bool) float64 {
	return 1.0
}

func newCategoricalAlgorithm(discount, beta float64, cellNum int) CategoricalAlgorithm {
	if cellNum == 0 {
		return NoOpCategoricalAlgorithm{}
	}
	return newSDLE(discount, beta, cellNum)
}

type CategoricalAlgorithm interface {
	input([]int, bool) float64
}

type NoOpCategoricalAlgorithm struct {
}

func (c NoOpCategoricalAlgorithm) input([]int, bool) float64 {
	return 1.0
}
