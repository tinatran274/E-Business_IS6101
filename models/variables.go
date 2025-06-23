package models

var (
	PendingStatus  string = "pending"
	RejectedStatus string = "rejected"
	ActiveStatus   string = "active"
	InactiveStatus string = "inactive"
	DeletedStatus  string = "deleted"

	ValidStatus = map[string]bool{
		PendingStatus:  true,
		RejectedStatus: true,
		ActiveStatus:   true,
		InactiveStatus: true,
	}

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

	DefaultAge           int32   = 18
	DefaultHeight        int32   = 160
	DefaultWeight        int32   = 50
	DefaultGender        string  = "female"
	DefaultExerciseLevel string  = "no exercise"
	DefaultAim           string  = "maintain weight"
	DefaultShippingCost  float64 = 30000

	MinAge        int     = 18
	MaxAge        int     = 100
	MinHeight     int     = 140
	MaxHeight     int     = 220
	MinWeight     int     = 30
	MaxWeight     int     = 200
	LimitCalories float64 = 1000

	LimitTextLength     int = 40
	LimitLongTextLength int = 100
	LimitImageSize      int = 5 * 1024 * 1024 // 5MB
	LimitImageWidth     int = 2000
	LimitImageHeight    int = 2000

	ValidOrderBy = map[string]bool{
		"updated_at": true,
		"created_at": true,
		"name":       true,
		"calories":   true,
	}

	ValidSortBy = map[string]bool{
		"asc":  true,
		"desc": true,
	}

	SortByDefault  string = "asc"
	OrderByDefault string = "updated_at"
)
