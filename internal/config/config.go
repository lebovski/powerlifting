package config

import "fmt"

type Config struct {
	Name         string        `json:"name"`
	Maximums     Maximum       `json:"maximum"`
	Modification Modification  `json:"modification"`
	Round        float64       `json:"round"`
	MicroCycles  []*MicroCycle `json:"micro_cycle"`
}

type Max struct {
	Weight float64 `json:"weight"`
	Repeat int     `json:"repeat"`
}

type Maximum map[string]*Max

func (m Maximum) String() string {
	res := "Максимум\n"

	for k, v := range m {
		res += fmt.Sprintf("%v: %v/%v\n", k, v.Weight, v.Repeat)
	}

	return res
}

type Mod struct {
	MaxPercent int `json:"max_percent"`
	AddRepeat  int `json:"add_repeat"`
}

type Modification map[string]*Mod

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
