// ? RECEIVE USER INPUT

package main

import (
	"bufio" // * package for input output
	"fmt"
	"os" // * operating system package
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // * NewReader method need the source of information, terminal = standard input (Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n') // * name is variable to receive user input after press enter, _ is actually an error
	name = strings.TrimSpace(name)     // * remove any white space

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	myBill := createBill()

	fmt.Println(myBill)
}
