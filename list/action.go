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

// 消耗的内存过大
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var totalA, totalB1, totalB2 int
	var endA, endB *ListNode
	// 计算A的长度,记录A的最后元素
	loopA := headA
	for {
		if loopA == nil {
			break
		} else {
			endA = loopA
		}
		loopA = loopA.Next
		totalA++
	}
	// 计算B的长度,记录B的最后元素
	loopB := headB
	for {
		if loopB == nil {
			break
		} else {
			endB = loopB
		}
		loopB = loopB.Next
		totalB1++
	}
	if endA != endB { // 没有交集
		return nil
	}
	// 反转A
	loopA = headA
	endA = nil
	for {
		if loopA == nil {
			break
		}
		tmp := loopA.Next
		loopA.Next = endA
		endA = loopA
		loopA = tmp
	}
	// 再次计算B的长度
	loopB = headB
	for {
		if loopB == nil {
			break
		}
		loopB = loopB.Next
		totalB2++
	}
	pastB := (totalB2 + totalB1 - totalA) / 2
	ret := headB
	i := 0
	for {
		if i == pastB {
			break
		}
		ret = ret.Next
		i++
	}
	// 再次反转A
	for {
		if endA == nil {
			break
		}
		tmp := endA.Next
		endA.Next = loopA
		loopA = endA
		endA = tmp
	}
	return ret
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	loopA := headA
	loopB := headB
	for {
		if loopA == nil && loopB == nil {
			break
		}
		if loopA == nil {
			loopA = headB
		}
		if loopB == nil {
			loopB = headA
		}
		if loopA == loopB {
			return loopA
		}
		loopA = loopA.Next
		loopB = loopB.Next
	}
	return nil
}

// 任意链表
func deleteDuplicates(head *ListNode) *ListNode {
	data := make(map[int]bool, 0)
	loop := head
	post := head
	for {
		if loop == nil {
			break
		}
		if _, ok := data[loop.Val]; ok {
			post.Next = loop.Next
		} else {
			data[loop.Val] = true
			post = loop
		}
		loop = loop.Next

	}
	return head
}

// 排序链表去重
func deleteDuplicates1(head *ListNode) *ListNode {
	var val int
	loop := head
	post := head
	for {
		if loop == nil {
			break
		}
		if val == loop.Val {
			post.Next = loop.Next
		} else {
			val = loop.Val
			post = loop
		}
		loop = loop.Next
	}
	return head
}

// 排序链表去重
func deleteDuplicates2(head *ListNode) *ListNode {
	loop := head
	for {
		if loop == nil && loop.Next == nil {
			break
		}
		if loop.Next.Val == loop.Val {
			loop.Next = loop.Next.Next
		} else {
			loop = loop.Next
		}
	}
	return head
}

/*
链表至少包含两个节点。
链表中所有节点的值都是唯一的。
给定的节点为非末尾节点并且一定是链表中的一个有效节点。
不要从你的函数中返回任何结果。
注意:
没有给链表的头节点，node就是链表中需要删除的节点位置
*/
func deleteNode(node *ListNode) {
	// 可以理解为把下一节点的信息拷贝的当前节点
	// 从而完成了当前节点的删除
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
