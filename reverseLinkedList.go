package main

func (head *ListNode) reverseList() *ListNode {
	var prev *ListNode
	currNode := head

	for {
		if currNode == nil {
			break
		}
		// ambil temp
		temp := currNode.Next
		// assign next nya yg sekarang dengan nilai setelahnya
		currNode.Next = prev
		// geser pointer kedua ke kanan
		prev = currNode
		// geser pointer utama ke kanan
		currNode = temp
	}
	return prev
}
