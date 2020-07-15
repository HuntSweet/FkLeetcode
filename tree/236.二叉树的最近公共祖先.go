package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if p == root || q == root{
		return root
	}

	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)

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
