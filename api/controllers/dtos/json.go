// Package dtos stores the API DTOs and functionalities to convert DTOs to Models and vice versa
// as well as functionality to serve json and error
package dtos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mzampetakis/prods-api/api/app"
)

// JSON responds to a request a json with the provided data alongside with the provided statusCode
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
	return
}

type ServeError struct {
	TraceID        string `json:"trace_id"`
	Timestamp      string `json:"timestamp"`
	Message        string `json:"message"`
	Code           string `json:"code"`
	HTTPStatusCode int    `json:"http_status_code"`
	HTTPStatus     string `json:"http_status"`
}

// ERROR responds to a request an error with the provided data
func ERROR(w http.ResponseWriter, ctx context.Context, err error) {
	JSON(w, app.StatusCode(err), ServeError{
		TraceID:        fmt.Sprintf("%v", ctx.Value("request_id")),
		Timestamp:      time.Now().Format(time.RFC3339),
		Message:        app.ErrorMessage(err),
		Code:           app.ErrorCode(err),
		HTTPStatusCode: app.StatusCode(err),
		HTTPStatus:     http.StatusText(app.StatusCode(err)),
	})
	return
}
