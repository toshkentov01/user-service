package service

import (
	"database/sql"
	"strings"

	newerrors "github.com/toshkentov01/alif-tech-task/user-service/new_errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	l "github.com/toshkentov01/alif-tech-task/user-service/pkg/logger"
)

const (
	// InternalServerError ...
	InternalServerError = "Internal Server Error"

	// AlreadyExistsError ...
	AlreadyExistsError = "Already Exists"

	// NotFoundError ...
	NotFoundError = "Not Found"

	// InvalidArgumentError ...
	InvalidArgumentError = "Invalid Argument"
)

// errorHandler function for handling errors in product service
func errorHandler(logger l.Logger, err error, message string, req ...interface{}) error {

	if err == nil {
		return nil

	} else if strings.Contains(err.Error(), "duplicate key value") {
		logger.Error(message+": Already exists, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.AlreadyExists, AlreadyExistsError)

	} else if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
		logger.Error(message+": Invalid argument, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.InvalidArgument, InvalidArgumentError)

	} else if err == newerrors.ErrAlreadyExists {
		logger.Error(message+": Already Exists, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.AlreadyExists, AlreadyExistsError)

	} else if err == sql.ErrNoRows {
		logger.Error(message+": Error No Rows, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.NotFound, NotFoundError)

	} else if err == newerrors.ErrInvalidField {
		logger.Error(message+": Invalid Field, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.InvalidArgument, InvalidArgumentError)

	} else if strings.Contains(err.Error(), "violates foreign key constraint") {
		logger.Error(message+"Invalid input"+err.Error(), l.Any("req", req))
		return status.Error(codes.InvalidArgument, "Invalid input")

	} else if err != nil {
		logger.Error(message+": Internal Server Error, error: "+err.Error(), l.Any("req", req))
		return status.Error(codes.Internal, InternalServerError)
	}

	return nil
}