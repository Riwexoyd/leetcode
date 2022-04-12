package main

import "fmt"

func main() {
	var result = firstBadVersion(1)
	fmt.Println(result)
}

func isBadVersion(version int) bool {
	return version == 1
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	bad := -1
	for left <= right {
		middle := left + (right-left)/2

		isBad := isBadVersion(middle)

		if !isBad {
			left = middle + 1
		} else {
			right = middle - 1
			bad = middle
		}
	}

	return bad
}
