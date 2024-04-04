package config

type Exercise struct {
	Name          string        `json:"name"`
	ExerciseTypes []string      `json:"exercise"`
	Training      TrainingUnits `json:"training"`
	UseWarm       bool          `json:"warm"`
}
