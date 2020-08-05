package linkedList

//归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//找到中点
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	//分别进行排序
	left, right := sortList(head), sortList(mid)

	//合并排序后的两个链表
	h := &ListNode{0, nil}
	cur := h
	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}

	if right != nil {
		cur.Next = right
	}

	return h.Next
}
