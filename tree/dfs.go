package tree

//从上往下
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil{
		return res
	}

	dfs(root,&res)
	return res
}

func dfs(node *TreeNode,res *[]int)  {
	if node == nil{
		return
	}

	*res = append(*res,node.Val)
	dfs(node.Left,res)
	dfs(node.Right,res)
}

//从下往上（分治)
func divideAndConquer(root *TreeNode) (res []int) {
	if root == nil{
		return res
	}

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	res = append(res,root.Val)
	res = append(res,left...)
	res = append(res,right...)

	return
}

//104.反向层次遍历
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) != 0{
		re := []int{}
		l := len(stack)
		for i:=0;i<l;i++{
			re = append(re,stack[i].Val)
			if stack[i].Left != nil{
				stack = append(stack,stack[i].Left)
			}
			if stack[i].Right != nil{
				stack = append(stack,stack[i].Right)
			}
		}
		stack = stack[l:]
		res = append(res,re)
	}

	//记住反转的话，一定要小于1/2长度，因为超过了之后，又反转回去了
	for i:=0;i<len(res)/2;i++{
		res[i],res[len(res)-1-i] = res[len(res)-1-i],res[i]
	}
	return res
}
