package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


/**
 * 最近公共祖先所在层的节点数值和​
 * @param root TreeNode类 根节点
 * @param p int整型 节点值
 * @param q int整型 节点值
 * @return int整型
 */
func findAndSum( root *TreeNode ,  p int ,  q int ) int {
	// write code here
	//进行层次遍历
	if root == nil{
		return 0
	}
	stack := []*TreeNode{root}
	node := lowest(root,p,q)
	count := 0
	for len(stack) != 0{
		l := len(stack)
		for _,v := range stack{
			if v == node{
				for _,i := range stack[:l]{
					count+=i.Val
					return count
				}
			}
			if v.Left != nil{
				stack = append(stack,v.Left)
			}
			if v.Right != nil{
				stack = append(stack,v.Right)
			}

		}
		stack = stack[l:]
	}
	return 0
}

func lowest(root *TreeNode,p,q int) *TreeNode {
	if root == nil{
		return nil
	}
	if p == root.Val || q == root.Val{
		return root
	}

	left := lowest(root.Left,p,q)
	right := lowest(root.Right,p,q)

	if left != nil && right != nil{
		return root
	}

	if left == nil && right != nil{
		return right
	}

	if left != nil && right == nil{
		return left
	}

	return nil
}