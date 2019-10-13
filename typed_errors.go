package goerror

type ErrorType string

const Untyped ErrorType = "untyped"

type Error struct {
	ErrorType        ErrorType              `json:"error"`
	ErrorDescription string                 `json:"error_description,omitempty"`
	ErrorData        map[string]interface{} `json:"error_data,omitempty"`
	Cause            *Error                 `json:"error_cause,omitempty"`
}

func (e *Error) Type() ErrorType {
	return e.ErrorType
}

func (e *Error) Error() string {
	return string(e.ErrorType)
}

type typedError interface {
	Type() ErrorType
}

func (e *Error) Is(err error) bool {
	if apiErr, ok := err.(typedError); ok {
		return e.ErrorType == apiErr.Type()
	}
	return false
}

func NewUntypedError(err error) *Error {
	return &Error{
		ErrorType:        Untyped,
		ErrorDescription: err.Error(),
	}
}

func NewTypedError(cause error, errorType ErrorType, description string) *Error {
	var apiErr *Error
	if cause != nil {
		var ok bool
		apiErr, ok = cause.(*Error)
		if !ok {
			apiErr = NewUntypedError(cause)
		}
	}
	return &Error{
		ErrorType:        errorType,
		ErrorDescription: description,
		Cause:            apiErr,
	}
}
