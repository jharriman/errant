package errant

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TestType ErrorType = "test_type"
	BadType  ErrorType = "bad_type"
)

// Is should return true when matching two concrete Errs
func TestIsStruct(t *testing.T) {
	err1 := NewTypedError(nil, TestType, "test error 1")
	err2 := NewTypedError(nil, TestType, "test error 2")
	require.True(t, Is(err1, err2))

	err3 := NewTypedError(nil, BadType, "test error 3")
	require.False(t, Is(err1, err3))
}

// Is should return true when matching an Err against a ErrorType.
func TestIsErrorType(t *testing.T) {
	err := NewTypedError(nil, TestType, "test error")
	require.True(t, Is(err, TestType))
	require.False(t, Is(err, BadType))
}
