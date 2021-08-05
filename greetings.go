package main // * because this file is part of main.go

import "fmt"

// * anything that declare inside this top/root level of file, it automatically made accessible at any other file with the package main (inside main function)
var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("you scored this many points:", score) // score is coming from main.go
}
