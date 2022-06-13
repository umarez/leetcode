package main

func containsDuplicate(nums []int) bool {
	var m = map[int]int{}
	for _, n := range nums {
		m[n] = m[n] + 1
		if m[n] > 1 {
			return true
		}
	}
	return false
}
