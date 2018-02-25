package smartsifter

import "math"

type SmartSifter struct {
	categoricalAlgorism CategoricalAlgorism
	continuousAlgorism  ContinuousAlgorism
}

func NewSmartSifter(discount, alpha, beta float64, cellNum, mixtureNum, dim int) *SmartSifter {
	return &SmartSifter{
		categoricalAlgorism: newCategoricalAlgorism(discount, beta, cellNum),
		continuousAlgorism:  newContinuousAlgorism(discount, alpha, mixtureNum, dim),
	}
}

func (ss *SmartSifter) Input(X []int, Y []float64, update bool) float64 {
	return ss.logLoss(
		ss.categoricalAlgorism.input(X, update),
		ss.continuousAlgorism.input(Y, update),
	)
}

func (ss SmartSifter) logLoss(p, q float64) float64 {
	return -math.Log(p * q)
}
