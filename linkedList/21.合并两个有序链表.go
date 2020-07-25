package linkedList

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	nl := &ListNode{}
	pre := nl
	for l1 != nil && l2 != nil{
		if l1.Val >= l2.Val{
			nl.Next = l2
			l2 = l2.Next
		}else{
			nl.Next = l1
			l1 = l1.Next
		}
		nl = nl.Next
	}

	if l1 != nil{
		nl.Next = l1
	}
	if l2 != nil{
		nl.Next = l2
	}

	return pre.Next
}
