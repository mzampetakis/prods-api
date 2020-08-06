package repositories

import (
	"context"
	"database/sql"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mzampetakis/prods-api/api/app"
	"github.com/sirupsen/logrus"
)

type DatastoreIface interface {
	GetProducts(context.Context, app.Filter) ([]*ProductFetchModel, error)
	GetProduct(context.Context, int64) (*ProductFetchModel, error)
	CreateProduct(context.Context, ProductCreateModel) (int64, error)
	UpdateProduct(context.Context, int64, ProductCreateModel) error
	DeleteProduct(context.Context, int64) error
	AssignProductsToCategory(context.Context, int64, ProductsCategoryUpdateModel) error

	GetCategories(context.Context, app.Filter) ([]*CategoryFetchModel, error)
	GetCategory(context.Context, int64) (*CategoryFetchModel, error)
	CreateCategory(context.Context, CategoryCreateModel) (int64, error)
	UpdateCategory(context.Context, int64, CategoryCreateModel) error
	DeleteCategory(context.Context, int64) error
}

type DB struct {
	*sql.DB
}

func NewDB(DBType string, DBURL string) (*DB, error) {
	var db *sql.DB
	db, err := sql.Open(DBType, DBURL)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	myDB := &DB{db}
	return myDB, nil
}

func (db *DB) MigrateDB() error {
	logrus.Info("Migrating DB schema...")
	stmt, err := ioutil.ReadFile("api/repositories/schema_script.sql")
	if err != nil {
		panic(err)
	}
	db.executeCommands(string(stmt))
	logrus.Info("DB Schema Migrated.")
	return nil
}

func (db *DB) SeedData() error {
	logrus.Info("Seeding sample data into DB...")
	stmt, err := ioutil.ReadFile("api/repositories/data_script.sql")
	if err != nil {
		panic(err)
	}
	db.executeCommands(string(stmt))
	logrus.Info("Sample Data Seeding Completed.")
	return nil
}

func (db *DB) executeCommands(statements string) {
	for _, statement := range strings.Split(statements, ";") {
		if len(strings.Trim(statement, " ")) == 0 {
			continue
		}
		_, err := db.Exec(statement)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}
}
