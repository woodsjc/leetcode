package main

import (
	"fmt"
)

func oneAway(s1 string, s2 string) bool {
	// move through string and compare to other
	// whenever hit a difference up counter and move on
	// if second counter update return false
	totalChanges := 0
	s2Offset := 0
	lenDiff := len(s2) - len(s1)
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if i >= len(s2) { //can only delete
			if totalChanges <= 1 {
				return true
			}
			return false
		}

		if s1[i] != s2[i+s2Offset] {
			totalChanges++
			if totalChanges > 1 {
				return false
			}

			if len(s1) < len(s2) { //can only insert
				s2Offset++
				continue
			} else if len(s1) == len(s2) { //can only replace
				continue
			} else if len(s1) > len(s2) { //can only delete
				s2Offset--
				continue
			}
		}
	}

	return true
}

func main() {
	type strPair struct {
		s1 string
		s2 string
	}

	pairs := []strPair{
		{"abc", "abcd"},
		{"abcowiefj", "oqiwueabcd"},
		{"abcxdx", "abcdxdxd"},
		{"abc", "abc"},
		{"abca", "abcd"},
		{"xydef", "ydef"},
		{"xydef", "yydef"},
		{"xydef", "yxdef"},
	}

	for _, pair := range pairs {
		if oneAway(pair.s1, pair.s2) {
			fmt.Printf("%v pair is one away.\n", pair)
		} else {
			fmt.Printf("%v pair is not one away.\n", pair)
		}
	}
}
