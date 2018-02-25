package smartsifter

import "math"

type SDLE struct {
	t             int
	r             float64
	beta          float64
	cellNum       int
	occurrences   []float64
	probabilities []float64
	categorizer   Categorizer
}

func newSDLE(discount, beta float64, cellNum int) *SDLE {
	sdle := &SDLE{
		r:           discount,
		beta:        beta,
		cellNum:     cellNum,
		categorizer: noOpCategorizer{},
	}
	sdle.initialize()
	return sdle
}

func (sdle *SDLE) initialize() {
	sdle.t = 1
	sdle.occurrences = make([]float64, sdle.cellNum)
	sdle.probabilities = make([]float64, sdle.cellNum)
}

func (sdle *SDLE) input(x []int, update bool) float64 {
	idx := sdle.categorizer.Index(x)
	q := sdle.probabilities[idx]
	if update {
		sdle.update(idx)
	}
	return q / float64(sdle.categorizer.Size(idx))
}

func (sdle *SDLE) update(idx int) {
	for i := 0; i < sdle.cellNum; i++ {
		delta := 0.0
		if i == idx {
			delta = 1.0
		}
		t := (1-sdle.r)*sdle.occurrences[i] + delta
		q := (t + sdle.beta) * sdle.r / ((1 - math.Pow((1-sdle.r), float64(sdle.t))) + sdle.r*float64(sdle.cellNum)*sdle.beta)

		sdle.occurrences[i] = t
		sdle.probabilities[i] = q
	}
	sdle.t += 1
}
