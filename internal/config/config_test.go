package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	a := &Config{
		MicroCycles: []*MicroCycle{
			{
				Name: "cycle_1",
				Warm: TrainingUnits{{20, 8, 1}, {32, 7, 1}, {44, 6, 1}, {56, 5, 1}},
				Trainings: []*Exercise{
					{
						Name:          "train 1",
						ExerciseTypes: []ExerciseType{Squat, Press},
						Training:      TrainingUnits{{68, 4, 1}, {76, 3, 1}, {80, 3, 2}},
					},
					{
						Name:          "train 2",
						ExerciseTypes: []ExerciseType{Lift, Press},
						Training:      TrainingUnits{{64, 5, 1}, {68, 3, 4}},
					},
					{
						Name:          "train 3",
						ExerciseTypes: []ExerciseType{Squat, Press},
						Training:      TrainingUnits{{68, 4, 1}, {80, 3, 1}, {84, 2, 1}, {88, 1, 3}},
					},
					{
						Name:          "train 4",
						ExerciseTypes: []ExerciseType{Lift, Press},
						Training:      TrainingUnits{{68, 4, 1}, {76, 3, 1}, {80, 2, 3}},
					},
					{
						Name:          "train 5",
						ExerciseTypes: []ExerciseType{Squat, Press},
						Training:      TrainingUnits{{64, 5, 1}, {68, 4, 3}},
					},
					{
						Name:          "train 6",
						ExerciseTypes: []ExerciseType{Lift, Press},
						Training:      TrainingUnits{{68, 4, 1}, {80, 3, 1}, {84, 2, 1}, {88, 1, 3}},
					},
				},
			},
		},
	}

	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
