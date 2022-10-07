package config

import "fmt"

type Config struct {
	Name         string        `json:"name"`
	Maximums     *Maximum      `json:"maximum"`
	Modification *Modification `json:"modification"`
	Round        float64       `json:"round"`
	MicroCycles  []*MicroCycle `json:"micro_cycle"`
}

type Maximum struct {
	PressWeight float64 `json:"press_weight"`
	PressRepeat int     `json:"press_repeat"`
	LiftWeight  float64 `json:"lift_weight"`
	LiftRepeat  int     `json:"lift_repeat"`
	SquatWeight float64 `json:"squat_weight"`
	SquatRepeat int     `json:"squat_repeat"`
}

func (m *Maximum) String() string {
	return fmt.Sprintf(
		"Max Жим:%v Повторы:%v\nMax Становая:%v Повторы:%v\nMax Присед:%v Повторы:%v\n",
		m.PressWeight, m.PressRepeat, m.LiftWeight, m.LiftRepeat, m.SquatWeight, m.SquatRepeat,
	)
}

type Modification struct {
	AddPressRepeat int `json:"add_press_repeat"`
	AddLiftRepeat  int `json:"add_lift_repeat"`
	AddSquatRepeat int `json:"add_squat_repeat"`
	MaxPercent     int `json:"max_percent"`
}

type MicroCycle struct {
	Name      string        `json:"name"`
	Warm      TrainingUnits `json:"warm"`
	Trainings []*Exercise   `json:"trainings"`
}

type TrainingUnits []TrainingUnit

type TrainingUnit [3]int

func (t TrainingUnit) Percent() int     { return t[0] }
func (t TrainingUnit) Repetitions() int { return t[1] }
func (t TrainingUnit) Count() int       { return t[2] }
