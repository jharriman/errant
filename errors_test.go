package errant

import "fmt"

func ExampleWrap() {
	err := New("base error")
	err2 := Wrap(err, "level 1 error")
	err3 := Wrap(err2, "level 2 error")

	fmt.Println(Unwrap(err3).Error())
	// Output: level 1 error: base error
}
