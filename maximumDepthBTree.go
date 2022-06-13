package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// find max depth value between left root and right root
	max := 0
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	if a > b {
		max = a
	} else {
		max = b
	}

	return 1 + max

}
