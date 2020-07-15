package tree

//

//时间复杂度较高
func isValidBST2(root *TreeNode) bool {
	if root == nil{
		return true
	}

	var f func(root *TreeNode)
	min := -1<<32
	t := true
	f = func(root *TreeNode){
		if root == nil{
			return
		}

		f(root.Left)
		if min >= root.Val{
			t = false
			return
		}
		min = root.Val
		f(root.Right)
	}

	f(root)

	return t
}
