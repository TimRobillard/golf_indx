package errors

import (
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int               `json:"statusCode"`
	Msg        string            `json:"msg"`
	Context    map[string]string `json:"context"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InternalServerError(err error) APIError {
	slog.Error(err.Error())
	return APIError{
		StatusCode: http.StatusInternalServerError,
		Msg:        "Something went wrong",
		Context:    map[string]string{"msg": err.Error()},
	}
}

func BadRequestError(msg string, context map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Msg:        msg,
		Context:    context,
	}
}

func UnauthorizedError(msg string, context map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnauthorized,
		Msg:        msg,
		Context:    context,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

func NotImplementedError() APIError {
	return NewAPIError(http.StatusNotImplemented, fmt.Errorf("not implemented"))
}
