package errant

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func Example_status() {
	underlyingError := New("not found in database")
	err := NewNotFound(underlyingError, "item not found")

	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	err.WriteHTTP(recorder, request)

	fmt.Println(recorder.Body.String())
	// Output:
	// {"error":"not_found","error_description":"item not found","error_cause":{"error":"untyped","error_description":"not found in database"}}
}
