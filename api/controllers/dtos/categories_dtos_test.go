package dtos

import (
	"reflect"
	"testing"

	"github.com/mzampetakis/prods-api/api/repositories"
)

func TestConvertCategoryResponseModelToDto(t *testing.T) {
	//Prepare
	categoryID := int64(1)
	categoryTitle := "some title"
	categoryImageURL := "https://some.image"
	categorySort := int64(3)
	categoryCreatedAt := "2020-05-25T20:06:40+03:00"
	categoryUpdatedAt := "2020-05-25T20:06:40+03:00"
	categoryFetchModel := repositories.CategoryFetchModel{
		ID:        categoryID,
		Title:     &categoryTitle,
		ImageURL:  &categoryImageURL,
		Sort:      &categorySort,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}
	excpectedCategoryResponseDto := CategoryResponseDto{
		ID:        categoryID,
		Title:     &categoryTitle,
		ImageURL:  &categoryImageURL,
		Sort:      &categorySort,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}

	//Act
	categoryResponseDto := ConvertCategoryResponseModelToDto(categoryFetchModel)

	//Assert
	if !reflect.DeepEqual(excpectedCategoryResponseDto, categoryResponseDto) {
		t.Errorf("Excpected CategoryResponseDto %v but got %v", excpectedCategoryResponseDto, categoryResponseDto)
	}
}

func TestConvertCategoryRequestDtoToModel(t *testing.T) {
	//Prepare
	categoryTitle := "some title"
	categoryImageURL := "https://some.image"
	categorySort := int64(3)
	categoryRequestDto := CategoryRequestDto{
		Title:    &categoryTitle,
		ImageURL: &categoryImageURL,
		Sort:     &categorySort,
	}
	excpectedCategoryCreateModel := repositories.CategoryCreateModel{
		Title:    &categoryTitle,
		ImageURL: &categoryImageURL,
		Sort:     &categorySort,
	}

	//Act
	categoryCreateModel := ConvertCategoryRequestDtoToModel(categoryRequestDto)

	//Assert
	if !reflect.DeepEqual(excpectedCategoryCreateModel, categoryCreateModel) {
		t.Errorf("Excpected CategoryResponseDto %v but got %v", excpectedCategoryCreateModel, categoryCreateModel)
	}
}
