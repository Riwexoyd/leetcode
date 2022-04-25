package main

import "fmt"

func main() {
	var nums = []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	k := 3

	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	swapArray(nums, 0, len(nums)-k-1)
	swapArray(nums, len(nums)-k, len(nums)-1)
	swapArray(nums, 0, len(nums)-1)
}

func swapArray(nums []int, start int, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start = start + 1
		end = end - 1
	}
}
