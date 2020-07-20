package linkedList

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	nu := &ListNode{}
	nu.Next = head
	head = nu
	var rm int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val{
			rm = head.Next.Val
			for head.Next != nil && head.Next.Val == rm{
				head.Next = head.Next.Next
			}
		}else{
			head = head.Next
		}

	}
	return nu.Next
}
