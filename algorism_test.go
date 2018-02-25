package smartsifter

import "testing"

func TestNewCategoricalAlgorism(t *testing.T) {
	cellNum := 0
	algo := newCategoricalAlgorism(1.0, 1.0, cellNum)
	if _, ok := algo.(NoOpCategoricalAlgorism); !ok {
		t.Errorf("newCategoricalAlgorism should return NoOpCategoricalAlgorism when cellNum is zero.")
	}

	cellNum = 1
	algo = newCategoricalAlgorism(1.0, 1.0, cellNum)
	if _, ok := algo.(*SDLE); !ok {
		t.Errorf("newCategoricalAlgorism should return SDLE algorism when cellNum > 0.")
	}
}

func TestNewContinuousAlgorism(t *testing.T) {
	mixtureNum := 0
	algo := newContinuousAlgorism(1.0, 1.0, mixtureNum, 2)
	if _, ok := algo.(NoOpContinuousAlgorism); !ok {
		t.Errorf("newContinuousAlgorism should return NoOpContinuousAlgorism when mixtureNum is zero.")
	}

	mixtureNum = 1
	algo = newContinuousAlgorism(1.0, 1.0, mixtureNum, 2)
	if _, ok := algo.(*SDEM); !ok {
		t.Errorf("newContinuousAlgorism should return SDEM algorism when mixtureNum > 0.")
	}
}
