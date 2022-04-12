package main

import "fmt"

func main() {
	var nums = []int{
		1, 3, 5, 6,
	}
	var target int

	target = 7

	var result = searchInsert(nums, target)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2

		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}

	return left
}
