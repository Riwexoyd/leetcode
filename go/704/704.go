package main

import "fmt"

func main() {
	var nums = []int{
		-1, 0, 3, 5, 9, 12,
	}
	var target int

	target = 9

	var result = search(nums, target)
	fmt.Println(result)
}

func search(nums []int, target int) int {
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

	return -1
}
