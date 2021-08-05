package main // * every project should have at least one file with package main (file with package main will run first)

import "fmt" // * fmt is use for format string and print them out to the console

// * func main is entry point for the application, only one in a project, will run first in file
func main() {

	// * print to terminal
	// fmt.Println("Let's Go!")

	// ? VARIABLES
	// * STRINGS
	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// * variable data type cannot be change
	// nameOne = 21 // * -> error
	// nameOne = "peach"

	// fmt.Println(nameOne)

	// * variable declaration shorthand (cannot use outside of the func)
	// nameFour := "yoshi" // * is equal to var nameFour = "yoshi"

	// fmt.Println(nameFour)

	// * INTEGERS
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// * bits & memory
	// * define how much bits/memory the variable can use on range
	// var numOne int8 = 25 // * int8 -> -128 ~ 127
	// var numTwo int8 = -128
	// * unsigned, the value cannot be minus
	// var numThree uint8 = 25 // * uint8 -> 0 ~ 255

	// * FLOATS
	// * the bits must be defined
	// var scoreOne float32 = -1.5
	// var scoreTwo float64 = 9418490289080092084.8
	// scoreThree := 1.5 // * -> default: float64

	// ? fmt PACKAGE

	// age := 35
	// name := "shaun"

	// * Print, not add a new line
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("new line \n")

	// * Println auto add a new line after
	// fmt.Println("hello ninjas!")
	// fmt.Println("goodbye ninjas!")
	// fmt.Println("my age is", age, "and my name is", name)

	// * Printf (formatted strings), add varieble that will use to print at the second argument
	// * at the first argument add %_ as format specifier
	// fmt.Printf("my name is %v and my name is %v \n", age, name)
	// fmt.Printf("my name is %q and my name is %q", age, name) // * %q add a quote around the variable when output, number will outputed as '#'
	// fmt.Printf("age is of type %T \n", age) // * %T outputed the variabel type
	// fmt.Printf("you scored %f points \n", 225.55) // * outputed float variable
	// fmt.Printf("you scored %0.2f points\n", 225.55) // * outputed float variable with 2 decimal points

	// * Sprintf (save formatted strings)
	// var str = fmt.Sprintf("my name is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string is:", str)

	// ? ARRAYS AND SLICES
	// * Arrays
	// var ages [3]int = [3]int{20, 25, 30} // * once an array been declared, the length cannot be change
	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"yoshi", "mario", "peach", "bowser"}
	// names[1] = "luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// * Slices (it use arrays under the hood), it more like array in other language
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85) // * overwrite/push a value to array

	// fmt.Println(scores, len(scores))

	// * Slice ranges, the output is slice
	// rangeOne := names[1:3]  // * [from:to(before)]
	// rangeTwo := names[2:]   // * from index 2, to the last index
	// rangeThree := names[:3] // * from first index to the before index 3

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)

	// ? STANDARD LIBRARY
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello")) // * return boolean
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // * return new string after replace without change the original value
	// fmt.Println(strings.ToUpper(greeting))     // * return new string
	// fmt.Println(strings.Index(greeting, "a")) // * return the position (number), if the value doesn't exist it'll return -1
	// fmt.Println(strings.Split(greeting, " ")) // * return array

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages) // * sort the original slices value
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30) // * return the position (after sorted), if the value doesn't exist it'll return the slices last index + 1
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))

	// ? LOOPS
	// x := 0
	// for x < 5 { // * while loop like
	// 	fmt.Println("value of x is:", x)
	// 	x++ // * infinite loop without this
	// }

	// for i := 0; i < 5; i++ { // * normal loop
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ { // * normal loop by slices length
	// 	fmt.Println(names[i])
	// }

	// for index, val := range names { // * for range, for.. in like loop
	// 	fmt.Printf("the value at pos %v is %v \n", index, val)
	// 	val = "new string"
	// }

	// for _, val := range names { // * for range without index, _ is used to unused variable
	// 	fmt.Print(val, ",")
	// 	val = "new string"
	// }

	// fmt.Println(names) // * changing val in a loop does not mutate the original slice

	// ? BOOLEANS AND CONDITIONS
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, val := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // * skip next line and continue the loop
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break // * stop the loop
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}
}
