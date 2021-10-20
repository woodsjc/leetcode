package main

import (
	"fmt"
)

func main() {
	stringList := []string{
		"aaaaaabbbbbbbccccccc",
		"woiejfowiqpweuipqweip",
		"zzzbccccbbb",
		"oqwieoqiwueoiquweouqweoiwqeu",
		"thisisit",
	}

	for _, s := range stringList {
		fmt.Printf("%s: compressed-%s\n", s, stringCompression(s))
	}
}

func stringCompression(str string) string {
	compressed := ""

	for i := 0; i < len(str); i++ {
		count := 1
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				count++
				i++
			} else {
				break
			}
		}
		compressed += fmt.Sprintf("%c%d", str[i], count)
	}

	if len(compressed) < len(str) {
		return compressed
	}
	return str
}
