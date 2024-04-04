package calculator

import (
	"fmt"

	"powerlifting/internal/config"
)

type Calculator struct {
	cfg *config.Config
	max maximum
}

func (c *Calculator) Calculate() ([]string, error) {
	res := make([]string, 0)

	for _, mc := range c.cfg.MicroCycles {
		res = append(res, mc.Name)

		previousName := ""
		for _, trs := range mc.Trainings {
			if previousName != trs.Name {
				res = append(res, trs.Name)
			} else {
				res = res[:len(res)-1]
			}
			previousName = trs.Name

			for _, ex := range trs.ExerciseTypes {
				var tr training
				var err error

				if trs.UseWarm {
					tr, err = c.getTraining(append(mc.Warm, trs.Training...), ex)
				} else {
					tr, err = c.getTraining(trs.Training, ex)
				}

				if err != nil {
					return nil, err
				}

				res = append(res, fmt.Sprintf("%s: %s", ex, tr))
			}

			res = append(res, "")
		}

		res = append(res, "")
	}

	return res, nil
}

func (c *Calculator) getTraining(tr []config.TrainingUnit, ex string) (training, error) {
	res := make(training, 0, len(tr))

	maxWeight := c.max[ex]
	repeatAdder := 0

	mod, ok := c.cfg.Modification[ex]
	if ok {
		repeatAdder = mod.AddRepeat
	}

	for _, s := range tr {
		weight := maxWeight * float64(s.Percent()) / 100

		res = append(res, &unit{
			weight: round(weight, c.cfg.Round),
			repeat: s.Repetitions() + repeatAdder,
			count:  s.Count(),
		})
	}

	return res, nil
}

func NewCalculator(cfg *config.Config) *Calculator {
	return &Calculator{
		cfg: cfg,
		max: calculateMax(cfg),
	}
}
