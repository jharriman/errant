package errant

import (
	"errors"
	"fmt"
)

// While go adopted many great features of github.com/pkg/errors, they chose
// not to retain the concrete implementation of a 'wrapped' struct.
//
// The struct provided here allows you to use Wrap and Unwrap without
// writing your own implementation.
type wrapped struct {
	msg   string
	cause error
}

func (w wrapped) Error() string {
	if w.cause != nil {
		return fmt.Sprintf("%s: %s", w.msg, w.cause.Error())
	}
	return w.msg
}

func (w wrapped) Unwrap() error {
	return w.cause
}

func Wrap(cause error, msg string) error {
	return wrapped{msg: msg, cause: cause}
}

func Wrapf(cause error, msg string, args ...interface{}) error {
	return wrapped{msg: fmt.Sprintf(msg, args...), cause: cause}
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func New(text string) error {
	return errors.New(text)
}

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
