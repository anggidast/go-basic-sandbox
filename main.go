// ? RECEIVE USER INPUT

package main

import (
	"bufio" // * package for input output
	"fmt"
	"os" // * operating system package
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) { // * reusable to show prompt and get input
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // * NewReader method need the source of information, terminal = standard input (Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n') // * name is variable to receive user input after press enter, _ is actually an error
	// name = strings.TrimSpace(name)     // * remove any white space

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
