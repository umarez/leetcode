package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, n := range nums {
		diff := target - n
		if _, ok := m[diff]; ok {
			return []int{m[diff], index}
		}
		m[n] = index
	}
	return []int{}
}
