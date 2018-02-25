package smartsifter

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestGMM(t *testing.T) {
	gmm := GMM{}

	x := mat.NewVecDense(1, []float64{0.0})
	avg := mat.NewVecDense(1, []float64{0.0})
	vcv := mat.NewDense(1, 1, []float64{1.0})

	d := gmm.distribution(x, avg, vcv)
	if fmt.Sprintf("%f", d) != "0.398942" {
		t.Errorf("Standard normal distribution should return 0.398942 when passed 0, but %f", d)

	}
}
