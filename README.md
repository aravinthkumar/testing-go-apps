# Testing Go Applications
This repository is for learning how to test Go applications.

Several test commands which can be used,

- `go test` to run the test in the current directory.
- `go test <package>` e.g. `git test github.com/aravinthkumar/testing-go-apps/code/message` to run a specific package.
- `go test -v` for providing more verbose (for example information of passed tests).
- `go test -run Greet` for running a specific test
- `go test -cover` to see the test coverage.
- `go test -coverprofile cover.out` prints out the coverage details into a `cover.out` file.
- `go tool cover -func cover.out` to open the coverage details in the terminal using go tool
- `go tool cover -html cover.out` to open the coverage details in the browser using go tool

## Special Type of Test : Table Driven Tests

- Here a single function is tested against all possible scenarios.

> following example shows where multiple scenarios are run against the function.

```go 
	scenarios := []struct {
		input  string
		output string
	}{
		{input: "Gophers", output: "Hello, Gophers"},
		{input: "", output: "Hello,   "},
	}
	for _, s := range scenarios {
		result := Greet(s.input)
		if result != s.output {
			T.Errorf(" Greet didn't match, actual: '%v' but expected: '%v'", result, s.output)
		}
	}
```


##  To Do 

- Learning Notes
- Benchmarking Test
- Profiling Test


  
