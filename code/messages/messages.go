package messages

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}

func Depart(name string) string {
	return fmt.Sprintf("Bye, %v", name)
}
