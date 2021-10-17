//Strategy
//count all letters into hash table
//check all counts even or at most 1 odd length
package main

import (
	"fmt"
	"unicode"
)

func isPalindromePermutation(str string) bool {
	table := map[rune]int{}
	totalOdd := 0

	for _, c := range str {
		if unicode.IsLetter(c) {
			table[c] += 1
		}
	}

	for _, value := range table {
		if value%2 == 0 {
			continue
		}

		if totalOdd > 0 {
			return false
		}
		totalOdd++
	}

	return true
}

func main() {
	strList := []string{
        "aazzxzzaa",
        "deffed",
        "goingniog",
		"lkasjdfoiwjfaoeifjowiejf",
		"voiuwefjoimcoiwjceoijpoqupioeqwepo",
		"aaaaaaaaaaaaaaaaaaaaabbccdd",
		"oiuoiuoiuoiuoiuoiuoiuoiuoi",
		"qweqweqweqweqewqweqewqweqweqweqweqweqwex",
        "aabbcc",
	}

	for _, str := range strList {
        if isPalindromePermutation(str) {
            fmt.Printf("%s is a palindrome permutation.\n", str)
        } else {
            fmt.Printf("%s is not a palindrome permutation.\n", str)
        }
	}
}
