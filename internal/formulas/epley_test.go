package formulas

import "testing"

func TestRepeatEpley(t *testing.T) {
	testRepeat(t, RepeatEpley, &testCase{
		m: 100,
		w: 60,
		r: 20,
	})
}

func TestMaxWeightEpley(t *testing.T) {
	testMaxWeight(t, MaxWeightEpley, &testCase{
		m:     100,
		w:     60,
		r:     20,
		delta: 0.3,
	})
}

func TestWorkWeightEpley(t *testing.T) {
	testWorkWeight(t, WorkWeightEpley, &testCase{
		m:     100,
		w:     100,
		r:     1,
		delta: 4,
	})
}
