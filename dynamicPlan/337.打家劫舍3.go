package dynamicPlan

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
 }

func rob(root *TreeNode) int {
	//思路：节点有两个状态：1.偷 偷的话，那么获得的最大金额由左右子节点不偷时的最大金额加上节点的金额组成
	//2.不偷  由左右子节点(偷时、不偷时)的最大金额加起来（ps：重要一环)
	//
	if root == nil{
		return 0
	}

	res := robInternal(root)

	return max(res[0],res[1])
}
func robInternal(root *TreeNode)[2]int{
	temp := [2]int{}
	if root == nil{
		return temp
	}

	left := robInternal(root.Left)
	right := robInternal(root.Right)

	//不偷
	temp[0] =  max(left[1],left[0]) + max(right[1],right[0])
	//偷
	temp[1] = left[0] + right[0] + root.Val
	return temp
}
func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}
