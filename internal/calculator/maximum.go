package calculator

import (
	"powerlifting/internal/config"
	"powerlifting/internal/formulas"
)

type maximum map[string]float64

func calculateMax(cfg *config.Config) maximum {
	res := make(maximum, len(cfg.Maximums))

	for ex := range cfg.Maximums {
		maxCoeff := float64(1)
		repeat := uint(1)

		if c, ok := cfg.Modification[ex]; ok {
			maxCoeff = float64(c.MaxPercent) / 100
			repeat = uint(1 + c.AddRepeat)
		}

		weight := formulas.MaxWeight(cfg.Maximums[ex].Weight, uint(cfg.Maximums[ex].Repeat))

		res[ex] = formulas.WorkWeight(weight*maxCoeff, repeat)
	}

	return res
}
