package main

import "fmt"

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
