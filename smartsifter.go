package smartsifter

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
