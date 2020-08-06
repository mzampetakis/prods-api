package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/mzampetakis/prods-api/api/app"
)

type ProductFetchModel struct {
	ID          int64   `json:"id"`
	CategoryID  *int64  `json:"category_id"`
	Title       *string `json:"title"`
	ImageURL    *string `json:"image_url"`
	Price       *int64  `json:"price"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductCreateModel struct {
	CategoryID  *int64  `json:"category_id"`
	Title       *string `json:"title"`
	ImageURL    *string `json:"image_url"`
	Price       *int64  `json:"price"`
	Description *string `json:"description"`
}

type ProductsCategoryUpdateModel []int64

func isProductsValidField(field string) bool {
	var model ProductFetchModel
	val := reflect.ValueOf(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if field == val.Type().Field(i).Tag.Get("json") {
			return true
		}
	}
	return false
}

func (db *DB) GetProducts(ctx context.Context, filter app.Filter) ([]*ProductFetchModel, error) {
	if !isProductsValidField(filter.SortBy) {
		return nil, &app.Error{Op: "repositories.GetProducts", Code: app.EINVALID, Message: "Invalid SortBy field: " + filter.SortBy}
	}
	query := fmt.Sprintf("SELECT id, category_id, title, image_url, price, description, created_at, updated_at FROM products ORDER BY %s %s LIMIT %d OFFSET %d",
		filter.SortBy, filter.SortDirection, filter.Limit, filter.Offset)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, &app.Error{Op: "repositories.GetProducts", Code: app.EINTERNAL, Err: err, Message: "Could not query Products from DB"}
	}
	defer rows.Close()

	prods := make([]*ProductFetchModel, 0)
	for rows.Next() {
		prod := new(ProductFetchModel)
		err := rows.Scan(&prod.ID, &prod.CategoryID, &prod.Title, &prod.ImageURL, &prod.Price, &prod.Description, &prod.CreatedAt, &prod.UpdatedAt)
		if err != nil {
			return nil, &app.Error{Op: "repositories.GetProducts", Code: app.EINTERNAL, Err: err, Message: "Could not fetch Products from DB"}
		}
		prods = append(prods, prod)
	}
	if err = rows.Err(); err != nil {
		return nil, &app.Error{Op: "repositories.GetProducts", Code: app.EINTERNAL, Err: err, Message: "Could not fetch Products from DB"}
	}
	return prods, nil
}

func (db *DB) GetProduct(ctx context.Context, productID int64) (*ProductFetchModel, error) {
	row := db.QueryRowContext(ctx, "SELECT id, category_id, title, image_url, price, description, created_at, updated_at FROM products WHERE id= ?",
		productID)

	prod := new(ProductFetchModel)
	err := row.Scan(&prod.ID, &prod.CategoryID, &prod.Title, &prod.ImageURL, &prod.Price, &prod.Description, &prod.CreatedAt, &prod.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &app.Error{Op: "repositories.GetProduct", Code: app.ENOTFOUND, Err: err, Message: "Product not found."}
		}
		return nil, &app.Error{Op: "repositories.GetProduct", Code: app.EINTERNAL, Err: err, Message: "Could not query Product from DB"}
	}
	return prod, nil
}

func (db *DB) CreateProduct(ctx context.Context, product ProductCreateModel) (int64, error) {
	res, err := db.ExecContext(ctx, "INSERT INTO products (category_id, title, image_url, price, description) VALUES (?, ?, ?, ?, ?)",
		product.CategoryID, product.Title, product.ImageURL, product.Price, product.Description)
	if err != nil {
		return -1, &app.Error{Op: "repositories.CreateProduct", Code: app.EINTERNAL, Err: err, Message: "Could not execute insert Product to DB"}
	}
	if rowsAffected, err := res.RowsAffected(); err != nil && rowsAffected != 1 {
		return -1, &app.Error{Op: "repositories.CreateProduct", Code: app.EINTERNAL, Err: err, Message: "Could not insert Product to DB"}
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		insertedID = 0
	}
	return insertedID, nil
}

func (db *DB) UpdateProduct(ctx context.Context, productID int64, product ProductCreateModel) error {
	res, err := db.ExecContext(ctx, "UPDATE products SET category_id=?, title=?, image_url=?, price=?, description=? WHERE id = ?",
		product.CategoryID, product.Title, product.ImageURL, product.Price, product.Description, productID)
	if err != nil {
		return &app.Error{Op: "repositories.UpdateProduct", Code: app.EINTERNAL, Err: err, Message: "Could not execute update Product in DB"}
	}
	if rowsAffected, err := res.RowsAffected(); err != nil && rowsAffected != 1 {
		return &app.Error{Op: "repositories.UpdateProduct", Code: app.EINTERNAL, Err: err, Message: "Could not update Product in DB"}
	}
	return nil
}

func (db *DB) DeleteProduct(ctx context.Context, productID int64) error {
	_, err := db.ExecContext(ctx, "DELETE FROM products WHERE id=?",
		productID)
	if err != nil {
		return &app.Error{Op: "repositories.DeleteProduct", Code: app.EINTERNAL, Err: err, Message: "Could not delete Product from DB"}
	}
	return nil
}

func (db *DB) AssignProductsToCategory(ctx context.Context, categoryID int64, productsCategory ProductsCategoryUpdateModel) error {
	query := fmt.Sprintf("UPDATE products SET category_id=? WHERE id IN (%s)", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(productsCategory)), ", "), "[]"))
	res, err := db.ExecContext(ctx, query,
		categoryID)
	if err != nil {
		return &app.Error{Op: "repositories.AssignProductsToCategory", Code: app.EINTERNAL, Err: err, Message: "Could not execute update Products' categories in DB"}
	}
	if rowsAffected, err := res.RowsAffected(); err != nil && rowsAffected != 1 {
		return &app.Error{Op: "repositories.AssignProductsToCategory", Code: app.EINTERNAL, Err: err, Message: "Could not update Products' categories in DB"}
	}
	return nil
}
