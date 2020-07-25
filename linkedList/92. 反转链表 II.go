package linkedList

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	pre.Next = head
	cur := head

	i := 0
	for i < m-1 && cur != nil{
		next := cur.Next

		pre = cur
		cur = next
		i++
	}

	j := 0
	//要反转的次数等于位置之差
	for cur != nil && j < n-m{
		next := cur.Next

		//判断是否为空，以防止n过大导致 nil
		if next != nil{
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}else{
			break
		}

		j++
	}

	return dummy.Next
}
