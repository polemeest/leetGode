package main

import (
	"fmt"
	"strings"
	"unicode"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	newStr := string(s[0])
	for i := 1; i < len(s); i++ {
		if !unicode.IsSpace(rune(s[i])) {
			if unicode.IsSpace(rune(s[i-1])) {
				newStr = string(s[i])
			} else {
				newStr += string(s[i])
			}
		}
	}
	return len(newStr)
}

func lengthOfLastWordShort(s string) int {
	array := strings.Fields(s)
	return len(array[(len(array) - 1)])
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordShort(s))
}
