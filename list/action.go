package list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	var ret *ListNode
	if l1.Val <= l2.Val {
		ret = &ListNode{
			Val: l1.Val,
		}
		l1 = l1.Next
	} else {
		ret = &ListNode{
			Val: l2.Val,
		}
		l2 = l2.Next
	}
	loop := ret
	for {
		if l1 == nil && l2 != nil {
			loop.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
			loop = loop.Next
			continue
		}
		if l2 == nil && l1 != nil {
			loop.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
			loop = loop.Next
			continue
		}
		if l1 == nil && l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			loop.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			loop.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		loop = loop.Next
	}
	return ret
}

/*
生成新的反转列表，原来的列表不变化
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{
		Val: head.Val,
	}
	head = head.Next
	for {
		if head == nil {
			break
		}
		tmp := &ListNode{
			Val: head.Val,
		}
		tmp.Next = ret
		ret = tmp
		head = head.Next
	}
	return ret
}

/*
O(n)的时间复杂度，O(n)的空间复杂度
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	rhead := reverseList(head)
	for {
		if head == nil {
			break
		}
		if head.Val != rhead.Val {
			return false
		}
		head = head.Next
		rhead = rhead.Next
	}
	return true
}

/*
O(n)的时间复杂度，O(1)的空间复杂度
*/
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针找到中间值
	fast := head
	slow := head
	isOdd := false
	for {
		if fast == nil {
			break
		}
		if fast.Next == nil {
			isOdd = true
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	if isOdd {
		slow = slow.Next
	}
	// 反转慢指针链表
	var rslow *ListNode
	for {
		if slow == nil {
			break
		}
		tmp := slow.Next
		slow.Next = rslow
		rslow = slow
		slow = tmp
	}
	// 比较前后部分
	hrslow := rslow
	for {
		if rslow == nil {
			break
		}
		if head.Val != rslow.Val {
			return false
		}
		head = head.Next
		rslow = rslow.Next
	}
	// 没有如下代码，原链表被修改
	for {
		if hrslow == nil {
			break
		}
		tmp := hrslow.Next
		hrslow.Next = slow
		slow = hrslow
		hrslow = tmp
	}
	return true
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func removeElements(head *ListNode, val int) *ListNode {
	loop := &ListNode{}
	loop.Next = head
	ret := loop
	for {
		if loop.Next == nil {
			break
		}
		if loop.Next.Val == val {
			loop.Next = loop.Next.Next
		} else {
			loop = loop.Next
		}
	}
	return ret.Next
}
