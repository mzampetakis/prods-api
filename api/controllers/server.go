package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mzampetakis/prods-api/api/controllers/middlewares"
	"github.com/mzampetakis/prods-api/api/services"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/redis"
)

type Handler struct {
	AppServices services.FunctionalitiesIface
}

func (h *Handler) ServerRun(addr string, prefix string) {
	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api-docs.json", swagger).Methods("GET")
	apiRouter := router.PathPrefix(prefix).Subrouter().StrictSlash(true)

	ringOpt := &redis.RingOptions{
		Addrs: map[string]string{
			"server": ":6379",
		},
	}
	cacheClient, _ := cache.NewClient(
		cache.ClientWithAdapter(redis.NewAdapter(ringOpt)),
		cache.ClientWithTTL(time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)

	apiRouter.Use(middlewares.LogRequest)
	h.initializeRoutes(apiRouter, cacheClient)
	fmt.Println("Listening at " + addr)
	logrus.Fatal(http.ListenAndServe(addr, router))
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "docs/swagger.json")
}
