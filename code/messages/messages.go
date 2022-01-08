package messages

import "fmt"

func Greet(name string, id int) string {
	return fmt.Sprintf("Hello, %v\n your ID is %v\n", name, id)
}

func Depart(name string) string {
	return fmt.Sprintf("Bye, %v\n", name)
}
