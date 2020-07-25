package linkedList

//移动链表在组合
func partition(head *ListNode, x int) *ListNode {
	left,right := &ListNode{},&ListNode{}

	pre,next := left,right
	for head != nil{
		if head.Val >= x{
			right.Next = head
			right = right.Next
		}else{
			left.Next = head
			left = left.Next
		}
		head = head.Next
	}

	//防止成环
	right.Next = nil

	left.Next = next.Next

	return pre.Next



}
