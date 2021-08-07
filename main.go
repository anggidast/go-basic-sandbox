// ? CUSTOM DATA TYPE USING STRUCT
// * struct is another data type in Go, basically is a blueprint which describe a type of data
// * kind of custom object structure (like class in OOP)

package main

import "fmt"

func main() {
	myBill := newBill("mario's bill") // * use the newBill function from bill.go

	myBill.format() // * use bill format function
	// fmt.Println(myBill) // * result: {mario's bill map[] 0}

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	myBill.updateTip(10)

	fmt.Println(myBill.format())
}
