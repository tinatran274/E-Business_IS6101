package usecases

import (
	"context"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type StatisticUseCase interface {
	CreateStatistic(
		ctx context.Context,
		statistic *models.Statistic,
	) (*models.Statistic, error)

	GetStatisticByUserIdAndDate(
		ctx context.Context,
		userID uuid.UUID,
		updatedAt time.Time,
	) (*models.Statistic, error)

	UpdateStatisticByUserIdAndDate(
		ctx context.Context,
		statistic *models.Statistic,
	) (*models.Statistic, error)
}
