package goerror

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	InvalidRequest         ErrorType = "invalid_request"
	MethodNotAllowed       ErrorType = "method_not_allowed"
	ServerError            ErrorType = "server_error"
	UnsupportedContentType ErrorType = "unsupported_content_type"
)

type StatusError struct {
	statusCode int
	*Error
}

func (s *StatusError) WriteHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.statusCode)
	if err := json.NewEncoder(w).Encode(s.Error); err != nil {
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
		Error: &Error{
			ErrorType:        MethodNotAllowed,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewServerError(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusInternalServerError,
		Error: &Error{
			ErrorType:        ServerError,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewUnsupportedContentType(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusUnsupportedMediaType,
		Error: &Error{
			ErrorType:        UnsupportedContentType,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func NewInvalidRequest(cause error, description string, options ...Option) *StatusError {
	return &StatusError{
		statusCode: http.StatusBadRequest,
		Error: &Error{
			ErrorType:        InvalidRequest,
			ErrorDescription: description,
			Cause:            getAPIError(cause),
		},
	}
}

func getAPIError(err error) *Error {
	if err == nil {
		return nil
	}
	apiErr, isErrorType := err.(*Error)
	if !isErrorType {
		apiErr = NewUntypedError(err)
	}
	return apiErr
}
