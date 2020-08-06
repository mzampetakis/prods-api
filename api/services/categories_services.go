package services

import (
	"strings"

	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/repositories"
	"golang.org/x/net/context"
)

func (s *Service) GetCategories(ctx context.Context, filter app.Filter) ([]*repositories.CategoryFetchModel, error) {
	if filter.Limit <= 0 {
		filter.Limit = 3
	}
	if len(filter.SortBy) == 0 {
		filter.SortBy = "sort"
	}
	filter.SortDirection = strings.ToUpper(filter.SortDirection)
	if filter.SortDirection != "" && filter.SortDirection != app.ASC && filter.SortDirection != app.DESC {
		return nil, &app.Error{Op: "services.GetCategories", Code: app.EINVALID, Message: "Invalid SortDirection field: " + filter.SortDirection}
	}

	categs, err := s.DB.GetCategories(ctx, filter)
	if err != nil {
		return nil, &app.Error{Op: "services.GetCategories", Err: err}
	}
	return categs, nil
}

func (s *Service) GetCategory(ctx context.Context, categoryID int64) (*repositories.CategoryFetchModel, error) {
	categ, err := s.DB.GetCategory(ctx, categoryID)
	if err != nil {
		return nil, &app.Error{Op: "services.GetCategory", Err: err}
	}
	return categ, nil
}

func (s *Service) CreateCategory(ctx context.Context, category repositories.CategoryCreateModel) (int64, error) {
	if category.Title == nil || len(*category.Title) == 0 {
		return -1, &app.Error{Op: "services.CreateCategory", Code: app.EINVALID, Message: "Title cannot be empty."}
	}
	insertedID, err := s.DB.CreateCategory(ctx, category)
	if err != nil {
		return -1, &app.Error{Op: "services.CreateCategory", Err: err}
	}
	return insertedID, nil
}

func (s *Service) UpdateCategory(ctx context.Context, categoryID int64, category repositories.CategoryCreateModel) error {
	if category.Title == nil || len(*category.Title) == 0 {
		return &app.Error{Op: "services.UpdateCategory", Code: app.EINVALID, Message: "Title cannot be empty."}
	}
	err := s.DB.UpdateCategory(ctx, categoryID, category)
	if err != nil {
		return &app.Error{Op: "services.UpdateCategory", Err: err}
	}
	return nil
}

func (s *Service) DeleteCategory(ctx context.Context, categoryID int64) error {
	err := s.DB.DeleteCategory(ctx, categoryID)
	if err != nil {
		return &app.Error{Op: "services.DeleteCategory", Err: err}
	}
	return nil
}
