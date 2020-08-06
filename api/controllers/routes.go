package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mzampetakis/prods-api/api/controllers/middlewares"
	cache "github.com/victorspringer/http-cache"
)

func (h *Handler) initializeRoutes(router *mux.Router, cacheClient *cache.Client) {
	router.Use(middlewares.AcceptJSON)
	router.Use(middlewares.ContentTypeJSON)
	router.Use(middlewares.Recovery)
	router.Use(cacheClient.Middleware)

	// Home Route
	router.HandleFunc("/", h.Home).Methods("GET")

	// Products Routes
	router.HandleFunc("/products", h.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID:[0-9]+}", h.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{productID:[0-9]+}", h.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products/{productID:[0-9]+}", h.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/products/category/{categoryID:[0-9]+}", h.AssignProductsToCategory).Methods(http.MethodPut)

	// Categories Routes
	router.HandleFunc("/categories", h.GetAllCategories).Methods(http.MethodGet)
	router.HandleFunc("/categories/{categoryID:[0-9]+}", h.GetCategory).Methods(http.MethodGet)
	router.HandleFunc("/categories", h.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/categories/{categoryID:[0-9]+}", h.UpdateCategory).Methods(http.MethodPut)
	router.HandleFunc("/categories/{categoryID:[0-9]+}", h.DeleteCategory).Methods(http.MethodDelete)

}
