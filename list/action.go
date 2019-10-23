package list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
