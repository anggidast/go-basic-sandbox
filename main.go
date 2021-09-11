// ? CUSTOM DATA TYPE USING STRUCT
// * struct is another data type in Go, basically is a blueprint which describe a type of data
// * kind of custom object structure (like class in OOP)

package main

import "fmt"

func main() {
	myBill := newBill("mario's bill") // * use the newBill function from bill.go

	// * other way to make struct data/object
	var b2 bill
	b2.name = "Anggi"
	b2.items = map[string]float64{
		"onion soup": 4.5,
		"veg pie": 8.95,
	}
	b2.tip = 10


	b3 := bill{
		"Dastarina",
		map[string]float64{
			"onion soup": 4.5,
			"veg pie": 8.95,
		},
		10,
	}

	fmt.Println(b2, b3)

	myBill.format() // * use bill format function
	// fmt.Println(myBill) // * result: {mario's bill map[] 0}

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	myBill.updateTip(10)

	fmt.Println(myBill.format())
}
