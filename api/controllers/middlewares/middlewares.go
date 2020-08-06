// Package middlewares provides middlewares for the usage in http
package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mzampetakis/prods-api/api/app"
	"github.com/mzampetakis/prods-api/api/controllers/dtos"
	"github.com/sirupsen/logrus"
)

// LogRequest logs each request with details
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New()
		ctx := context.WithValue(r.Context(), "request_id", requestID)
		logrus.Printf("%s : %s: Method: %s | URL: %s%s | Proto: %s",
			time.Now().Format(time.RFC3339), requestID, r.Method, r.Host, r.URL, r.Proto)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AcceptJSON stricts usage to accept only 'application/json' or '*/*' request header
func AcceptJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept"), "application/json") && !strings.Contains(r.Header.Get("Accept"), "*/*") {
			msg := "Header accept: '" + r.Header.Get("Accept") + "' is not accepted. application/json or */* should be accepted."
			dtos.ERROR(w, r.Context(), &app.Error{Op: "AcceptJSON", Err: errors.New(msg), Code: app.ENOTACCEPTED, Message: msg})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ContentTypeJSON adds 'Content-Type'='application/json' header to each response
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware allows a panic within an API call to recover
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			var err error
			recovery := recover()
			if recovery != nil {
				switch t := recovery.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				logrus.Error(err.Error())
				dtos.ERROR(w, r.Context(), &app.Error{Code: app.EINTERNAL, Err: err})
			}

		}()
		next.ServeHTTP(w, r)
	})
}
