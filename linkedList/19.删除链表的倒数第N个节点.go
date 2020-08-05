package linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//思路:使用双指针，因为有可能删除头节点，所以设置哑巴节点
	cur := head
	dummy := &ListNode{}
	dummy.Next = head
	firstNode := dummy
	i := 0

	for cur != nil{
		if i >= n {
			firstNode = firstNode.Next
		}
		if cur.Next == nil{
			firstNode.Next = firstNode.Next.Next
			break
		}
		cur = cur.Next
		i++
	}

	return dummy.Next

}