package formulas

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	m     float64 // Полный максимум.
	w     float64 // Вес штанги.
	r     uint    // Количество повторений.
	delta float64 // Погрешность.
}

func testRepeat(t *testing.T, f func(m, w float64) uint, c *testCase) {
	res := f(c.m, c.w)
	require.Equal(t, c.r, res, fmt.Sprintf("expected: %v, was: %v", c.r, res))
}

func testMaxWeight(t *testing.T, f func(w float64, r uint) float64, c *testCase) {
	res := f(c.w, c.r)
	require.InDelta(t, c.m, res, c.delta, fmt.Sprintf("expected: %v, was: %v", c.m, res))
}

func testWorkWeight(t *testing.T, f func(m float64, r uint) float64, c *testCase) {
	res := f(c.m, c.r)
	require.InDelta(t, c.w, res, c.delta, fmt.Sprintf("expected: %v, was: %v", c.w, res))
}
