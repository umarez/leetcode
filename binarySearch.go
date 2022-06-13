package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for {
		if !(l <= r) {
			break
		}

		m := l + (r-l)/2

		// check if target is in mid
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
