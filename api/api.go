package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mzampetakis/prods-api/api/controllers"
	"github.com/mzampetakis/prods-api/api/repositories"
	"github.com/mzampetakis/prods-api/api/services"
	"github.com/sirupsen/logrus"
)

func init() {
	loadEnvVars()

}
func loadEnvVars() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Warn(".env file not found")
	}
}

func Run() {
	dbConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	db, err := repositories.NewDB("mysql", dbConnectionURL)
	if err != nil {
		logrus.Errorf("Could not connect ot DB: %s", err.Error())
		return
	}
	defer db.Close()
	if os.Getenv("MIGRATE_DB") == "true" {
		db.MigrateDB()
	}
	if os.Getenv("SEED_DATA") == "true" {
		db.SeedData()
	}
	sv := &services.Service{DB: db}
	h := controllers.Handler{AppServices: sv}
	h.ServerRun(":"+os.Getenv("SERVER_PORT"), os.Getenv("API_PREFIX"))
}
