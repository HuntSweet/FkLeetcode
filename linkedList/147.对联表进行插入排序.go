package linkedList


//还有时间更少的方法
func insertionSortList2(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	right := insertionSortList2(head.Next)
	dummy := &ListNode{}
	dummy.Next = right
	pre := dummy
	cur := right

	flag := 0
	for cur != nil{
		if cur.Val >= head.Val{
			head.Next = cur
			pre.Next = head

			flag = 1
			break
		}
		cur = cur.Next
		pre = pre.Next
	}

	if flag == 0{
		pre.Next = head
		head.Next = nil
	}
	return dummy.Next
}
