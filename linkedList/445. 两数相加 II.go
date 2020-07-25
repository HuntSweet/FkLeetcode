package linkedList

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	sta1 := make([]*ListNode,0)
	sta2 := make([]*ListNode,0)

	for l1 != nil{
		sta1 =append(sta1,l1)
		l1 = l1.Next
	}

	for l2 != nil{
		sta2 = append(sta2,l2)
		l2 = l2.Next
	}

	newList := &ListNode{}
	head := newList
	n := 0
	for len(sta1) != 0 || len(sta2) != 0 || n > 0{
		len1 := len(sta1)
		len2 := len(sta2)

		var v1,v2 int
		if len1 > 0{
			v1 = sta1[len1-1].Val
			sta1 = sta1[:len1-1]
		}
		if len2 > 0{
			v2 = sta2[len2-1].Val
			sta2 = sta2[:len2-1]
		}


		node := &ListNode{Val:(v1+v2+n)%10}
		node.Next = newList.Next
		newList.Next = node

		n = (v1+v2+n)/10
	}

	return head.Next
}
