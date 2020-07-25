package linkedList

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for head != nil{
		if head.Val == val{
			pre.Next = head.Next
		}else{
			pre = pre.Next
		}
		head = head.Next
	}

	return dummy.Next
}
