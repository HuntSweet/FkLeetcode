package linkedList

//跟86一样
func oddEvenList(head *ListNode) *ListNode {
	left, right := &ListNode{},&ListNode{}
	pre,next := left,right

	n := 0
	for head != nil{
		if n % 2 == 0{
			left.Next = head
			left = left.Next
		}else{
			right.Next = head
			right = right.Next
		}
		n++
		head = head.Next
	}

	right.Next = nil
	left.Next = next.Next

	return pre.Next
}