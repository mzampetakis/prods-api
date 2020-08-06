// Package dtos stores the API DTOs and functionalities to convert DTOs to Models and vice versa
// as well as functionality to serve json and error
package dtos

import (
	"github.com/mzampetakis/prods-api/api/repositories"
)

type ProductResponseDto struct {
	ID          int64   `json:"id"`
	CategoryID  *int64  `json:"category_id"`
	Title       *string `json:"title"`
	ImageURL    *string `json:"image_url"`
	Price       *int64  `json:"price"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductRequestDto struct {
	CategoryID  *int64  `json:"category_id"`
	Title       *string `json:"title"`
	ImageURL    *string `json:"image_url"`
	Price       *int64  `json:"price"`
	Description *string `json:"description"`
}
type CreateProductResponseDto struct {
	ID int64 `json:"id"`
}

type ProductsResponseDto []ProductResponseDto

type ProductsCategoryUpdateRequestDto struct {
	ProductIDs []int64 `json:"product_ids"`
}

func ConvertProductResponseModelToDto(product repositories.ProductFetchModel) ProductResponseDto {
	return ProductResponseDto{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Title:       product.Title,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ConvertProductRequestDtoToModel(product ProductRequestDto) repositories.ProductCreateModel {
	return repositories.ProductCreateModel{
		CategoryID:  product.CategoryID,
		Title:       product.Title,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		Description: product.Description,
	}
}

func ConvertProductsCategoryUpdateRequestDtoToModel(productsCategory ProductsCategoryUpdateRequestDto) repositories.ProductsCategoryUpdateModel {
	return repositories.ProductsCategoryUpdateModel(productsCategory.ProductIDs)
}

func ConvertProductsResponseModelToDto(products []*repositories.ProductFetchModel) ProductsResponseDto {
	productsResponseDto := make(ProductsResponseDto, 0)
	for _, product := range products {
		productsResponseDto = append(productsResponseDto, ConvertProductResponseModelToDto(*product))
	}
	return productsResponseDto
}

func ConvertCreateProductResponseModelToDto(productID int64) CreateProductResponseDto {
	return CreateProductResponseDto{
		ID: productID,
	}
}
