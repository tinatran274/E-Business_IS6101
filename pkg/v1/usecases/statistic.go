package usecases

import (
	"context"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type StatisticUseCase struct {
	statisticRepo models.StatisticRepository
}

func NewStatisticUseCase(
	statisticRepo models.StatisticRepository,
) *StatisticUseCase {
	return &StatisticUseCase{
		statisticRepo: statisticRepo,
	}
}

func (s *StatisticUseCase) CreateStatistic(
	ctx context.Context,
	statistic *models.Statistic,
) (*models.Statistic, error) {
	createdStatistic, err := s.statisticRepo.CreateStatistic(ctx, statistic)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}
	return createdStatistic, nil
}

func (s *StatisticUseCase) GetStatisticByUserIdAndDate(
	ctx context.Context,
	userID uuid.UUID,
	updatedAt time.Time,
) (*models.Statistic, error) {
	statistic, err := s.statisticRepo.GetStatisticByUserIdAndDate(ctx, userID, updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			newStatistic := models.NewStatistic(
				updatedAt,
				userID,
				0,
				0,
				0,
				0,
				0,
			)
			return newStatistic, nil
		}

		return nil, response.NewInternalServerError(err)
	}

	return statistic, nil
}

func (s *StatisticUseCase) UpdateStatisticByUserIdAndDate(
	ctx context.Context,
	authInfo models.AuthenticationInfo,
	updatedAt time.Time,
	statistic *models.Statistic,
) (*models.Statistic, error) {
	existingStatistic, err := s.statisticRepo.GetStatisticByUserIdAndDate(
		ctx,
		authInfo.User.ID,
		updatedAt,
	)
	if err != nil && err != pgx.ErrNoRows {
		return nil, response.NewInternalServerError(err)
	}

	if existingStatistic == nil {
		_, err = s.statisticRepo.CreateStatistic(ctx, statistic)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}

		return statistic, nil
	}

	updatedStatistic, err := s.statisticRepo.UpdateStatisticByUserIdAndDate(
		ctx,
		statistic,
	)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return updatedStatistic, nil
}
