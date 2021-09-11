package main // * every project should have at least one file with package main (file with package name will run first)

import (
	"fmt" // * fmt is use for format string and print them out tod the console
	"strings"
)

// * Basic Functions
// func sayGreeting(n string) { // * n param only receive string data arg
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) { // * multiple params and function as argument
// 	for _, v := range n { // * for every n slices elements (v), invoke the function with v as argument
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 { // * if the func return data, declare the data type before {
// 	return math.Pi * r * r
// }

// * Multiple Return Values
func getInitials(n string) (string, string) { // * the func have 2 params with string data type both
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") // * split the string with space as separator

	initials := []string{}
	for _, v := range names {
		initials = append(initials, v[:1]) // * push first character at every names elements (v), v[0] cannot use at string var
	}

	if len(initials) > 1 {
		return initials[0], initials[1] // * return 2 value
	}

	return initials[0], "" // * return 1 value and 1 empty string value
}

// * Variadic Function
// * is a function with dynamic parameters amount (works like spread operator/rest parameter in JS)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// * Function as Value and as Parameter
func getGoodBye(name string) string {
	return "Good bye " + name
}

type Filter func(string) string // * bisa langsung bisa dengan type declaration

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter (name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

// * Anonymous Function
type Blacklist func(string) bool
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// * Recursive Function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1) // recursion is here
	}
}

// * Defer Function
// * the function is schedule to invoke/execute after other function executed (even the other function is error) always executed
func logging() {
	fmt.Println("executing function done")
}

func runApplication() {
	defer logging() // * put defer function at the first of other function
	fmt.Println("run application")
}

// * Panic function
// * function that use to stop program when error appear
// * defer function will always be executed even use panic function

// * Recover function
// * function that use to catch panic data
// * when use recover, panic process will be stop and the program/function will run continue

func endFunc() {
	fmt.Println("function end")
	message := recover()
	if message != nil {
		fmt.Println("error happen", message)
	}
}

func runFunc(error bool) {
	defer endFunc()
	if error {
		panic("error")
	}
	fmt.Println("function run")
}

// * func main is entry point for the application, only one in one file, will run first in file
func main() {
	// * Basic Functions
	// 	sayGreeting("mario")
	// 	sayGreeting("luigi")
	// 	sayBye("mario")

	// 	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	// 	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	// 	a1 := circleArea(10.5)
	// 	a2 := circleArea(15)

	// 	fmt.Println(a1, a2)
	// 	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	// * Multiple Return Values
	// fn1, sn1 := getInitials("tifa lockhart") // * multi variabel store/declaration when it get multiple return
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("cloud strife")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("barret")
	// fmt.Println(fn3, sn3)

	// * Variadic Function
	// fmt.Println(sumAll(1,2,3,4,5))
	
	// * Send slice to variadic function
	// slice := []int{1,2,3,4,5}
	// fmt.Println(sumAll(slice...))

	// * Function as value/variable
	// goodBye := getGoodBye
	// fmt.Println(goodBye("anggi"))

	// * Function as parameter
	// sayHelloWithFilter("Anggi", spamFilter)
	// sayHelloWithFilter("Anjing", spamFilter)

	// * Anonymous Function
	// blacklist := func(name string) bool {
	// 	return name == "admin"
	// }

	// registerUser("admin", blacklist)
	// registerUser("admin", func(name string) bool {
	// 	return name == "root"
	// })

	// * Recursive
	// fmt.Println(factorialRecursive(6))

	// * Panic and Recover
	runFunc(false)
	fmt.Println("will execute even runFunc error/panic")

}
