package main

import "fmt"

func main() {
	var nums = []int{
		-4, -1, 0, 3, 10,
	}

	var result = sortedSquares(nums)
	fmt.Println(result)
}

func sortedSquares(nums []int) []int {

	result := make([]int, len(nums))

	left := 0
	right := len(nums) - 1
	index := len(result) - 1

	for index >= 0 {
		leftValue := nums[left] * nums[left]
		rightValue := nums[right] * nums[right]

		if leftValue > rightValue {
			result[index] = leftValue
			left = left + 1
		} else {
			result[index] = rightValue
			right = right - 1
		}

		index = index - 1
	}

	return result
}
