package smartsifter

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type GMM struct {
}

func (gmm GMM) probabilityDensity(x, rates []float64, avgs []*mat.VecDense, vcvs []*mat.Dense) (float64, []float64) {
	p := 0.0
	ps := make([]float64, len(rates))
	for i := 0; i < len(rates); i++ {
		p_ := gmm.probabilityDensityAt(x, rates[i], avgs[i], vcvs[i])
		ps[i] = p_
		p += p_
	}
	return p, ps
}

func (gmm GMM) probabilityDensityAt(x []float64, rate float64, avg *mat.VecDense, vcv *mat.Dense) float64 {
	return rate * gmm.distribution(mat.NewVecDense(len(x), x), avg, vcv)
}

func (gmm GMM) distribution(x, avg *mat.VecDense, vcv *mat.Dense) float64 {
	dim, _ := x.Dims()
	c := 1 / (math.Pow(math.Sqrt(2*math.Pi), float64(dim)) * math.Sqrt(mat.Det(vcv)))

	var sub mat.VecDense
	sub.SubVec(x, avg)

	var ivcv mat.Dense
	ivcv.Inverse(vcv)

	return c * math.Exp(-0.5*mat.Inner(&sub, &ivcv, &sub))
}
