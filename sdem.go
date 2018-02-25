package smartsifter

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type SDEM struct {
	r            float64
	alpha        float64
	dim          int
	mixtureNum   int
	mixtureRates []float64
	averages     []*mat.VecDense
	averages_    []*mat.VecDense
	matrices     []*mat.Dense
	matrices_    []*mat.Dense
	gmm          GMM
}

func newSDEM(discount, alpha float64, mixtureNum, dim int) *SDEM {
	sdem := &SDEM{
		r:          discount,
		alpha:      alpha,
		mixtureNum: mixtureNum,
		dim:        dim,
		gmm:        GMM{},
	}
	sdem.initialize()
	return sdem
}

func (sdem *SDEM) initialize() {
	rand.Seed(time.Now().UnixNano())

	sdem.mixtureRates = make([]float64, sdem.mixtureNum)
	sdem.averages = make([]*mat.VecDense, sdem.mixtureNum)
	sdem.averages_ = make([]*mat.VecDense, sdem.mixtureNum)
	sdem.matrices = make([]*mat.Dense, sdem.mixtureNum)
	sdem.matrices_ = make([]*mat.Dense, sdem.mixtureNum)
	for i := 0; i < sdem.mixtureNum; i++ {
		sdem.mixtureRates[i] = 1.0 / float64(sdem.mixtureNum)

		sdem.averages[i] = uniformVector(sdem.dim)
		var average_ mat.VecDense
		average_.ScaleVec(sdem.mixtureRates[i], sdem.averages[i])
		sdem.averages_[i] = &average_

		sdem.matrices[i] = identifyMatrix(sdem.dim)
		var matrix_ mat.Dense
		matrix_.Outer(1.0, sdem.averages[i], sdem.averages[i])
		matrix_.Add(&matrix_, sdem.matrices[i])
		matrix_.Scale(sdem.mixtureRates[i], &matrix_)
		sdem.matrices_[i] = &matrix_
	}
}

func (sdem *SDEM) input(x []float64, update bool) float64 {
	p, ps := sdem.gmm.probabilityDensity(x, sdem.mixtureRates, sdem.averages, sdem.matrices)
	if update {
		sdem.update(x, p, ps)
	}
	return p
}

func (sdem *SDEM) update(x []float64, p float64, ps []float64) {
	for k := 0; k < sdem.mixtureNum; k++ {
		// gamma
		gamma := (1.0-sdem.alpha*sdem.r)*(ps[k]/p) + (sdem.alpha*sdem.r)/float64(sdem.mixtureNum)

		// mixtureRate
		mixtureRate := (1.0-sdem.r)*sdem.mixtureRates[k] + sdem.r*gamma

		// mu_
		var mu_ mat.VecDense
		mu_.ScaleVec(1.0-sdem.r, sdem.averages_[k])
		var mu_2 mat.VecDense
		mu_2.ScaleVec(sdem.r*gamma, mat.NewVecDense(len(x), x))
		mu_.AddVec(&mu_, &mu_2)

		// mu
		var mu mat.VecDense
		mu.ScaleVec(1.0/mixtureRate, &mu_)

		// lambda_
		var lambda_ mat.Dense
		lambda_.Scale(1.0-sdem.r, sdem.matrices_[k])
		var lambda_2 mat.Dense
		lambda_2.Outer(sdem.r*gamma, mat.NewVecDense(len(x), x), mat.NewVecDense(len(x), x))
		lambda_.Add(&lambda_, &lambda_2)

		// lambda
		var lambda mat.Dense
		lambda.Scale(1.0/mixtureRate, &lambda_)
		var lambda2 mat.Dense
		lambda2.Outer(-1.0, &mu, &mu)
		lambda.Add(&lambda, &lambda2)

		// Update parameters
		sdem.mixtureRates[k] = mixtureRate
		sdem.averages_[k] = &mu_
		sdem.averages[k] = &mu
		sdem.matrices_[k] = &lambda_
		sdem.matrices[k] = &lambda
	}
}
