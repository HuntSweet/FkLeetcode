package tree

func levelOrder(root *TreeNode) [][]int {

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
		res = append(res,re)
		stack = stack[l:]
	}

	return res
}
