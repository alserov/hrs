package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Error struct {
	Message string
	Status  int
}

const (
	ErrNotFound int = iota
	ErrBadRequest
	ErrNotAllowed
	ErrInternal
)

func (e *Error) Error() string {
	return e.Message
}

func NewError(status int, msg string) error {
	return &Error{
		Message: msg,
		Status:  status,
	}
}

const internalError = "internal error"

func FromGRPCError(err error) (int, string) {
	st, _ := status.FromError(err)
	switch st.Code() {
	case codes.NotFound:
		return http.StatusNotFound, st.Message()
	case codes.InvalidArgument:
		return http.StatusBadRequest, st.Message()
	case codes.PermissionDenied:
		return http.StatusMethodNotAllowed, st.Message()
	default:
		return http.StatusInternalServerError, internalError
	}
}

func FromError(err error) (int, string) {
	var e *Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError, internalError
	}

	switch e.Status {
	case ErrNotFound:
		return http.StatusNotFound, e.Message
	case ErrBadRequest:
		return http.StatusBadRequest, e.Message
	case ErrNotAllowed:
		return http.StatusMethodNotAllowed, e.Message
	default:
		return http.StatusInternalServerError, internalError
	}
}
