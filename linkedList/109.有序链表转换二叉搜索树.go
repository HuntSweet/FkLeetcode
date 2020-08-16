package linkedList
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func sortedListToBST(head *ListNode) *TreeNode {
	//思路;找到中点，递归生成平衡树
	if head == nil{
		return nil
	}
	var pre *ListNode
	cur,next := head,head

	for next != nil && next.Next != nil { //end or end.Next
		pre = cur
		cur = cur.Next      //move 1
		next = next.Next.Next //move 2
	}

	if pre != nil {
		pre.Next = nil
	}


	node := cur.Next

	root := new(TreeNode)
	root.Val = cur.Val
	//如果不返回，则会无限循环
	if head == cur{
		return root
	}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(node)

	return root
}
