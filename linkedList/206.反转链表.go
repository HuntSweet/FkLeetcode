package linkedList

 type ListNode struct {
	     Val int
	     Next *ListNode
 }

 //双指针
func reverseList3(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head

	for cur != nil{
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre

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