package main // * every project should have at least one file with package main (file with package main will run first)

import "fmt"

var score = 99.5 // * root level (package scope) to so it accessible inside any file with package name

// * func main is entry point for the application, only one in a project, will run first in file
func main() {

	// * all var and func inside this func main is coming from greetings.go

	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

	// * variable inside this function is unaccessible at other file

}

// * to run the package (form different file), we also have to run all the file that connect to main.go -> go run main.go greetings.go
