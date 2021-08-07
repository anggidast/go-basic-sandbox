// ? RECEIVE USER INPUT

package main

import (
	"bufio" // * package for input output
	"fmt"
	"os" // * operating system package
	"strconv"
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

	// * SWITCH STATEMENT
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64) // * parsing float
		if err != nil {                         // * if parsing error
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64) // * parsing float
		if err != nil {                       // * if parsing error
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip add - ", tip)
		promptOptions(b)

	case "s":

		fmt.Println("you choose to save the bill", b)
	default:
		fmt.Println("invalid option!")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
