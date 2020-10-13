package errs

import "net/http"

type ErrWrapper struct {
	Message string `json:"message" example:"INTERNAL_SERVER_ERROR"`
	Code    int    `json:"code" example:"500"`
}

func WrapError(err error, code int) ErrWrapper {
	return ErrWrapper{
		Message: err.Error(),
		Code:    code,
	}
}

var (
	InternalServerErr      = ErrWrapper{"INTERNAL_SERVER_ERROR", http.StatusInternalServerError}
	IncorrectBodyErr       = ErrWrapper{"INCORRECT_BODY", http.StatusBadRequest}
	UnauthorizedErr        = ErrWrapper{"UNAUTHORIZED", http.StatusUnauthorized}
	IncorrectMultipartData = ErrWrapper{"INCORRECT_MULTIPART_DATA", http.StatusBadRequest}
	StatusOK               = ErrWrapper{"OK", http.StatusOK}
	IncorrectQueryParams   = ErrWrapper{"INCORRECT_QUERY_PARAMS", http.StatusBadRequest}
)
