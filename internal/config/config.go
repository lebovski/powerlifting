package config

import (
	"fmt"
	"strings"
)

// ExerciseType отвечает за тип упражнения
type ExerciseType int8

const (
	Press ExerciseType = iota + 1
	PressMax
	Lift
	LiftMax
	Squat
	SquatMax
)

func (i ExerciseType) String() string {
	switch i {
	case Press:
		return "Жим лёжа"
	case PressMax:
		return "Жим лёжа + max"
	case Lift:
		return "Становая тяга"
	case LiftMax:
		return "Становая тяга + max"
	case Squat:
		return "Присед"
	case SquatMax:
		return "Присед + max"
	default:
		panic(fmt.Sprintf("unexpected ExerciseType: %T", i))
	}
}

func (i ExerciseType) MarshalJSON() ([]byte, error) {
	switch i {
	case Press:
		return []byte("\"press\""), nil
	case PressMax:
		return []byte("\"press_max\""), nil
	case Lift:
		return []byte("\"lift\""), nil
	case LiftMax:
		return []byte("\"lift_max\""), nil
	case Squat:
		return []byte("\"squat\""), nil
	case SquatMax:
		return []byte("\"squat_max\""), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %v", i)
	}
}

func (i *ExerciseType) UnmarshalJSON(b []byte) (err error) {
	str := strings.ToLower(strings.Trim(string(b), "\""))
	switch str {
	case "press":
		*i = Press
	case "press_max":
		*i = PressMax
	case "lift":
		*i = Lift
	case "lift_max":
		*i = LiftMax
	case "squat":
		*i = Squat
	case "squat_max":
		*i = SquatMax
	default:
		err = fmt.Errorf("unknown exercise type: %v", i)
	}
	return
}

type Config struct {
	Round       float64       `json:"round"`
	MicroCycles []*MicroCycle `json:"micro_cycle"`
}

type MicroCycle struct {
	Name      string        `json:"name"`
	Warm      TrainingUnits `json:"warm"`
	Trainings []*Exercise   `json:"trainings"`
}

type Exercise struct {
	Name          string         `json:"name"`
	ExerciseTypes []ExerciseType `json:"exercise_types"`
	Training      TrainingUnits  `json:"training"`
}

type TrainingUnits []TrainingUnit

type TrainingUnit [3]int

func (t TrainingUnit) Percent() int     { return t[0] }
func (t TrainingUnit) Repetitions() int { return t[1] }
func (t TrainingUnit) Count() int       { return t[2] }
