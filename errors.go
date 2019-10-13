// This package wraps the go errors library, the default error library
// does not have support for cleanly 'Wrap'-ing akin to github.com/pkg/errors.
package goerror

import (
	"errors"
	"fmt"
)

type wrapped struct {
	msg   string
	cause error
}

func (w wrapped) Error() string {
	return fmt.Sprintf("%s: %s", w.msg, w.cause.Error())
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
