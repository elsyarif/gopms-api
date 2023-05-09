package common

import (
	"errors"
)

const (
	NotFoundError                = "NotFound"
	notFoundErrorMessage         = "resource tidak ditemukan"
	ValidationError              = "ValidationError"
	validationErrorMessage       = "terjadi kesalahan pada request"
	NotAuthenticatedError        = "NotAuthenticated"
	notAuthenticatedErrorMessage = "tidak dapat diautentikasi"
	NotAuthorizedError           = "NotAuthorized"
	notAuthorizedErrorMessage    = "resource tidak diizinkan"
	UnknownError                 = "UnknownError"
	unknownErrorMessage          = "terjadi kesalahan pada server kami"
	ResourceAlreadyExists        = "ResourceAlreadyExists"
	alreadyExistsErrorMessage    = "resource already exists"
	InvalidTokenError            = "InvalidToken"
	InvalidTokenErrorMessage     = "invalid token"
)

type AppError struct {
	Errors error
	Type   string
}

func (app *AppError) Error() string {
	return app.Errors.Error()
}

func NewError(err error, errType string) *AppError {
	return &AppError{
		Errors: err,
		Type:   errType,
	}
}

func NewErrorType(errType string) *AppError {
	var err error

	switch errType {
	case NotFoundError:
		err = errors.New(notFoundErrorMessage)
	case ValidationError:
		err = errors.New(validationErrorMessage)
	case NotAuthenticatedError:
		err = errors.New(notAuthenticatedErrorMessage)
	case NotAuthorizedError:
		err = errors.New(notAuthorizedErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(alreadyExistsErrorMessage)
	case InvalidTokenError:
		err = errors.New(InvalidTokenErrorMessage)
	default:
		err = errors.New(unknownErrorMessage)
	}

	return &AppError{
		Errors: err,
		Type:   errType,
	}
}
