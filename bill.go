package main

import (
	"fmt"
	"os"
)

type bill struct { // * declare a new struct, a kinda blueprint of any bill object
	name  string // * name properties is string, etc
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill { // * function with bill data type return value
	b := bill{ // * declare a new bill variable
		name:  name,                 // * initial value of name properties is name (param)
		items: map[string]float64{}, // * initial value of items properties is empty map
		tip:   0,                    // * initial value of tip is 0
	}

	return b
}

// make method/function that associate with struct using receiver function
// limiting the function only to some struct

// format bill
func (b *bill) format() string { // * this function is receive the bill struct and can access the struct using b variable (copied)
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // * ex: pie: ...5.99
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip) // * %-25v is special character means the output variable will have 25 characters long
	// * -25 is 25 chars long at the right side
	// * 25 is 25 chars long at the left side

	return fs
}

// to update struct use pointer
// * RULE OF THUMB, WHENEVER CALLING A METHOD/FUNCTION WHERE UPDATING THE VALUE, SHOUL PASSING A POINTER
// * STRUCT POINTERS ARE AUTOMATICALLY DEREFERENCE
// * USE POINTER AT EVERY RECEIVER FUNCTION CAN DECREASE MEMORY USAGE, BECAUSE GO WILL ALWAYS COPYING THE ARGUMENT AT EVERY FUNCTION CALL (INVOKE)

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip // * updating struct value doesn't need dereference as long as the receiver function receive the struct pointer
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format()) // * byte slices to save file

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) // * 0644 is permission
	if err != nil {
		panic(err) // * stop the program and show the error
	}

	fmt.Println("Bill was saved to file")
}
