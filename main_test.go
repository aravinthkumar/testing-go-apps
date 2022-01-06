package main_test

import (
	"testing"
)

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
