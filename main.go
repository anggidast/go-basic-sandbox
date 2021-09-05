package main // * every project should have at least one file with package main (file with package main will run first)

import "fmt"

// * func main is entry point for the application, only one in a project, will run first in file
func main() {

	// ? MAPS
	// * like object in javascript
	// * all the key must have the same type in a single map
	// * and all the value must have the same type in a single map

	// map_name := map[key-data-type]value-data-type
	menu := map[string]float64{ // * all key data type in menu map is string, and all value in menu map is float
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"coffee pudding": 3.55,
	}

	// * make a new empty map
	var book map[string]string = make(map[string]string)
	fmt.Println(book)

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// * looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// * integers as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	// * updating map item
	phonebook[267584967] = "bowser"
	fmt.Println(phonebook)

	phonebook[845775485] = "yoshi"
	fmt.Println(phonebook)

	// * add new map item
	phonebook[1234567] = "anggi"

	// * delete map item
	delete(phonebook, 267584967) // will delete properties 267584967 on phonebook map
	fmt.Println(phonebook)

}
