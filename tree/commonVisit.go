package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//通用非递归遍历二叉树
func PreorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil{
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) != 0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil{
			//根据哪种遍历，来确定入栈顺序
			//右
			if node.Right != nil{
				stack = append(stack,node.Right)
			}
			//左
			if node.Left != nil{
				stack = append(stack,node.Left)
			}
			//根
			stack = append(stack,node)
			stack = append(stack,nil)
		}else{
			res = append(res,stack[len(stack)-1].Val)
			stack = stack[:len(stack)-1]
		}
	}

	return res
}
