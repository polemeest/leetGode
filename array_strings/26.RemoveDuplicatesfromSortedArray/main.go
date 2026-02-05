package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	current_num := nums[0]
	last_index := 1
	for i, v := range nums {
		if i >= last_index && current_num < v {
			nums[last_index] = v
			current_num = v
			last_index++
		}
	}
	return last_index
}

func main() {
	nums := []int{1, 2}
	fmt.Println(removeDuplicates(nums))
}
