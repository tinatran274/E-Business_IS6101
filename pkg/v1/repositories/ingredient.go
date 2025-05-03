package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type IngredientRepository struct {
	db *db.Queries
}

func NewIngredientRepository(db *db.Queries) *IngredientRepository {
	return &IngredientRepository{
		db: db,
	}
}

func (r *IngredientRepository) GetIngredientByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Ingredient, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetIngredientByID").
			Observe(time.Since(t).Seconds())
	}()

	ingredient, err := r.db.GetIngredientById(ctx, id)
	if err != nil {
		return nil, err
	}

	category, err := r.db.GetCategoryById(ctx, ingredient.CategoryID)
	if err != nil {
		return nil, err
	}

	ingredientModel := models.ToIngredient(ingredient)
	ingredientModel.Category = models.ToCategory(category)
	return ingredientModel, nil
}

func (r *IngredientRepository) GetIngredients(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.Ingredient, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetIngredients").
			Observe(time.Since(t).Seconds())
	}()

	arg := db.GetIngredientsParams{
		Limit:   filter.Limit,
		Offset:  filter.Offset,
		Keyword: filter.Keyword,
		SortBy:  filter.SortBy,
		OrderBy: filter.OrderBy,
	}
	ingredients, err := r.db.GetIngredients(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Ingredient, len(ingredients))
	for i, ingredient := range ingredients {
		result[i] = models.ToIngredient(ingredient)
		category, err := r.db.GetCategoryById(ctx, ingredient.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i].Category = models.ToCategory(category)
	}

	return result, nil
}

func (r *IngredientRepository) CountIngredients(
	ctx context.Context,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountIngredients").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountIngredients(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *IngredientRepository) GetIngredientByDishId(
	ctx context.Context,
	dishID uuid.UUID,
) ([]*models.Ingredient, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetIngredientByDishId").
			Observe(time.Since(t).Seconds())
	}()

	ingredients, err := r.db.GetIngredientByDishId(ctx, dishID)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Ingredient, len(ingredients))
	for i, ingredient := range ingredients {
		result[i] = &models.Ingredient{
			ID:          ingredient.ID,
			Name:        ingredient.Name,
			Description: ingredient.Description,
			Removal:     ingredient.Removal,
			Kcal:        ingredient.Kcal,
			Protein:     ingredient.Protein,
			Lipits:      ingredient.Lipits,
			Glucids:     ingredient.Glucids,
			Canxi:       ingredient.Canxi,
			Phosphor:    ingredient.Phosphor,
			Fe:          ingredient.Fe,
			VitaminA:    ingredient.VitaminA,
			VitaminB1:   ingredient.VitaminB1,
			VitaminB2:   ingredient.VitaminB2,
			VitaminC:    ingredient.VitaminC,
			VitaminPp:   ingredient.VitaminPp,
			BetaCaroten: ingredient.BetaCaroten,
		}
		category, err := r.db.GetCategoryById(ctx, ingredient.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i].Category = models.ToCategory(category)
	}

	return result, nil
}
