package errorhandler

import (
	"errors"
	"net/http"
)

const (
	errBadRequest          = "Bad Request"
	errUnauthorized        = "Unauthorized"
	errForbidden           = "Forbidden"
	errUnprocessableEntity = "Unprocessable Entity"
	errInternalServerError = "Internal Server Error"
)

var (
	ErrBadRequest          = errors.New(errBadRequest)
	ErrUnauthorized        = errors.New(errUnauthorized)
	ErrForbidden           = errors.New(errForbidden)
	ErrUnprocessableEntity = errors.New(errUnprocessableEntity)
	ErrInternalServerError = errors.New(errInternalServerError)
)

type ErrorResponse struct {
	Status      int    `json:"status"`
	Error       string `json:"error"`
	ErrorDetail string `json:"error_detail"`
}

func ApiErrorHandle(errorKind error, msg string) *ErrorResponse {
	switch errorKind {
	case ErrBadRequest:
		return &ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       errBadRequest,
			ErrorDetail: msg,
		}
	case ErrUnauthorized:
		return &ErrorResponse{
			Status:      http.StatusUnauthorized,
			Error:       errUnauthorized,
			ErrorDetail: msg,
		}
	case ErrForbidden:
		return &ErrorResponse{
			Status:      http.StatusForbidden,
			Error:       errForbidden,
			ErrorDetail: msg,
		}
	case ErrUnprocessableEntity:
		return &ErrorResponse{
			Status:      http.StatusUnprocessableEntity,
			Error:       errUnprocessableEntity,
			ErrorDetail: msg,
		}
	default:
		return &ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       errInternalServerError,
			ErrorDetail: msg,
		}
	}
}
