package list

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	tmp := ret
	overflow := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		a1 := ListNode{}
		a2 := ListNode{}
		sum := ListNode{}
		if l1 != nil {
			a1.Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a2.Val = l2.Val
			l2 = l2.Next
		}
		sum.Val = a1.Val + a2.Val + overflow
		if sum.Val > 9 {
			sum.Val = sum.Val - 10
			overflow = 1
		} else {
			overflow = 0
		}
		tmp.Next = &sum
		tmp = &sum
	}
	if overflow == 1 {
		tmp.Next = &ListNode{Val: 1}
	}
	return ret.Next
}
