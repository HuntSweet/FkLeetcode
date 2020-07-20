package linkedList

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	n := head.Next
	p := head
	for n != nil{
		if n.Val == p.Val{
			p.Next = n.Next
		}else{
			p = n
		}
		n = n.Next
	}

	return head

}

