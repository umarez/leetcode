package main

func main() {
	// tc := []int{1, 2, 3, 1}
	// println(containsDuplicate(tc))
	// println(anagram("rat", "car"))
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	// println(bestTimeBuyStock([]int{7, 6, 4, 3, 1}))
	// println(validParentheses("["))
	t := ListNode{
		Val:  2,
		Next: nil,
	}

	t.addNode(3)
	t.addNode(5)
	t.reverseList().printNode()
}
