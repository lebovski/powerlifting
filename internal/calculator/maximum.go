package calculator

import (
	"powerlifting/internal/config"
	"powerlifting/internal/formulas"
)

type maximum struct {
	pressWeight float64
	liftWeight  float64
	squatWeight float64
}

func calculateMax(cfg *config.Config) *maximum {
	var (
		maxCoeff = float64(cfg.Modification.MaxPercent) / 100

		pressWeight = formulas.MaxWeight(cfg.Maximums.PressWeight, uint(cfg.Maximums.PressRepeat))
		liftWeight  = formulas.MaxWeight(cfg.Maximums.LiftWeight, uint(cfg.Maximums.LiftRepeat))
		squatWeight = formulas.MaxWeight(cfg.Maximums.SquatWeight, uint(cfg.Maximums.SquatRepeat))

		pressRepeat = uint(1 + cfg.Modification.AddPressRepeat)
		liftRepeat  = uint(1 + cfg.Modification.AddLiftRepeat)
		squatRepeat = uint(1 + cfg.Modification.AddSquatRepeat)
	)

	res := &maximum{
		pressWeight: formulas.WorkWeight(pressWeight*maxCoeff, pressRepeat),
		liftWeight:  formulas.WorkWeight(liftWeight*maxCoeff, liftRepeat),
		squatWeight: formulas.WorkWeight(squatWeight*maxCoeff, squatRepeat),
	}

	return res
}
