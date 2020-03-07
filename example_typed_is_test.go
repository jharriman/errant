package errant

import (
	"fmt"
)

const PackageSpecificError ErrorType = "package_specific_error"

func Example_is() {
	underlyingErr := New("underlying error")
	err := NewTypedError(underlyingErr, PackageSpecificError, "a special type of error")

	fmt.Println(Is(err, PackageSpecificError))

	fmt.Println(Unwrap(err).Error())
	fmt.Println(Unwrap(err).(*Err).ErrorDescription)
	// Output:
	// true
	// untyped
	// underlying error
}
