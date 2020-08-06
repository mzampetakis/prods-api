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

// GetAllProducts godoc
// Id GetAllProducts
// @Summary Retrive Products - uses filtering
// @Description Retrieve a list of products
// @Tags Products
// @Produce json
// @Param offset query integer false "Offset of the results"
// @Param limit query integer false "Limit the results"
// @Param sortby query string false "Sort by of the results"
// @Param sortdirection query string false "Sort direction of the results (ASC|DESC)"
// @Success 200 {object} dtos.ProductsResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /products [get]
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	filter := new(app.Filter)
	r.ParseForm()
	schema.NewDecoder().Decode(filter, r.Form)
	products, err := h.AppServices.GetProducts(r.Context(), *filter)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetAllProducts", Err: err})
		return
	}
	dtos.JSON(w, http.StatusOK, dtos.ConvertProductsResponseModelToDto(products))
}

// GetProduct godoc
// Id GetProduct
// @Summary Retrives single Product
// @Description Retrieve a Product
// @Tags Products
// @Produce json
// @Param product_id path integer true "Product ID to retrieve"
// @Success 200 {object} dtos.ProductResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 404 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /products/{product_id} [get]
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.ParseInt(params["productID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetProduct", Err: err, Code: app.EINVALID})
		return
	}
	product, err := h.AppServices.GetProduct(r.Context(), productID)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.GetProduct", Err: err})
		return
	}
	dtos.JSON(w, http.StatusOK, dtos.ConvertProductResponseModelToDto(*product))

}

// CreateProduct godoc
// Id CreateProduct
// @Summary Creates a Product
// @Description Create a new Product
// @Tags Products
// @Produce json
// @Param product body dtos.ProductRequestDto true "Product's data to create"
// @Success 201 {object} dtos.CreateProductResponseDto
// @Failure 400 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /products [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct dtos.ProductRequestDto
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.CreateProduct", Code: app.EINVALID, Err: err, Message: "Data validation error."})
		return
	}
	insertedID, err := h.AppServices.CreateProduct(r.Context(), dtos.ConvertProductRequestDtoToModel(newProduct))
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.CreateBook", Err: err})
		return
	}
	dtos.JSON(w, http.StatusCreated, dtos.ConvertCreateProductResponseModelToDto(insertedID))
}

// UpdateProduct godoc
// Id UpdateProduct
// @Summary Updates a Product
// @Description Update a Product
// @Tags Products
// @Produce json
// @Param product_id path integer true "Product ID to update"
// @Param product body dtos.ProductRequestDto true "Product's data to update"
// @Success 204
// @Failure 400 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /products/{product_id} [put]
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.ParseInt(params["productID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateProduct", Code: app.EINVALID, Err: err})
		return
	}
	var updateProduct dtos.ProductRequestDto
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateProduct", Code: app.EINVALID, Err: err})
		return
	}
	err = h.AppServices.UpdateProduct(r.Context(), productID, dtos.ConvertProductRequestDtoToModel(updateProduct))
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.UpdateProduct", Err: err})
		return
	}
	dtos.JSON(w, http.StatusNoContent, nil)
}

// DeleteProduct godoc
// Id DeleteProduct
// @Summary Deletes a Product
// @Description Deletes a Product
// @Tags Products
// @Produce json
// @Param product_id path integer true "Product ID to delete"
// @Success 204
// @Failure 500 {object} dtos.ServeError
// @Router /products/{product_id} [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.ParseInt(params["productID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.DeleteProduct", Code: app.EINVALID, Err: err})
		return
	}

	err = h.AppServices.DeleteProduct(r.Context(), productID)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.DeleteProduct", Err: err})
		return
	}
	dtos.JSON(w, http.StatusNoContent, nil)
}

// AssignProductsToCategory godoc
// Id AssignProductsToCategory
// @Summary Assing Products to a category
// @Description Assing Products to a category
// @Tags Products
// @Produce json
// @Param category_id path integer true "Category ID to assign products to"
// @Param products_category body dtos.ProductsCategoryUpdateRequestDto true "Products' ID to assign to the category"
// @Success 204
// @Failure 400 {object} dtos.ServeError
// @Failure 404 {object} dtos.ServeError
// @Failure 500 {object} dtos.ServeError
// @Router /products/category/{category_id} [put]
func (h *Handler) AssignProductsToCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryID"], 10, 64)
	if err != nil {
		logrus.Warn(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.AssignProductsToCategory", Code: app.EINVALID, Err: err})
		return
	}
	var productsCategoryUpdate dtos.ProductsCategoryUpdateRequestDto
	err = json.NewDecoder(r.Body).Decode(&productsCategoryUpdate)
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.AssignProductsToCategory", Code: app.EINVALID, Err: err})
		return
	}
	err = h.AppServices.AssignProductsToCategory(r.Context(), categoryID, dtos.ConvertProductsCategoryUpdateRequestDtoToModel(productsCategoryUpdate))
	if err != nil {
		logrus.Errorf(err.Error())
		dtos.ERROR(w, r.Context(), &app.Error{Op: "handlers.AssignProductsToCategory", Err: err})
		return
	}
	dtos.JSON(w, http.StatusNoContent, nil)
}
