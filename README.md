# errant
A package for faster error writing. Makes using `errors.Wrap` and `errors.Is` a breeze.

## Usage
### `Wrap` and `Unwrap`
```go
package main

import (
    "fmt"

    "github.com/jharriman/errant"
)

func main() {
    err := errant.New("base error")
    err2 := errant.Wrap(err, "level 1 error")
    err3 := errant.Wrap(err2, "level 2 error")

    fmt.Println(errant.Unwrap(err3))
    // Output: level 1 error: base error
}
```

### `Is` with `ErrorType`
```go
package main

import (
    "fmt"

    "github.com/jharriman/errant"
)

const PackageSpecificError ErrorType = "package_specific_error"

func main() {
	underlyingErr := errant.New("underlying error")
	err := errant.NewTypedError(underlyingErr, PackageSpecificError, "a special type of error")

	fmt.Println(errant.Is(err, PackageSpecificError))

	fmt.Println(errant.Unwrap(err).Error())
	fmt.Println(errant.Unwrap(err).(*errant.Err).ErrorDescription)
	// Output:
	// true
	// untyped
	// underlying error
}
```

### HTTP errors
```go
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"

    "github.com/jharriman/errant"
)

func main() {
	underlyingError := errant.New("not found in database")
	err := errant.NewNotFound(underlyingError, "item not found")

	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	err.WriteHTTP(recorder, request)

	fmt.Println(recorder.Body.String())
	// Output:
	// {"error":"not_found","error_description":"item not found","error_cause":{"error":"untyped","error_description":"not found in database"}}
}
```