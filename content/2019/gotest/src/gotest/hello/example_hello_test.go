package hello_test

import (
	"fmt"
	"gotest/hello"
)

func ExampleHello_extpkg() {
	fmt.Println(hello.Hello("Gopher"))

	// Output: Hello, Gopher!
}
