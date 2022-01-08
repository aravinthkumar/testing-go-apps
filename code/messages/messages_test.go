// messages is used since we are using whitebox testing ,
//all the menbers of the messages package can be used
// message_test used then member can be access only by importing the package.
package messages

import "testing"

// Test is prefixed for Go Test runner to identity this function as test
func TestGreet(T *testing.T) {
	actual := Greet("Gophers")
	expected := "Hello, Gophers"
	if actual != expected {
		T.Errorf("Greet didn't matched, actual: '%v' but expected: '%v' ", actual, expected)
	}
}

func TestDepart(T *testing.T) {
	actual := Depart("Gophers")
	expected := "Bye, Gophers"
	if actual != expected {
		T.Errorf("Greet didn't matched, actual: '%v' but expected: '%v' ", actual, expected)
	}
}
