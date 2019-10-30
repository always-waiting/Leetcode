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

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	step := 1
	loop := head
	var start, end, preStart, postEnd *ListNode
	for {
		if loop == nil {
			break
		}
		if step == m {
			start = loop
		}
		if step == m-1 {
			preStart = loop
		}
		if step == n {
			end = loop
			postEnd = end.Next
			break
		}
		step++
		loop = loop.Next
	}
	loop = start
	for {
		if loop == end {
			loop.Next = postEnd
			if preStart == nil {
				head = loop
			} else {
				preStart.Next = loop
			}
			break
		}
		tmp := loop.Next
		loop.Next = postEnd
		postEnd = loop
		loop = tmp
	}
	return head
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	step := 1
	var prev, cur, con, tail *ListNode
	cur = head
	for {
		if step == m {
			con = prev
			tail = cur
		}
		tmp := cur.Next
		if step > m && step < n {
			cur.Next = prev
		}
		if step == n {
			tail.Next = tmp
			if con == nil {
				head = cur
			} else {
				con.Next = cur
			}
			cur.Next = prev
			break
		}
		prev = cur
		cur = tmp
		step++
	}
	return head
}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	fast := head
	slow := dummy
	var del *ListNode
	var delflag bool
	for {
		if fast == nil {
			break
		}
		if fast.Next == nil {
			if delflag {
				if del.Val == fast.Val {
					slow.Next = fast.Next
				} else {
					del = nil
					delflag = false
					slow = fast
				}
			}
			break
		}
		if fast.Val == fast.Next.Val {
			slow.Next = fast.Next.Next
			del = fast
			delflag = true
			fast = fast.Next.Next
		} else {
			if delflag {
				if del.Val == fast.Val {
					slow.Next = fast.Next
				} else {
					del = nil
					delflag = false
					slow = fast
				}
			} else {
				del = nil
				delflag = false
				slow = fast
			}
			fast = fast.Next
		}
	}
	return dummy.Next
}

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？
*/
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	loop := head
	for {
		if slow == loop {
			break
		}
		loop = loop.Next
		slow = slow.Next
	}
	return loop
}
