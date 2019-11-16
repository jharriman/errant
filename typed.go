package errant

type ErrorType string

func (et ErrorType) Type() ErrorType {
	return et
}

func (et ErrorType) Error() string {
	return string(et)
}

const Untyped ErrorType = "untyped"

type Err struct {
	ErrorType        ErrorType              `json:"error"`
	ErrorDescription string                 `json:"error_description,omitempty"`
	ErrorData        map[string]interface{} `json:"error_data,omitempty"`
	Cause            *Err                   `json:"error_cause,omitempty"`
}

func (e *Err) Type() ErrorType {
	return e.ErrorType
}

func (e *Err) Error() string {
	return string(e.ErrorType)
}

type typedError interface {
	Type() ErrorType
}

func (e *Err) Is(err error) bool {
	if apiErr, ok := err.(typedError); ok {
		return e.ErrorType == apiErr.Type()
	}
	return false
}

func NewUntypedError(err error) *Err {
	return &Err{
		ErrorType:        Untyped,
		ErrorDescription: err.Error(),
	}
}

func NewTypedError(cause error, errorType ErrorType, description string) *Err {
	var apiErr *Err
	if cause != nil {
		var ok bool
		apiErr, ok = cause.(*Err)
		if !ok {
			apiErr = NewUntypedError(cause)
		}
	}
	return &Err{
		ErrorType:        errorType,
		ErrorDescription: description,
		Cause:            apiErr,
	}
}
