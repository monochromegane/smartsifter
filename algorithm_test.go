package smartsifter

import "testing"

func TestNewCategoricalAlgorithm(t *testing.T) {
	cellNum := 0
	algo := newCategoricalAlgorithm(1.0, 1.0, cellNum)
	if _, ok := algo.(NoOpCategoricalAlgorithm); !ok {
		t.Errorf("newCategoricalAlgorithm should return NoOpCategoricalAlgorithm when cellNum is zero.")
	}

	cellNum = 1
	algo = newCategoricalAlgorithm(1.0, 1.0, cellNum)
	if _, ok := algo.(*SDLE); !ok {
		t.Errorf("newCategoricalAlgorithm should return SDLE algorism when cellNum > 0.")
	}
}

func TestNewContinuousAlgorithm(t *testing.T) {
	mixtureNum := 0
	algo := newContinuousAlgorithm(1.0, 1.0, mixtureNum, 2)
	if _, ok := algo.(NoOpContinuousAlgorithm); !ok {
		t.Errorf("newContinuousAlgorithm should return NoOpContinuousAlgorithm when mixtureNum is zero.")
	}

	mixtureNum = 1
	algo = newContinuousAlgorithm(1.0, 1.0, mixtureNum, 2)
	if _, ok := algo.(*SDEM); !ok {
		t.Errorf("newContinuousAlgorithm should return SDEM algorism when mixtureNum > 0.")
	}
}
