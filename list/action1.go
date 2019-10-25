package list

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func swapPairs1(head *ListNode) *ListNode {
	ret := &ListNode{Next: head}
	cur := head
	loop := ret
	for {
		if cur == nil || cur.Next == nil {
			break
		}
		tmp := cur.Next
		cur.Next = tmp.Next
		loop.Next = tmp
		tmp.Next = cur
		loop = cur
		cur = cur.Next
	}
	return ret.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}
