package smartsifter

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func uniformVector(dim int) *mat.VecDense {
	vec := make([]float64, dim)
	for i, _ := range vec {
		vec[i] = rand.Float64()
	}
	return mat.NewVecDense(dim, vec)
}

func identifyMatrix(dim int) *mat.Dense {
	matrix := mat.NewDense(dim, dim, nil)
	j := 0
	for i := 0; i < dim; i++ {
		matrix.Set(i, j, 1.0)
		j += 1
	}
	return matrix
}
