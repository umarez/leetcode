package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) addNode(val int) {
	currNode := head
	newNode := ListNode{
		Val:  val,
		Next: nil,
	}
	for {
		if currNode.Next == nil {
			break
		}
		currNode = currNode.Next
	}

	currNode.Next = &newNode
}

func (head *ListNode) printNode() {
	currNode := head

	for {
		fmt.Println(currNode.Val)
		if currNode.Next == nil {
			break
		}
		currNode = currNode.Next
	}

}
