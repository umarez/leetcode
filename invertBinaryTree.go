package main

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    // swap node children
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
   
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}