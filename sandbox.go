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

func hourglassSum(arr [][]int32) int32 {
	var hourglasses []int32
	for i := 0; i < len(arr); i++ { // column array
		for j := 0; j < len(arr[i]); j++ { // row array
			if i < len(arr)-2 && j < len(arr[i])-2 { // row array limitation
				var sum int32 = 0           // hourglass sum result
				for k := j; k <= j+2; k++ { // hourglass first row loop
					// fmt.Println(arr[i][k], "row1")
					sum += arr[i][k]
				}
				// fmt.Println(arr[i+1][j+1], "middle")
				sum += arr[i+1][j+1]        // hourglass middle
				for l := j; l <= j+2; l++ { // hourglass third row loop
					// fmt.Println(arr[i+2][l], "row2")
					sum += arr[i+2][l]
				}

				hourglasses = append(hourglasses, sum)
			}
		}

	}
	// fmt.Println(hourglasses)
	var result int32 = 0
	for i := 0; i < len(hourglasses); i++ {
		if hourglasses[i] > result {
			result = hourglasses[i]
		}
	}

	return result
}

func sandbox() {
	// palindromeChecker("kodok")
	// isPalindrome := palindromeChecker("adba", 0)
	// fmt.Println(isPalindrome)

	// fmt.Println(sockMerchant(20, []int32{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5}))

	input := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	input2 := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	fmt.Println(hourglassSum(input))
	fmt.Println(hourglassSum(input2))
}
