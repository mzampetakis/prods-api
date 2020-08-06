package services

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/repositories"
	"golang.org/x/net/context"
)

type DBMock struct{}

func (db *DBMock) GetCategories(ctx context.Context, filter app.Filter) ([]*repositories.CategoryFetchModel, error) {
	var categories []*repositories.CategoryFetchModel
	categoryTitle := "Laptops"
	categoryImageURL := "https://category200.image"
	categorySort := int64(2)
	categories = append(categories, &repositories.CategoryFetchModel{
		ID:        200,
		Title:     &categoryTitle,
		ImageURL:  &categoryImageURL,
		Sort:      &categorySort,
		CreatedAt: "2020-05-25 21:02:15",
		UpdatedAt: "2020-05-25 21:05:15",
	})

	return categories, nil
}

func (db *DBMock) GetCategory(ctx context.Context, ID int64) (*repositories.CategoryFetchModel, error) {
	if ID == 201 {
		categoryTitle := "Monitors"
		categoryImageURL := "https://category201.image"
		categorySort := int64(3)
		return &repositories.CategoryFetchModel{
			ID:        201,
			Title:     &categoryTitle,
			ImageURL:  &categoryImageURL,
			Sort:      &categorySort,
			CreatedAt: "2020-05-24 21:02:15",
			UpdatedAt: "2020-05-24 21:05:15",
		}, nil
	}

	return nil, &app.Error{Op: "repositories.getCategory", Code: app.ENOTFOUND, Err: sql.ErrNoRows}
}

func (db *DBMock) CreateCategory(ctx context.Context, newCategory repositories.CategoryCreateModel) (int64, error) {
	return 201, nil
}

func (db *DBMock) UpdateCategory(ctx context.Context, categoryID int64, updateCategory repositories.CategoryCreateModel) error {
	return nil
}

func (db *DBMock) DeleteCategory(ctx context.Context, categoryID int64) error {
	return nil
}

func (db *DBMock) GetProducts(ctx context.Context, filter app.Filter) ([]*repositories.ProductFetchModel, error) {
	var products []*repositories.ProductFetchModel
	productTitle := "Flash Drive 1TB"
	productImageURL := "https://product200.image"
	productPrice := int64(1050)
	productCategory := int64(201)
	products = append(products, &repositories.ProductFetchModel{
		ID:         200,
		Title:      &productTitle,
		ImageURL:   &productImageURL,
		Price:      &productPrice,
		CategoryID: &productCategory,
		CreatedAt:  "2020-05-25 21:02:15",
		UpdatedAt:  "2020-05-25 21:05:15",
	})

	return products, nil
}

func (db *DBMock) GetProduct(ctx context.Context, ID int64) (*repositories.ProductFetchModel, error) {
	if ID == 201 {
		productTitle := "Flash Drive 1TB"
		productImageURL := "https://product200.image"
		productPrice := int64(1050)
		productCategory := int64(201)
		return &repositories.ProductFetchModel{
			ID:         201,
			Title:      &productTitle,
			ImageURL:   &productImageURL,
			Price:      &productPrice,
			CategoryID: &productCategory,
			CreatedAt:  "2020-05-25 21:02:15",
			UpdatedAt:  "2020-05-25 21:05:15",
		}, nil
	}

	return nil, &app.Error{Op: "repositories.Getproduct", Code: app.ENOTFOUND, Err: sql.ErrNoRows}
}

func (db *DBMock) CreateProduct(ctx context.Context, newProduct repositories.ProductCreateModel) (int64, error) {
	return 201, nil
}

func (db *DBMock) UpdateProduct(ctx context.Context, productID int64, updateProduct repositories.ProductCreateModel) error {
	return nil
}

func (db *DBMock) DeleteProduct(ctx context.Context, productID int64) error {
	return nil
}

func (db *DBMock) AssignProductsToCategory(ctx context.Context, categoryID int64, productsCategory repositories.ProductsCategoryUpdateModel) error {
	return nil
}

