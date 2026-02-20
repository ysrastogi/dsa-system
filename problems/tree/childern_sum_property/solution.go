package childernsumproperty

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func childernSumProperty(root *TreeNode) {
	if root == nil{
		return 
	}
	child := 0
	if root.Left != nil{
		child += root.Left.Val
	}
	if root.Right != nil{
		child += root.Right.Val
	}
	if child >= root.Val{
		root.Val = child
	}else{
		if root.Left != nil{
			root.Left.Val = root.Val
		}
		if root.Right != nil{
			root.Right.Val = root.Val
		}
	}
	childernSumProperty(root.Left)
	childernSumProperty(root.Right)
	total := 0
	if root.Left != nil{
		total += root.Left.Val
	}
	if root.Right != nil{
		total += root.Right.Val
	}
	if root.Left != nil || root.Right != nil{
		root.Val = total
	}
}
	
