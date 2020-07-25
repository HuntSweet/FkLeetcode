package linkedList

func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{}
	pre := dummy
	pre.Next = head


	for head != nil && head.Next != nil{

		firstNode := head
		secondNode := head.Next

		pre.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		pre = firstNode
		head = firstNode.Next

	}

	return dummy.Next
}
