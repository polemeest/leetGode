package main

import "fmt"

func maxProfit(prices []int) int {
	current_max := 0
	global_max := 0
	for i := 1; i < len(prices); i++ {
		current_max = max(current_max+(prices[i]-prices[i-1]), 0)
		global_max = max(global_max, current_max)
	}
	return global_max
}

func maxProfitPointers(prices []int) int {
	l, r, global_max := 0, 1, 0
	for r < len(prices) {
		if prices[l] > prices[r] {
			l = r
		} else {
			global_max = max(global_max, prices[r]-prices[l])
		}
		r++
	}
	return global_max
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
	fmt.Println(maxProfitPointers(nums))
}
