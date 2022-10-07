package formulas

import "testing"

func TestRepeatBrzycki(t *testing.T) {
	testRepeat(t, RepeatBrzycki, &testCase{
		m: 100,
		w: 60,
		r: 15,
	})
}

func TestMaxWeightBrzycki(t *testing.T) {
	testMaxWeight(t, MaxWeightBrzycki, &testCase{
		m:     100,
		w:     60,
		r:     20,
		delta: 30,
	})
}

func TestWorkWeightBrzycki(t *testing.T) {
	testWorkWeight(t, WorkWeightBrzycki, &testCase{
		m:     100,
		w:     100,
		r:     1,
		delta: 4,
	})
}
