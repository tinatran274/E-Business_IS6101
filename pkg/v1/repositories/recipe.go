package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type RecipeRepository struct {
	db *db.Queries
}

func NewRecipeRepository(db *db.Queries) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (r *RecipeRepository) CreateRecipe(
	ctx context.Context,
	recipe *models.Recipe,
) (*models.Recipe, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateRecipe").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateRecipe(ctx, db.CreateRecipeParams{
		DishID:       recipe.DishID,
		IngredientID: recipe.IngredientID,
		Unit:         recipe.Unit,
	})
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r *RecipeRepository) DeleteRecipe(
	ctx context.Context,
	dishId, ingredientId uuid.UUID,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteRecipe").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteRecipe(ctx, db.DeleteRecipeParams{
		DishID:       dishId,
		IngredientID: ingredientId,
	})
	if err != nil {
		return err
	}

	return nil
}
