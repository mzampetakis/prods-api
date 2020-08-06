package app

import (
	"bytes"
	"fmt"
	"net/http"
)

// Filter is used in all GET listings
type Filter struct {
	Offset        int    `schema:"offset"`
	Limit         int    `schema:"limit"`
	SortBy        string `schema:"sortby"`
	SortDirection string `schema:"sortdirection"`
}

const (
	ASC  string = "ASC"
	DESC string = "DESC"
)

// Error is the way we pass and stack our errors
type Error struct {
	// Machine-readable error code.
	Code string
	// Human-readable message.
	Message string
	// Logical operation and nested error.
	Op string
	// Detailed Error Responce
	Err error
}

const (
	ECONFLICT    = "conflict"  // action cannot be performed
	EINTERNAL    = "internal"  // internal error
	EINVALID     = "invalid"   // validation failed
	ENOTFOUND    = "not_found" // entity does not exist
	ENOTACCEPTED = "not_accepted"
)

func StatusCode(err error) int {
	switch ErrorCode(err) {
	case ECONFLICT:
		return http.StatusConflict
	case EINTERNAL:
		return http.StatusInternalServerError
	case EINVALID:
		return http.StatusBadRequest
	case ENOTFOUND:
		return http.StatusNotFound
	case ENOTACCEPTED:
		return http.StatusNotAcceptable
	default:
		return http.StatusInternalServerError
	}
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
