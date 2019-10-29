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

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	total := 0
	loop := head
	for {
		if loop == nil {
			break
		}
		total++
		loop = loop.Next
	}
	newTail := head
	k = total - k
	for {
		if k > 0 {
			break
		}
		k = total + k
	}
	for {
		if k == 1 {
			break
		}
		if newTail.Next == nil {
			newTail = head
		} else {
			newTail = newTail.Next
		}
		k = k - 1
	}
	tmp := newTail.Next
	if tmp == nil {
		return head
	} else {
		newTail.Next = nil
		loop := tmp
		for {
			if loop.Next == nil {
				loop.Next = head
				break
			}
			loop = loop.Next
		}
	}
	return tmp
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	fast := head
	newTail := head
	for {
		if k == 0 {
			break
		}
		if fast.Next == nil {
			fast = head
		} else {
			fast = fast.Next
		}
		k--
	}
	for {
		if fast.Next == nil {
			break
		}
		newTail = newTail.Next
		fast = fast.Next
	}
	var ret *ListNode
	if newTail.Next == nil {
		ret = head
	} else {
		ret = newTail.Next
		newTail.Next = nil
		loop := ret
		for {
			if loop.Next == nil {
				loop.Next = head
				break
			}
			loop = loop.Next
		}
	}
	return ret
}
