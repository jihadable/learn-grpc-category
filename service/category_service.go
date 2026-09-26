package service

import (
	"context"

	categoryPb "github.com/jihadable/learn-grpc-proto/category"
)

type CategoryService struct {
	categoryPb.UnimplementedCategoryServiceServer
}

func (service *CategoryService) CreateCategory(ctx context.Context, req *categoryPb.Category) (*categoryPb.Category, error) {
	category := &categoryPb.Category{
		Id:   1,
		Name: "Category 5",
	}

	return category, nil
}

func (service *CategoryService) GetCategories(ctx context.Context, _ *categoryPb.Empty) (*categoryPb.Categories, error) {
	categories := &categoryPb.Categories{}
	category := &categoryPb.Category{
		Id:   5,
		Name: "Category 8",
	}

	categories.Data = append(categories.Data, category)

	return categories, nil
}

func (service *CategoryService) GetCategory(ctx context.Context, id *categoryPb.Id) (*categoryPb.Category, error) {
	category := &categoryPb.Category{
		Id:   2,
		Name: "Category 9",
	}

	return category, nil
}

func (service *CategoryService) UpdateCategory(ctx context.Context, id *categoryPb.Id) (*categoryPb.Category, error) {
	category := &categoryPb.Category{
		Id:   4,
		Name: "Category 7",
	}

	return category, nil
}

func (service *CategoryService) DeleteCategory(ctx context.Context, id *categoryPb.Id) (*categoryPb.Status, error) {
	status := &categoryPb.Status{
		Status: 1,
	}

	return status, nil
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}
