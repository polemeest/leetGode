package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
