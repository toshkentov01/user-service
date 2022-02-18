package newerrors

import "errors"

var (
	// ErrAlreadyExists ...
	ErrAlreadyExists = errors.New("error already exists")

	// ErrUsernameExists ...
	ErrUsernameExists = errors.New("username exists")

	// ErrEmailExists ...
	ErrEmailExists = errors.New("email exists")

	// ErrInvalidField ...
	ErrInvalidField = errors.New("invalid field for username/email")

	// ErrMaximumAmount ...
	ErrMaximumAmount = errors.New("maximum amount")

	// ErrNotEnoughCash ...
	ErrNotEnoughCash = errors.New("not enough cash")

	// ErrInvalidFieldForOperations ...
	ErrInvalidFieldForOperations = errors.New("invalid field for operation type")
)
