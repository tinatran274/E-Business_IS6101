package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type DishRepository struct {
	db *db.Queries
}

func NewDishRepository(db *db.Queries) *DishRepository {
	return &DishRepository{
		db: db,
	}
}

func (r *DishRepository) GetDishByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Dish, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetDishByID").
			Observe(time.Since(t).Seconds())
	}()

	dish, err := r.db.GetDishById(ctx, id)
	if err != nil {
		return nil, err
	}

	category, err := r.db.GetCategoryById(ctx, dish.CategoryID)
	if err != nil {
		return nil, err
	}

	ingredients, err := r.db.GetIngredientByDishId(ctx, id)
	if err != nil {
		return nil, err
	}

	ingredientModels := make([]*models.IngredientWithUnit, len(ingredients))
	for i, ing := range ingredients {
		nutritionPer100g := models.Ingredient{
			ID:          ing.ID,
			Name:        ing.Name,
			Description: ing.Description,
			Removal:     ing.Removal,
			Kcal:        ing.Kcal,
			Protein:     ing.Protein,
			Lipits:      ing.Lipits,
			Glucids:     ing.Glucids,
			Canxi:       ing.Canxi,
			Phosphor:    ing.Phosphor,
			Fe:          ing.Fe,
			VitaminA:    ing.VitaminA,
			VitaminB1:   ing.VitaminB1,
			VitaminB2:   ing.VitaminB2,
			VitaminC:    ing.VitaminC,
			BetaCaroten: ing.BetaCaroten,
			VitaminPp:   ing.VitaminPp,
		}

		ingredientModels[i] = &models.IngredientWithUnit{
			ID:               ing.ID,
			Name:             ing.Name,
			Description:      ing.Description,
			Removal:          ing.Removal,
			Unit:             ing.Unit,
			NutritionPer100g: &nutritionPer100g,
		}
		category, err := r.db.GetCategoryById(ctx, ing.CategoryID)
		if err != nil {
			return nil, err
		}

		ingredientModels[i].Category = models.ToCategory(category)
	}

	dishModel := models.ToDish(dish)
	dishModel.Category = *models.ToCategory(category)
	dishModel.Ingredients = ingredientModels
	dishModel = dishModel.CalculateNutritionalValues()
	return dishModel, nil
}

func (r *DishRepository) GetDishes(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.Dish, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetDishs").
			Observe(time.Since(t).Seconds())
	}()

	arg := db.GetDishesParams{
		Limit:   filter.Limit,
		Offset:  filter.Offset,
		Keyword: filter.Keyword,
		SortBy:  filter.SortBy,
		OrderBy: filter.OrderBy,
	}
	dishes, err := r.db.GetDishes(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Dish, len(dishes))
	for i, dish := range dishes {
		result[i] = models.ToDish(dish)
		category, err := r.db.GetCategoryById(ctx, dish.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i].Category = *models.ToCategory(category)
		ingredients, err := r.db.GetIngredientByDishId(ctx, dish.ID)
		if err != nil {
			return nil, err
		}

		ingredientModels := make([]*models.IngredientWithUnit, len(ingredients))
		for i, ing := range ingredients {
			nutritionPer100g := models.Ingredient{
				ID:          ing.ID,
				Name:        ing.Name,
				Description: ing.Description,
				Removal:     ing.Removal,
				Kcal:        ing.Kcal,
				Protein:     ing.Protein,
				Lipits:      ing.Lipits,
				Glucids:     ing.Glucids,
				Canxi:       ing.Canxi,
				Phosphor:    ing.Phosphor,
				Fe:          ing.Fe,
				VitaminA:    ing.VitaminA,
				VitaminB1:   ing.VitaminB1,
				VitaminB2:   ing.VitaminB2,
				VitaminC:    ing.VitaminC,
				BetaCaroten: ing.BetaCaroten,
				VitaminPp:   ing.VitaminPp,
			}

			ingredientModels[i] = &models.IngredientWithUnit{
				ID:               ing.ID,
				Name:             ing.Name,
				Description:      ing.Description,
				Removal:          ing.Removal,
				Unit:             ing.Unit,
				NutritionPer100g: &nutritionPer100g,
			}
			category, err := r.db.GetCategoryById(ctx, ing.CategoryID)
			if err != nil {
				return nil, err
			}

			ingredientModels[i].Category = models.ToCategory(category)
		}

		result[i].Ingredients = ingredientModels
		result[i] = result[i].CalculateNutritionalValues()
	}

	return result, nil
}

func (r *DishRepository) CountDishes(
	ctx context.Context,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountDishs").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountDishes(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *DishRepository) GetDishesByIngredientID(
	ctx context.Context,
	id uuid.UUID,
	filter models.FilterParams,
) ([]*models.Dish, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetDishesByIngredientID").
			Observe(time.Since(t).Seconds())
	}()

	arg := db.GetDishesByIngredientIdParams{
		IngredientID: id,
		Limit:        filter.Limit,
		Offset:       filter.Offset,
		SortBy:       filter.SortBy,
		OrderBy:      filter.OrderBy,
	}
	dishes, err := r.db.GetDishesByIngredientId(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Dish, len(dishes))
	for i, dish := range dishes {
		result[i] = models.ToDish(dish)
		category, err := r.db.GetCategoryById(ctx, dish.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i].Category = *models.ToCategory(category)
		ingredients, err := r.db.GetIngredientByDishId(ctx, dish.ID)
		if err != nil {
			return nil, err
		}

		ingredientModels := make([]*models.IngredientWithUnit, len(ingredients))
		for i, ing := range ingredients {
			nutritionPer100g := models.Ingredient{
				ID:          ing.ID,
				Name:        ing.Name,
				Description: ing.Description,
				Removal:     ing.Removal,
				Kcal:        ing.Kcal,
				Protein:     ing.Protein,
				Lipits:      ing.Lipits,
				Glucids:     ing.Glucids,
				Canxi:       ing.Canxi,
				Phosphor:    ing.Phosphor,
				Fe:          ing.Fe,
				VitaminA:    ing.VitaminA,
				VitaminB1:   ing.VitaminB1,
				VitaminB2:   ing.VitaminB2,
				VitaminC:    ing.VitaminC,
				BetaCaroten: ing.BetaCaroten,
				VitaminPp:   ing.VitaminPp,
			}

			ingredientModels[i] = &models.IngredientWithUnit{
				ID:               ing.ID,
				Name:             ing.Name,
				Description:      ing.Description,
				Removal:          ing.Removal,
				Unit:             ing.Unit,
				NutritionPer100g: &nutritionPer100g,
			}
			category, err := r.db.GetCategoryById(ctx, ing.CategoryID)
			if err != nil {
				return nil, err
			}

			ingredientModels[i].Category = models.ToCategory(category)
		}

		result[i].Ingredients = ingredientModels
		result[i] = result[i].CalculateNutritionalValues()
	}

	return result, nil
}

func (r *DishRepository) CountDishesByIngredientID(
	ctx context.Context,
	id uuid.UUID,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountDishesByIngredientID").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountDishesByIngredientId(ctx, id)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
