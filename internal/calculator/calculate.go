package calculator

import (
	"errors"
	"fmt"

	"powerlifting/internal/config"
)

type Calculator struct {
	cfg *config.Config
	max *maximum
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

			for _, et := range trs.ExerciseTypes {
				tr, err := c.getTraining(append(mc.Warm, trs.Training...), et)
				if err != nil {
					return nil, err
				}

				res = append(res, fmt.Sprintf("%s: %s", et, tr))
			}

			res = append(res, "")
		}

		res = append(res, "")
	}

	return res, nil
}

func (c *Calculator) getTraining(tr []config.TrainingUnit, exT config.ExerciseType) (training, error) {
	res := make(training, 0, len(tr))

	maxWeight := float64(0)
	repeatAdder := 0

	switch exT {
	case config.Press, config.PressMax:
		maxWeight = c.max.pressWeight
		repeatAdder = c.cfg.Modification.AddPressRepeat
	case config.Lift, config.LiftMax:
		maxWeight = c.max.liftWeight
		repeatAdder = c.cfg.Modification.AddLiftRepeat
	case config.Squat, config.SquatMax:
		maxWeight = c.max.squatWeight
		repeatAdder = c.cfg.Modification.AddSquatRepeat
	default:
		return nil, errors.New("unexpected exercise type")
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
