package gg

import (
	"net/http"
	"strings"
)

type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

const (
	ErrorRecordAlreadyExists      = ConstError("record already exists")
	ErrorRecordInUsed             = ConstError("record in used")
	ErrorRecordNotFound           = ConstError("record not found")
	ErrorReferencedRecordNotFound = ConstError("referenced record not found")
)

const (
	SQLForeignKeyViolation = "SQLSTATE 23503"
	SQLUniqueViolation     = "SQLSTATE 23505"
)

func HandleRetrievalError(err error) error {
	switch {
	case err == nil:
		return nil
	case strings.Contains(err.Error(), ErrorRecordNotFound.Error()):
		return ErrorRecordNotFound
	default:
		return err
	}
}

func HandleCreationError(err error) error {
	switch {
	case err == nil:
		return nil
	case strings.Contains(err.Error(), SQLForeignKeyViolation):
		return ErrorReferencedRecordNotFound
	case strings.Contains(err.Error(), SQLUniqueViolation):
		return ErrorRecordAlreadyExists
	default:
		return err
	}
}

func HandleDeletionError(err error) error {
	switch {
	case err == nil:
		return nil
	case strings.Contains(err.Error(), SQLForeignKeyViolation):
		return ErrorRecordInUsed
	default:
		return err
	}
}

func HandleStatusCode(err error) int {
	switch err {
	case ErrorInvalidStringToIntConversion:
		return http.StatusBadRequest
	case ErrorInvalidBase64Encoding:
		return http.StatusUnprocessableEntity
	case ErrorRecordAlreadyExists, ErrorRecordInUsed:
		return http.StatusConflict
	case ErrorRecordNotFound:
		return http.StatusNotFound
	case ErrorReferencedRecordNotFound:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
