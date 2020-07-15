package tree


func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	if maxDepth(root) == -1{
		return false
	}
	return true
}

//想法挺棒，值得细品
func maxDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left - right > 1 || right - left > 1{
		return -1
	}

	return max(maxDepth(root.Left)+1,maxDepth(root.Right)+1)

}



//超级笨方法
func isBalanced1(root *TreeNode) bool {
	if root == nil{
		return true
	}

	l := isBalanced1(root.Left)
	r := isBalanced1(root.Right)
	sub := maxDepth(root.Left)-maxDepth(root.Right)
	if l && r && sub<=1 && sub >= -1{
		return true
	}

	return false
}

func maxDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}

	return max(maxDepth(root.Left)+1,maxDepth(root.Right)+1)

}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}