package formulas

import "testing"

func TestRepeat(t *testing.T) {
	testRepeat(t, Repeat, &testCase{
		m: 100,
		w: 60,
		r: 15,
	})
}

func TestMaxWeight(t *testing.T) {
	cases := []*testCase{
		{
			m:     100,
			w:     60,
			r:     20,
			delta: 0.15,
		},
		{
			m:     100,
			w:     100,
			r:     1,
			delta: 0,
		},
	}

	for _, c := range cases {
		testMaxWeight(t, MaxWeight, c)
	}
}

func TestWorkWeight(t *testing.T) {
	testWorkWeight(t, WorkWeight, &testCase{
		m:     100,
		w:     100,
		r:     1,
		delta: 4,
	})
}
