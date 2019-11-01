package list

/*
给定一个单链表 L：L0→ L1→ …→ Ln-1→ Ln ，
将其重新排列后变为： L0→ Ln→ L1→ Ln-1→ L2→ Ln-2→ …
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast := head
	slow := head
	half := 0
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		half++
	}
	if fast != nil {
		slow = slow.Next
	}
	idx := 0
	cur := head
	for {
		if idx > half {
			break
		}
		a := cur
		cur = cur.Next
		b := slow
		n := 1
		for {
			if n+idx == half {
				break
			}
			if b == nil {
				break
			}
			b = b.Next
			n++
		}
		a.Next = b
		if b != cur {
			b.Next = cur
		}
		idx++
	}

}
