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
