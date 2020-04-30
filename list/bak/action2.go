package list

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:
输入: 4->2->1->3
输出: 1->2->3->4

示例 2:
输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/
func sortList(head *ListNode) *ListNode {
	total := 0
	loop := head
	for {
		if loop == nil {
			break
		}
		total++
		loop = loop.Next
	}
	intV := 1
	res := &ListNode{Next: head}
	for {
		if intV > total {
			break
		}
		pre, h := res, res.Next
		for {
			if h == nil {
				break
			}
			h1, i := h, intV
			for {
				if i == 0 || h == nil {
					break
				}
				h, i = h.Next, i-1
			}
			if i > 0 || h == nil {
				continue
			}
			h2, i := h, intV
			for {
				if i == 0 || h == nil {
					break
				}
				h, i = h.Next, i-1
			}
			c1, c2 := intV, intV-i
			for {
				if c1 == 0 || c2 == 0 {
					break
				}
				if h1.Val < h2.Val {
					pre.Next, h1, c1 = h1, h1.Next, c1-1
				} else {
					pre.Next, h2, c2 = h2, h2.Next, c2-1
				}
				pre = pre.Next
			}
			if c1 != 0 {
				pre.Next = h1
			} else {
				pre.Next = h2
			}
			for {
				if c1 > 0 || c2 > 0 {
					pre, c1, c2 = pre.Next, c1-1, c2-1
				} else {
					break
				}
			}
			pre.Next = h
		}
		intV = 2 * intV
	}
	return res.Next
}

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。

示例:
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, dummy.Next
	var large, tmp *ListNode
	for {
		if cur == nil {
			if tmp != nil {
				tmp.Next = nil
			}
			break
		}
		if cur.Val < x {
			pre.Next, pre, cur = cur, cur, cur.Next
		} else {
			if large == nil {
				large, tmp, cur = cur, cur, cur.Next
			} else {
				tmp.Next, tmp, cur = cur, cur, cur.Next
			}
		}
	}
	pre.Next = large
	return dummy.Next
}

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/
func oddEvenList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	step := 0
	pre, cur := dummy, dummy.Next
	var odd, tmp *ListNode
	for {
		if cur == nil {
			if tmp != nil {
				tmp.Next = nil
			}
			break
		}
		if step%2 == 0 {
			pre.Next, pre, cur = cur, cur, cur.Next
		} else {
			if odd == nil {
				odd, tmp, cur = cur, cur, cur.Next
			} else {
				tmp.Next, tmp, cur = cur, cur, cur.Next
			}

		}
		step++
	}
	pre.Next = odd
	return dummy.Next
}

/*
对链表进行插入排序。
插入排序算法：
1. 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
2. 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
3. 重复直到所有输入数据插入完为止。

示例 1：
输入: 4->2->1->3
输出: 1->2->3->4

示例 2：
输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/
/*
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next
	for {
		if cur == nil {
			break
		}
		if cur.Val < pre.Val {
			insert := insert
			pre.Next, cur = cur.Next, cur.Next
			if head.Val > insert.Val {
				insert.Next, head = head, insert
			} else {
				preloop, loop := head, head.Next
				for {
					if loop == cur {
						break
					}
					if insert.Val < loop.Val {
						insert.Next = loop
						preloop.Next = insert
						break
					}
					preloop, loop = loop, loop.Next
				}
			}
		} else {
			pre.Next, pre, cur = cur, cur, cur.Next
		}
	}
	return head
}
*/

/*
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:
输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1N, l2N := 0, 0
	cur1, cur2 := l1, l2
	for {
		if cur1 == nil && cur2 == nil {
			break
		}
		if cur1 != nil {
			cur1 = cur1.Next
			l1N++
		}
		if cur2 != nil {
			cur2 = cur2.Next
			l2N++
		}
	}
	ret := &ListNode{}
	var large, small *ListNode
	var largeN, smallN int
	if l1N < l2N {
		largeN = l2N
		large = l2
		smallN = l1N
		small = l1
	} else {
		largeN = l1N
		large = l1
		smallN = l2N
		small = l2
	}
	cur := ret
	totalRet := 1
	for {
		if large == nil && small == nil {
			break
		}
		if largeN == smallN {
			sum := &ListNode{
				Val: large.Val + small.Val,
			}
			cur.Next, cur = sum, sum
			largeN--
			smallN--
			large, small = large.Next, small.Next
		} else if largeN > smallN {
			sum := &ListNode{Val: large.Val}
			cur.Next, cur = sum, sum
			largeN--
			large = large.Next
		}
		totalRet++
	}
	bit := 2
	for {
		count := bit
		cur := ret
		postcur := ret.Next
		for {
			if count == totalRet {
				cur.Val = postcur.Val/10 + cur.Val
				postcur.Val = postcur.Val % 10
				break
			}
			cur, postcur = postcur, postcur.Next
			count++
		}
		if bit == totalRet {
			break
		}
		bit++
	}
	if ret.Val != 0 {
		return ret
	} else {
		return ret.Next
	}
	return nil
}
