package smartsifter

import "math"

type SmartSifter struct {
	categoricalAlgorithm CategoricalAlgorithm
	continuousAlgorithm  ContinuousAlgorithm
}

func NewSmartSifter(discount, alpha, beta float64, cellNum, mixtureNum, dim int) *SmartSifter {
	return &SmartSifter{
		categoricalAlgorithm: newCategoricalAlgorithm(discount, beta, cellNum),
		continuousAlgorithm:  newContinuousAlgorithm(discount, alpha, mixtureNum, dim),
	}
}

func (ss *SmartSifter) Input(X []int, Y []float64, update bool) float64 {
	return ss.logLoss(
		ss.categoricalAlgorithm.input(X, update),
		ss.continuousAlgorithm.input(Y, update),
	)
}

func (ss SmartSifter) logLoss(p, q float64) float64 {
	return -math.Log(p * q)
}
