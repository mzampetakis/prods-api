package services

import (
	"strings"

	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/repositories"
	"golang.org/x/net/context"
)

func (s *Service) GetProducts(ctx context.Context, filter app.Filter) ([]*repositories.ProductFetchModel, error) {
	if filter.Limit <= 0 {
		filter.Limit = 3
	}
	if len(filter.SortBy) == 0 {
		filter.SortBy = "id"
	}
	filter.SortDirection = strings.ToUpper(filter.SortDirection)
	if filter.SortDirection != "" && filter.SortDirection != app.ASC && filter.SortDirection != app.DESC {
		return nil, &app.Error{Op: "services.GetProducts", Code: app.EINVALID, Message: "Invalid SortDirection field: " + filter.SortDirection}
	}

	prods, err := s.DB.GetProducts(ctx, filter)
	if err != nil {
		return nil, &app.Error{Op: "services.GetProducts", Err: err}
	}
	return prods, nil
}

func (s *Service) GetProduct(ctx context.Context, productID int64) (*repositories.ProductFetchModel, error) {
	prod, err := s.DB.GetProduct(ctx, productID)
	if err != nil {
		return nil, &app.Error{Op: "services.GetProduct", Err: err}
	}
	return prod, nil
}

func (s *Service) CreateProduct(ctx context.Context, product repositories.ProductCreateModel) (int64, error) {
	if product.Title == nil || len(*product.Title) == 0 {
		return -1, &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Title cannot be empty."}
	}
	if product.Price == nil {
		return -1, &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Price cannot be empty."}
	}
	if product.CategoryID != nil {
		category, err := s.DB.GetCategory(ctx, *product.CategoryID)
		if err != nil || category.ID != *product.CategoryID {
			return -1, &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Invalid Category."}
		}
	}
	insertedID, err := s.DB.CreateProduct(ctx, product)
	if err != nil {
		return -1, &app.Error{Op: "services.CreateProduct", Err: err}
	}
	return insertedID, nil
}

func (s *Service) UpdateProduct(ctx context.Context, productID int64, product repositories.ProductCreateModel) error {
	if product.Title == nil || len(*product.Title) == 0 {
		return &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Title cannot be empty."}
	}
	if product.Price == nil {
		return &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Price cannot be empty."}
	}
	if product.CategoryID != nil {
		category, err := s.DB.GetCategory(ctx, *product.CategoryID)
		if err != nil || category.ID != *product.CategoryID {
			return &app.Error{Op: "services.CreateProduct", Code: app.EINVALID, Message: "Invalid Category."}
		}
	}
	err := s.DB.UpdateProduct(ctx, productID, product)
	if err != nil {
		return &app.Error{Op: "services.UpdateProduct", Err: err}
	}
	return nil
}

func (s *Service) DeleteProduct(ctx context.Context, productID int64) error {
	err := s.DB.DeleteProduct(ctx, productID)
	if err != nil {
		return &app.Error{Op: "services.DeleteProduct", Err: err}
	}
	return nil
}

func (s *Service) AssignProductsToCategory(ctx context.Context, CategoryID int64, productsCategory repositories.ProductsCategoryUpdateModel) error {
	category, err := s.DB.GetCategory(ctx, CategoryID)
	if err != nil || category.ID != CategoryID {
		return &app.Error{Op: "services.AssignProductsToCategory", Code: app.EINVALID, Message: "Invalid Category."}
	}
	err = s.DB.AssignProductsToCategory(ctx, CategoryID, productsCategory)
	if err != nil {
		return &app.Error{Op: "services.AssignProductsToCategory", Err: err}
	}
	return nil
}
