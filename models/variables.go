package models

var (
	ActiveStatus   string = "active"
	InactiveStatus string = "inactive"
	DeletedStatus  string = "deleted"

	GenderMap = map[string]bool{
		"male":   true,
		"female": true,
	}

	ExerciseMap = map[string]bool{
		"no exercise":                       true,
		"light exercise (1-3 days/week)":    true,
		"moderate exercise (3-5 days/week)": true,
		"intense exercise (6-7 days/week)":  true,
	}

	AimMap = map[string]bool{
		"lose weight":     true,
		"maintain weight": true,
		"gain weight":     true,
	}

	DefaultAge           int32  = 18
	DefaultHeight        int32  = 160
	DefaultWeight        int32  = 50
	DefaultGender        string = "female"
	DefaultExerciseLevel string = "no exercise"
	DefaultAim           string = "maintain weight"
)
