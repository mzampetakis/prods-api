package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/mzampetakis/prods-api/api/app"
)

type CategoryFetchModel struct {
	ID        int64   `json:"id"`
	Title     *string `json:"title"`
	ImageURL  *string `json:"image_url"`
	Sort      *int64  `json:"sort"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CategoryCreateModel struct {
	Title    *string `json:"title"`
	ImageURL *string `json:"image_url"`
	Sort     *int64  `json:"sort"`
}

func isCategoryValidField(field string) bool {
	var model CategoryFetchModel
	val := reflect.ValueOf(model)
	for i := 0; i < val.Type().NumField(); i++ {
		existing := val.Type().Field(i).Tag.Get("json")
		if field == existing {
			return true
		}
	}
	return false
}

func (db *DB) GetCategories(ctx context.Context, filter app.Filter) ([]*CategoryFetchModel, error) {
	if !isCategoryValidField(filter.SortBy) {
		return nil, &app.Error{Op: "repositories.GetCategories", Code: app.EINVALID, Message: "Invalid SortBy field: " + filter.SortBy}
	}
	query := fmt.Sprintf("SELECT id, title, image_url, sort, created_at, updated_at FROM categories ORDER BY %s %s LIMIT %d OFFSET %d",
		filter.SortBy, filter.SortDirection, filter.Limit, filter.Offset)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, &app.Error{Op: "repositories.GetCategories", Code: app.EINTERNAL, Err: err, Message: "Could not query Categories from DB"}
	}
	defer rows.Close()

	categs := make([]*CategoryFetchModel, 0)
	for rows.Next() {
		categ := new(CategoryFetchModel)
		err := rows.Scan(&categ.ID, &categ.Title, &categ.ImageURL, &categ.Sort, &categ.CreatedAt, &categ.UpdatedAt)
		if err != nil {
			return nil, &app.Error{Op: "repositories.GetCategories", Code: app.EINTERNAL, Err: err, Message: "Could not fetch Categories from DB"}
		}
		categs = append(categs, categ)
	}
	if err = rows.Err(); err != nil {
		return nil, &app.Error{Op: "repositories.GetCategories", Code: app.EINTERNAL, Err: err, Message: "Could not fetch Categories from DB"}
	}
	return categs, nil
}

func (db *DB) GetCategory(ctx context.Context, categoryID int64) (*CategoryFetchModel, error) {
	row := db.QueryRowContext(ctx, "SELECT id, title, image_url, sort, created_at, updated_at FROM categories WHERE id= ?",
		categoryID)
	categ := new(CategoryFetchModel)
	err := row.Scan(&categ.ID, &categ.Title, &categ.ImageURL, &categ.Sort, &categ.CreatedAt, &categ.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &app.Error{Op: "repositories.GetCategory", Code: app.ENOTFOUND, Err: err, Message: "Category not found."}
		}
		return nil, &app.Error{Op: "repositories.GetCategory", Code: app.EINTERNAL, Err: err, Message: "Could not query Category from DB"}
	}
	return categ, nil
}

func (db *DB) CreateCategory(ctx context.Context, category CategoryCreateModel) (int64, error) {
	res, err := db.ExecContext(ctx, "INSERT INTO categories (title, image_url, sort) VALUES (?, ?, ?)",
		category.Title, category.ImageURL, category.Sort)
	if err != nil {
		return -1, &app.Error{Op: "repositories.CreateCategory", Code: app.EINTERNAL, Err: err, Message: "Could not execute insert Category to DB"}
	}
	if rowsAffected, err := res.RowsAffected(); err != nil && rowsAffected != 1 {
		return -1, &app.Error{Op: "repositories.CreateCategory", Code: app.EINTERNAL, Err: err, Message: "Could not insert Category to DB"}
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		insertedID = 0
	}
	return insertedID, nil
}

func (db *DB) UpdateCategory(ctx context.Context, CategoryID int64, category CategoryCreateModel) error {
	res, err := db.ExecContext(ctx, "UPDATE categories SET title=?, image_url=?, sort=? WHERE id = ?",
		category.Title, category.ImageURL, category.Sort, CategoryID)
	if err != nil {
		return &app.Error{Op: "repositories.UpdateCategory", Code: app.EINTERNAL, Err: err, Message: "Could not execute update Category in DB"}
	}
	if rowsAffected, err := res.RowsAffected(); err != nil && rowsAffected != 1 {
		return &app.Error{Op: "repositories.UpdateCategory", Code: app.EINTERNAL, Err: err, Message: "Could not update Category in DB"}
	}
	return nil
}

func (db *DB) DeleteCategory(ctx context.Context, CategoryID int64) error {
	_, err := db.ExecContext(ctx, "DELETE FROM categories WHERE id=?",
		CategoryID)
	if err != nil {
		return &app.Error{Op: "repositories.DeleteCategory", Code: app.EINTERNAL, Err: err, Message: "Could not delete Category from DB"}
	}
	return nil
}
