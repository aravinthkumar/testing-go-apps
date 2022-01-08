// The name should be test to be identified by the test constracts as test package
// This will not be bundled with the binary
package main_test

import (
	"testing"
)

//  The function should start with Test to make the test constracts to know this is the test method
func TestAddition(t *testing.T) {
	a := 1
	b := 2
	actual := a + b
	expected := 3

	if actual != expected {
		t.Errorf("Value a '%v' and value b '%v' is not summing ", a, b)

	}

}

func TestSubtraction(t *testing.T) {
	a := 10
	b := 5
	actual := a - b
	expected := 3

	if actual != expected {
		t.Errorf("Value a '%v' and value b '%v' is not subtracting as expected ", a, b)

	}

}
