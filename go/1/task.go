package main

import "fmt"

func main() {
	var nums = []int{
		3, 3,
	}

	k := 6

	var result = twoSum(nums, k)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var result [2]int
	sums := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		temp := target - nums[i]

		if val, ok := sums[temp]; ok {
			result[0] = val
			result[1] = i
			break
		} else {
			sums[nums[i]] = i
		}
	}

	return result[:]
}
