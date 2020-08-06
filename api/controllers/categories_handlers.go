package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/controllers/dtos"
	"github.com/sirupsen/logrus"
)

// GetAllCategories godoc
// Id GetAllCategories
// @Summary Retrives Categories - uses filtering
// @Description Retrieve a list of Categories
// @Tags Categories
// @Produce json
// @Param offset query integer false "Offset of the results"
// @Param limit query integer false "Limit the results"
// @Param sortby query string false "Sort by of the results"
// @Param sortdirection query string false "Sort direction of the results (ASC|DESC)"
// @Success 200 {object} dtos.CategoriesResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /categories [get]
func (h *Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	filter := new(app.Filter)
	r.ParseForm()
	schema.NewDecoder().Decode(filter, r.Form)
	categories, err := h.AppServices.GetCategories(r.Context(), *filter)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetAllCategories", Err: err})
		return
	}
	dtos.JSON(w, http.StatusOK, dtos.ConvertCategoriesResponseModelToDto(categories))
}

// GetCategory godoc
// Id GetCategory
// @Summary Retrives single Category
// @Description Retrieves a Category
// @Tags Categories
// @Produce json
// @Param category_id path integer true "Category ID to retrieve"
// @Success 200 {object} dtos.CategoryResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 404 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /categories/{category_id} [get]
func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetCategory", Err: err, Code: app.EINVALID})
		return
	}
	category, err := h.AppServices.GetCategory(r.Context(), categoryID)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetCategory", Err: err})
		return
	}
	dtos.JSON(w, http.StatusOK, dtos.ConvertCategoryResponseModelToDto(*category))

}

// CreateCategory godoc
// Id CreateCategory
// @Summary Creates a Category
// @Description Create a new Category
// @Tags Categories
// @Produce json
// @Param category body dtos.CategoryRequestDto true "Category's data to create"
// @Success 201 {object} dtos.CreateCategoryResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /categories [post]
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory dtos.CategoryRequestDto
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.CreateCategory", Code: app.EINVALID, Err: err, Message: "Data validation error."})
		return
	}
	insertedID, err := h.AppServices.CreateCategory(r.Context(), dtos.ConvertCategoryRequestDtoToModel(newCategory))
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.CreateCategory", Err: err})
		return
	}
	dtos.JSON(w, http.StatusCreated, dtos.ConvertCreateCategoryResponseModelToDto(insertedID))
}

// UpdateCategory godoc
// Id UpdateCategory
// @Summary Updates a Category
// @Description Updates a Category
// @Tags Categories
// @Produce json
// @Param category_id path integer true "Category ID to update"
// @Param category body dtos.CategoryRequestDto true "Category's data to update"
// @Success 204
// @Failure 400 {object} dtos.ServeError
// @Failure 404 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /categories/{category_id} [put]
func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateCategory", Code: app.EINVALID, Err: err})
		return
	}
	var updateCategory dtos.CategoryRequestDto
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateCategory", Code: app.EINVALID, Err: err})
		return
	}
	err = h.AppServices.UpdateCategory(r.Context(), categoryID, dtos.ConvertCategoryRequestDtoToModel(updateCategory))
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateCategory", Err: err})
		return
	}
	dtos.JSON(w, http.StatusNoContent, nil)
}

// DeleteCategory godoc
// Id DeleteCategory
// @Summary Deletes a Category
// @Description Deletes a Category
// @Tags Categories
// @Produce json
// @Param category_id path integer true "Category ID to delete"
// @Success 204
// @Failure 500 {object} dtos.ServeError
// @Router /categories/{category_id} [delete]
func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.DeleteCategory", Code: app.EINVALID, Err: err})
		return
	}

	err = h.AppServices.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.DeleteCategory", Err: err})
		return
	}
	dtos.JSON(w, http.StatusNoContent, nil)
}