func TestGetCategories(t *testing.T) {
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())
	filter := app.Filter{}
	categs, err := mockService.GetCategories(ctx, filter)
	if err != nil {
		t.Errorf("Expected success but got error %s", err.Error())
	}
	if len(categs) != 1 {
		t.Errorf("Should get a Category list with 1 element but got %d", len(categs))
	}
}

func TestGetCategory(t *testing.T) {
	tests := map[string]struct {
		ID  int64
		err error
	}{
		"Category found":     {ID: 201, err: nil},
		"Category not found": {ID: 404, err: sql.ErrNoRows},
	}
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())

	for tName, tc := range tests {
		t.Run(tName, func(t *testing.T) {
			categ, err := mockService.GetCategory(ctx, tc.ID)
			if tc.err == nil && err != nil {
				t.Errorf("Expected success but got error %s", err.Error())
			}
			if tc.err != nil && err == nil {
				t.Errorf("Expected error %s but got no error", tc.err.Error())
			}
			if err != nil && tc.err != nil {
				if strings.Contains("err.Error()", "tc.err.Error()") {
					t.Errorf("Expected error %s but got error %s", err.Error(), tc.err.Error())
				}
			}
			if err == nil && categ.ID != tc.ID {
				t.Errorf("Should get a Category with ID %d but got ID %d", tc.ID, categ.ID)
			}
		})
	}
}

func TestCreateCategory_WhenNoTitleProvided_Fails(t *testing.T) {
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())
	categoryTitle := ""
	newCategory := repositories.CategoryCreateModel{
		Title: &categoryTitle,
	}

	categoryID, err := mockService.CreateCategory(ctx, newCategory)
	if err == nil {
		t.Errorf("Expected error but got no error")
	}
	if app.ErrorCode(err) != app.EINVALID {
		t.Errorf("Expected error code  %s, but got error code %s", app.ErrorCode(err), app.EINVALID)
	}
	excpectedCategoryID := int64(-1)
	if categoryID != excpectedCategoryID {
		t.Errorf("Expected categoryID to be %d but got %d", excpectedCategoryID, categoryID)
	}
}

func TestCreateCategory_WhenValidCategoryProvided_Succeeds(t *testing.T) {
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())
	categoryTitle := "some random title"
	newCategory := repositories.CategoryCreateModel{
		Title: &categoryTitle,
	}

	categoryID, err := mockService.CreateCategory(ctx, newCategory)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	excpectedCategoryID := int64(201)
	if categoryID <= 0 {
		t.Errorf("Expected categoryID to be %d but got %d", excpectedCategoryID, categoryID)
	}
}

func TestCreateProduct_WhenCategoryNotExists_Fails(t *testing.T) {
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())
	productTitle := "some random title"
	categoryID := int64(404)
	productPrice := int64(100)
	newProduct := repositories.ProductCreateModel{
		Title:      &productTitle,
		CategoryID: &categoryID,
		Price:      &productPrice,
	}

	productID, err := mockService.CreateProduct(ctx, newProduct)
	if err == nil {
		t.Errorf("Expected error but didn't got one")
	}
	if app.ErrorCode(err) != app.EINVALID {
		t.Errorf("Expected error code  %s, but got error code %s", app.ErrorCode(err), app.EINVALID)
	}
	excpectedProductID := int64(-1)
	if productID != excpectedProductID {
		t.Errorf("Expected productID to be %d but got %d", excpectedProductID, productID)
	}
}

func TestCreateProduct_WhenCategoryNotProvided_Succeeds(t *testing.T) {
	db := DBMock{}
	mockService := &Service{DB: &db}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", uuid.New())
	productTitle := "some random title"
	productPrice := int64(100)
	newProduct := repositories.ProductCreateModel{
		Title: &productTitle,
		Price: &productPrice,
	}

	productID, err := mockService.CreateProduct(ctx, newProduct)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	excpectedProductID := int64(201)
	if productID <= 0 {
		t.Errorf("Expected productID to be %d but got %d", excpectedProductID, productID)
	}
}
