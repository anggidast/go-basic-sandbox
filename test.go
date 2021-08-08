package main

import "fmt"

func anyType(a interface{}) interface{} {
	return a
}

func main() {
	menu := map[interface{}]interface{}{
		"soup":           4.99,
		4:                "string",
		1.3:              "6.99",
		"coffee pudding": true,
	}

	anyArray := [4]interface{}{"string", 5, 3.7, true}

	anySlice := []interface{}{"string", 5, 3.7, true}

	fmt.Println(anyType("string"))

	fmt.Println(menu)
	fmt.Println(anySlice)
	fmt.Println(anyArray)
}
