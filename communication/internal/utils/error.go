package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func FromError(err error) error {
	var e *Error
	if !errors.As(err, &e) {
		return status.Error(codes.Internal, internalError)
	}

	switch e.Status {
	case ErrNotFound:
		return status.Error(codes.NotFound, e.Message)
	case ErrBadRequest:
		return status.Error(codes.InvalidArgument, e.Message)
	case ErrNotAllowed:
		return status.Error(codes.PermissionDenied, e.Message)
	default:
		return status.Error(codes.Internal, internalError)
	}
}
