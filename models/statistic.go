package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type StatisticRepository interface {
	CreateStatistic(
		ctx context.Context,
		statistic *Statistic,
	) (*Statistic, error)
	GetStatisticByUserIdAndDate(
		ctx context.Context,
		userID uuid.UUID,
		updatedAt time.Time,
	) (*Statistic, error)
	GetStatisticByUserIdAndDateRange(
		ctx context.Context,
		userID uuid.UUID,
		startDate time.Time,
		endDate time.Time,
	) ([]*Statistic, error)
	UpdateStatisticByUserIdAndDate(
		ctx context.Context,
		statistic *Statistic,
	) (*Statistic, error)
}

type Statistic struct {
	UpdatedAt        time.Time `json:"updated_at"`
	UserID           uuid.UUID `json:"user_id"`
	MorningCalories  float64   `json:"morning_calories"`
	LunchCalories    float64   `json:"lunch_calories"`
	DinnerCalories   float64   `json:"dinner_calories"`
	SnackCalories    float64   `json:"snack_calories"`
	ExerciseCalories float64   `json:"exercise_calories"`
}

func NewStatistic(
	updatedAt time.Time,
	userID uuid.UUID,
	morningCalories float64,
	lunchCalories float64,
	dinnerCalories float64,
	snackCalories float64,
	exerciseCalories float64,
) *Statistic {
	return &Statistic{
		UpdatedAt:        updatedAt,
		UserID:           userID,
		MorningCalories:  morningCalories,
		LunchCalories:    lunchCalories,
		DinnerCalories:   dinnerCalories,
		SnackCalories:    snackCalories,
		ExerciseCalories: exerciseCalories,
	}
}

func ToStatistic(statistic *db.Statistic) *Statistic {
	return &Statistic{
		UpdatedAt:        statistic.UpdatedAt,
		UserID:           statistic.UserID,
		MorningCalories:  statistic.MorningCalories,
		LunchCalories:    statistic.LunchCalories,
		DinnerCalories:   statistic.DinnerCalories,
		SnackCalories:    statistic.SnackCalories,
		ExerciseCalories: statistic.ExerciseCalories,
	}
}
