package list

/*
修改l1和l2的链表信息
*/
func mergeTwoLists_bad1(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	if l1.Val <= l2.Val {
		ret = l1
		l1 = l1.Next
	} else {
		ret = l2
		l2 = l2.Next
	}
	loop := ret
	for {
		if l1 == nil && l2 != nil {
			loop.Next = l2
			l2 = l2.Next
			loop = loop.Next
			continue
		}
		if l2 == nil && l1 != nil {
			loop.Next = l1
			l1 = l1.Next
			loop = loop.Next
			continue
		}
		if l1 == nil && l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			loop.Next = l1
			l1 = l1.Next
		} else {
			loop.Next = l2
			l2 = l2.Next
		}
		loop = loop.Next
	}
	return ret
}

/*
赋值时，用了loop.Next = lx.Next
会产生2个效果
1. 丢数据
2. 在之后的循环中，修改lx的Next指针
从而产生死循环
*/
func mergeTwoLists_bad0(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	if l1.Val <= l2.Val {
		ret = l1
		l1 = l1.Next
	} else {
		ret = l2
		l2 = l2.Next
	}
	loop := ret
	i := 0
	for {
		if i > 10 {
			break
		}
		if l1.Next == nil && l2.Next != nil {
			loop.Next = l2.Next
			l2 = l2.Next
			loop = loop.Next
			continue
		}
		if l2.Next == nil && l1.Next != nil {
			loop.Next = l1.Next
			l1 = l1.Next
			loop = loop.Next
			continue
		}
		if l1.Next == nil && l2.Next == nil {
			break
		}
		if l1.Val <= l2.Val {
			loop.Next = l1.Next
			l1 = l1.Next
		} else {
			loop.Next = l2.Next
			l2 = l2.Next
		}
		loop = loop.Next
		i++
	}
	return ret
}
