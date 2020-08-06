package main

import (
	"github.com/mzampetakis/prods-api/api"
)

// @title API for prods-api
// @version 1.0
// @description This is the service that provides the API for prods-api.

// @contact.name Michalis Zampetakis
// @contact.email mzampetakis@gmail.com

// @host localhost:8080
// @BasePath /api
func main() {
	api.Run()
}
