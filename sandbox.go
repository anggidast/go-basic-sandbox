package main

import (
	"fmt"
	"sort"
)

// func palindromeChecker(n string, i int) bool {
// var temp string = ""

// for i := len(n); i > 0; i-- {
// 	temp += string(n[i-1])
// }
// return temp == n

// for i := 0; i < len(n)/2; i++ {
// 	fmt.Println(i)
// 	if string(n[i]) != string(n[len(n)-(i+1)]) {
// 		return false
// 	}
// }
// return true

// 	if i < len(n)/2 {
// 		if string(n[i]) != string(n[len(n)-(i+1)]) {
// 			return false
// 		}
// 		return palindromeChecker(n, i+1)
// 	}
// 	return true
// }

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var result int32 = 0
	int32AsIntValues := make([]int, len(ar))

	for i, val := range ar {
		int32AsIntValues[i] = int(val)
	}
	sort.Ints(int32AsIntValues)

	for i, val := range int32AsIntValues {
		ar[i] = int32(val)
	}
	fmt.Println(ar)
	var temp int32 = ar[0]
	for i := 0; i < len(ar)-1; i += 2 {
		fmt.Println(temp, ar[i+1])
		if ar[i+1] == temp {
			result++
		}
		if i < len(ar)-2 {
			temp = ar[i+2]
		}
	}

	return result
}

func main() {
	// palindromeChecker("kodok")
	// isPalindrome := palindromeChecker("adba", 0)
	// fmt.Println(isPalindrome)

	fmt.Println(sockMerchant(20, []int32{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5}))
}
