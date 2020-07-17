package linkedList

 type ListNode struct {
	     Val int
	     Next *ListNode
 }

 //双指针
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	n := head
	var c *ListNode
	c = nil
	for n != nil{
		head = head.Next
		n.Next = c
		c = n
		n = head
	}
	return c
}
 //递归1，超级低效
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	l := reverseList(head.Next)
	head.Next = nil
	t := l
	for l.Next != nil{
		l = l.Next
	}
	l.Next = head
	return t
}
//递归2
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	l := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return l
}