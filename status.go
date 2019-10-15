package goerror

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	InvalidRequest         ErrorType = "invalid_request"
	MethodNotAllowed       ErrorType = "method_not_allowed"
	NotFound               ErrorType = "not_found"
	ServerError            ErrorType = "server_error"
	Unauthorized           ErrorType = "unauthorized"
	UnsupportedContentType ErrorType = "unsupported_content_type"
)

type StatusError struct {
	statusCode int
	*Err
}

func (s *StatusError) WriteHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.statusCode)
	if err := json.NewEncoder(w).Encode(s.Err); err != nil {
		log.Println("Failed to decode request")
	}
}

type HTTPWriter interface {
	WriteHTTP(w http.ResponseWriter, r *http.Request)
}

type Option interface {
	Apply(s *StatusError)
}

func NewMethodNotAllowed(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusMethodNotAllowed,
		Err: &Err{
			ErrorType:        MethodNotAllowed,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewNotFound(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusNotFound,
		Err: &Err{
			ErrorType:        NotFound,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewServerError(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusInternalServerError,
		Err: &Err{
			ErrorType:        ServerError,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewUnauthorized(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusUnauthorized,
		Err: &Err{
			ErrorType:        Unauthorized,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewUnsupportedContentType(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusUnsupportedMediaType,
		Err: &Err{
			ErrorType:        UnsupportedContentType,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewInvalidRequest(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusBadRequest,
		Err: &Err{
			ErrorType:        InvalidRequest,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func getAPIError(err error) *Err {
	if err == nil {
		return nil
	}
	apiErr, isErrorType := err.(*Err)
	if !isErrorType {
		apiErr = NewUntypedError(err)
	}
	return apiErr
}
