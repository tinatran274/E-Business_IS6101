package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type StatisticRepository struct {
	db *db.Queries
}

func NewStatisticRepository(db *db.Queries) *StatisticRepository {
	return &StatisticRepository{
		db: db,
	}
}

func (r *StatisticRepository) CreateStatistic(
	ctx context.Context,
	statistic *models.Statistic,
) (*models.Statistic, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateStatistic").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateStatistic(ctx, db.CreateStatisticParams{
		UpdatedAt:        statistic.UpdatedAt,
		UserID:           statistic.UserID,
		MorningCalories:  statistic.MorningCalories,
		LunchCalories:    statistic.LunchCalories,
		DinnerCalories:   statistic.DinnerCalories,
		SnackCalories:    statistic.SnackCalories,
		ExerciseCalories: statistic.ExerciseCalories,
	})
	if err != nil {
		return nil, err
	}

	return statistic, nil
}

func (r *StatisticRepository) GetStatisticByUserIdAndDate(
	ctx context.Context,
	userID uuid.UUID,
	updatedAt time.Time,
) (*models.Statistic, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetStatisticByUserIdAndDate").
			Observe(time.Since(t).Seconds())
	}()

	statistic, err := r.db.GetStatisticByUserIdAndDate(ctx, db.GetStatisticByUserIdAndDateParams{
		UserID:    userID,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return nil, err
	}

	return models.ToStatistic(&statistic), nil
}

func (r *StatisticRepository) GetStatisticByUserIdAndDateRange(
	ctx context.Context,
	userID uuid.UUID,
	startDate time.Time,
	endDate time.Time,
) ([]*models.Statistic, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetStatisticByUserIdAndDateRange").
			Observe(time.Since(t).Seconds())
	}()

	statistic, err := r.db.GetStatisticByUserIdAndDateRange(
		ctx,
		db.GetStatisticByUserIdAndDateRangeParams{
			UserID:    userID,
			StartDate: startDate,
			EndDate:   endDate,
		},
	)
	if err != nil {
		return nil, err
	}

	statisticModels := make([]*models.Statistic, len(statistic))
	for i, s := range statistic {
		statisticModels[i] = models.ToStatistic(&s)
	}

	return statisticModels, nil
}

func (r *StatisticRepository) UpdateStatisticByUserIdAndDate(
	ctx context.Context,
	statistic *models.Statistic,
) (*models.Statistic, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateStatisticByUserIdAndDate").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateStatisticByUserIdAndDate(ctx, db.UpdateStatisticByUserIdAndDateParams{
		UserID:           statistic.UserID,
		UpdatedAt:        statistic.UpdatedAt,
		MorningCalories:  statistic.MorningCalories,
		LunchCalories:    statistic.LunchCalories,
		DinnerCalories:   statistic.DinnerCalories,
		SnackCalories:    statistic.SnackCalories,
		ExerciseCalories: statistic.ExerciseCalories,
	})
	if err != nil {
		return nil, err
	}

	return statistic, nil
}
