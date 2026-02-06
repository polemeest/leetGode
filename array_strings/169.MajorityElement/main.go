package main

import (
	"fmt"
	"sort"
)

type hMap struct {
	key   int
	value int
}

func majorityElement(nums []int) int {
	hashMap := make(map[int]int)
	l := len(nums) / 2
	for _, v := range nums {
		_, ok := hashMap[v]
		if ok != true {
			hashMap[v] = 1
		} else {
			hashMap[v] += 1
		}
	}
	items := []hMap{}

	for k, v := range hashMap {
		items = append(items, hMap{k, v})
	}

	sort.SliceStable(items, func(i int, j int) bool {
		return items[i].value > items[j].value
	})

	for _, v := range items {
		if v.value > l {
			return v.key
		}
	}
	return items[0].value
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
