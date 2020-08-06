package services

import (
	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/repositories"
	"golang.org/x/net/context"
)

type FunctionalitiesIface interface {
	GetProducts(context.Context, app.Filter) ([]*repositories.ProductFetchModel, error)
	GetProduct(context.Context, int64) (*repositories.ProductFetchModel, error)
	CreateProduct(context.Context, repositories.ProductCreateModel) (int64, error)
	UpdateProduct(context.Context, int64, repositories.ProductCreateModel) error
	DeleteProduct(context.Context, int64) error
	AssignProductsToCategory(context.Context, int64, repositories.ProductsCategoryUpdateModel) error

	GetCategories(context.Context, app.Filter) ([]*repositories.CategoryFetchModel, error)
	GetCategory(context.Context, int64) (*repositories.CategoryFetchModel, error)
	CreateCategory(context.Context, repositories.CategoryCreateModel) (int64, error)
	UpdateCategory(context.Context, int64, repositories.CategoryCreateModel) error
	DeleteCategory(context.Context, int64) error
}

type Service struct {
	DB repositories.DatastoreIface
}
