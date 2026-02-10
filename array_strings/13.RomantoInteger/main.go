package main

import "fmt"

var dictionary = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var dictionary2 = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	memory := ""
	for _, v := range s {
		current := memory + string(v)
		if i := dictionary[current]; i == 0 {
			res += dictionary[memory]
			memory = string(v)
		} else {
			memory = current
		}
	}
	if len(memory) > 0 {
		res += dictionary[memory]
	}
	return res
}

func romanToInt2(s string) (res int) {
	end := len(s) - 1
	for i := range s {
		if i < end && dictionary2[s[i]] < dictionary2[s[i+1]] {
			res -= dictionary2[s[i]]
		} else {
			res += dictionary2[s[i]]
		}
	}
	return
}

func main() {
	str := "LVIII"
	fmt.Println(romanToInt(str))
}
