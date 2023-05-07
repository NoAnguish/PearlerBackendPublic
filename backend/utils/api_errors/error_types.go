package api_errors

import (
	"errors"
	"net/http"
)

func NewInternalServerError(err error, message string) *Error {
	return &Error{
		Err:        err,
		statusCode: http.StatusInternalServerError,
		message:    message,
	}
}

func NewInternalDatabaseError(err error) *Error {
	return &Error{
		Err:        err,
		statusCode: http.StatusInternalServerError,
		message:    "internal database error occured",
	}
}

func NewDatabaseConnectionError(err error) *Error {
	return &Error{
		Err:        err,
		statusCode: http.StatusInternalServerError,
		message:    "could not establish connection with database",
	}
}

func NewS3UploadError(err error) *Error {
	return &Error{
		Err:        err,
		statusCode: http.StatusInternalServerError,
		message:    "error occurred while uploading to s3",
	}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Err:        errors.New(message),
		statusCode: http.StatusBadRequest,
		message:    message,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Err:        errors.New(message),
		statusCode: http.StatusNotFound,
		message:    message,
	}
}
