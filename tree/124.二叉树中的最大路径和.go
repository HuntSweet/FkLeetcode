package tree

func maxPathSum(root *TreeNode) int {

	var res = -1<<32
	var f func(node *TreeNode) int
	f = func (root *TreeNode) int{
		if root == nil{
			return 0
		}

		left := max(0,f(root.Left))
		right := max(0,f(root.Right))

		res = max(root.Val+right+left,res)

		return max(left,right) + root.Val
	}
	f(root)
	return res

}


func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}
