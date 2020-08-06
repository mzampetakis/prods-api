package controllers

import (
	"net/http"

	"github.com/mzampetakis/prods-api/api/controllers/dtos"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	dtos.JSON(w, http.StatusOK, "Welcome to prods-api. Please visit http://localhost:8080/swagger/ and use and use http://localhost:8080/api-docs.json  to get started")
}
