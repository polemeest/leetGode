package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	example := strs[0]
	last_index := 0
	for i := 1; i < len(strs); i++ {
		given_str := strs[i]
		if len(given_str) == 0 {
			return ""
		}
		for j := 0; j < len(example) && j < len(given_str); j++ {
			val := string(example[j])
			new_val := string(given_str[j])
			if val != new_val {
				last_index = j
				break
			}
			last_index = j + 1
		}
		example = example[:last_index]
	}
	return example
}

func main() {
	strs := []string{"abab", "aba", ""}
	fmt.Println(longestCommonPrefix(strs))
}
