package main

import (
	"fmt"
	"sort"
)

// 1
func firstRepeatedWord(sentence string) string {
	// Write your code here
	wordsArr := []string{}
	word := ""
	for i := 0; i < len(sentence); i++ {
		// fmt.Println(string(sentence[i]))
		if string(sentence[i]) == " " || string(sentence[i]) == "," || string(sentence[i]) == ":" || string(sentence[i]) == ";" || string(sentence[i]) == "-" || string(sentence[i]) == "." {
			if word != "" {
				wordsArr = append(wordsArr, word)
			}
			word = ""
		} else {
			word += string(sentence[i])
		}
		if i == len(sentence)-1 {
			if word != "" {
				wordsArr = append(wordsArr, word)
			}
			word = ""
		}
	}

	smallest := len(wordsArr)
	result := ""
	for j := 0; j <= len(wordsArr); j++ {
		for k := j + 1; k < len(wordsArr); k++ {
			if wordsArr[j] == wordsArr[k] {
				if k-j < smallest {
					smallest = k - j
					result = wordsArr[j]
				}
			}
		}
	}
	return result
}

// 2
func smallestString(substrings []string) string {
	// Write your code here
	result := ""

	// sort the array
	for i := 0; i < len(substrings); i++ {
		for j := i + 1; j < len(substrings); j++ {
			// fmt.Println(substrings[i] + substrings[j])
			// comparing which of the
			// two concatenation causes
			// lexiographically smaller
			// string
			if substrings[i]+substrings[j] > substrings[j]+substrings[i] {
				s := substrings[i]
				substrings[i] = substrings[j]
				substrings[j] = s
			}
		}
	}

	for i := 0; i < len(substrings); i++ {
		result += substrings[i]
	}

	return result
}

// 5
func largestArea(H, V []int, N, M, h, v int) int {
	// Stores all bars
	s1 := []int{}
	s2 := []int{}

	// Insert horizontal bars
	for i := 1; i <= N+1; i++ {
		s1 = append(s1, i)
	}

	// Insert vertictal bars
	for i := 1; i <= M+1; i++ {
		s2 = append(s2, i)
	}

	// Remove horizontal separators from s1
	filteredS1 := []int{}
	for i := 0; i < h; i++ {
		for j := range s1 {
			if s1[j] != H[i] {
				filteredS1 = append(filteredS1, s1[j])
			}
		}
		// s1 = s1.filter((el) => el != H[i]);
	}
	s1 = filteredS1

	// Remove vertical separators from s2
	filteredS2 := []int{}
	for i := 0; i < v; i++ {
		for j := range s2 {
			if s2[j] != H[i] {
				filteredS2 = append(filteredS2, s2[j])
			}
		}
		// s2 = s2.filter((el) => el != V[i]);
	}
	s2 = filteredS2

	// Stores left out horizontal and
	// vertical separators
	list1 := []int{}
	list2 := []int{}

	j := 0
	for i := 0; i < len(s1); i++ {
		j++
		// list1[j] = s1[i]
		list1 = append(list1, s1[i])
	}
	j = 0
	for i := 0; i < len(s2); i++ {
		j++
		// list2[j] = s2[i]
		list2 = append(list2, s2[i])
	}

	// var i = 0;
	// s1.forEach((element) => {
	//   list1[i++] = element;
	// });
	// i = 0;
	// s2.forEach((element) => {
	//   list2[i++] = element;
	// });

	// Sort both list in
	// ascending order
	sort.Ints(list1)
	sort.Ints(list2)
	// list1.sort((a, b) => a - b);
	// list2.sort((a, b) => a - b);

	maxH := 0
	p1 := 0
	maxV := 0
	p2 := 0

	// Find maximum difference of neighbors of list1
	for j := 0; j < len(s1); j++ {
		if maxH < list1[j]-p1 {
			maxH = list1[j] - p1
		}
		// maxH = max(maxH, list1[j] - p1);
		p1 = list1[j]
	}

	// Find max difference of neighbors of list2
	for j := 0; j < len(s2); j++ {
		if maxV < list2[j]-p2 {
			maxV = list2[j] - p2
		}
		// maxV = Math.max(maxV, list2[j] - p2);
		p2 = list2[j]
	}

	// Print largest volume
	return (maxV * maxH)
}

func main() {
	// fmt.Println(firstRepeatedWord("We work			   hard because hard work pays"))

	// fmt.Println(smallestString([]string{"c", "cc", "cca", "cccb"}))
	N := 3
	M := 2

	// Given arrays
	H := []int{1, 2, 3}
	V := []int{1, 2}
	h := len(H)
	v := len(V)

	fmt.Println(largestArea(H, V, N, M, h, v))
}
