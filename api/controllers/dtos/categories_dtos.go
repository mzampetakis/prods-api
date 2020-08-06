// Package dtos stores the API DTOs and functionalities to convert DTOs to Models and vice versa
// as well as functionality to serve json and error
package dtos

import (
	"github.com/mzampetakis/prods-api/api/repositories"
)

type CategoryResponseDto struct {
	ID        int64   `json:"id"`
	Title     *string `json:"title"`
	ImageURL  *string `json:"image_url"`
	Sort      *int64  `json:"sort"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CategoryRequestDto struct {
	Title    *string `json:"title"`
	ImageURL *string `json:"image_url"`
	Sort     *int64  `json:"sort"`
}

type CreateCategoryResponseDto struct {
	ID int64 `json:"id"`
}

type CategoriesResponseDto []CategoryResponseDto

func ConvertCategoryResponseModelToDto(category repositories.CategoryFetchModel) CategoryResponseDto {
	return CategoryResponseDto{
		ID:        category.ID,
		Title:     category.Title,
		ImageURL:  category.ImageURL,
		Sort:      category.Sort,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ConvertCategoryRequestDtoToModel(category CategoryRequestDto) repositories.CategoryCreateModel {
	return repositories.CategoryCreateModel{
		Title:    category.Title,
		ImageURL: category.ImageURL,
		Sort:     category.Sort,
	}
}

func ConvertCategoriesResponseModelToDto(categories []*repositories.CategoryFetchModel) CategoriesResponseDto {
	categoriesResponseDto := make(CategoriesResponseDto, 0)
	for _, category := range categories {
		categoriesResponseDto = append(categoriesResponseDto, ConvertCategoryResponseModelToDto(*category))
	}
	return categoriesResponseDto
}

func ConvertCreateCategoryResponseModelToDto(categoryID int64) CreateCategoryResponseDto {
	return CreateCategoryResponseDto{
		ID: categoryID,
	}
}
