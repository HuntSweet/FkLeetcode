package linkedList

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
