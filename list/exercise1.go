package list

/*
Content:
1. 链表相交
2. 回文链表
3. 移除重复节点
4. 删除中间节点 -- start1
5. 两个链表的第一个公共节点 -- same to 链表相交
*/

/*
链表相交

给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，
而不是基于节点的值。换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同
一节点（引用完全相同），则这两个链表相交。

示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例 3：
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。

注意：
如果两个链表没有交点，返回 null 。
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}
	return h1
}

/*
回文链表

编写一个函数，检查输入的链表是否是回文的。

示例 1：
输入： 1->2
输出： false

示例 2：
输入： 1->2->2->1
输出： true

进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func isPalindrome(head *ListNode) bool {
	/*
		if head == nil {
			return true
		}
		// O(n)空间复杂度,O(n)时间复杂度
		l := head.Reverse()
		for head != nil {
			if head.Val != l.Val {
				return false
			}
			head = head.Next
			l = l.Next
		}
		return true
	*/
	// 快慢指针，寻找中间位置，反转查看
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	var rSlow *ListNode
	for fast != nil && fast.Next != nil {
		tmp := slow
		fast = fast.Next.Next
		slow = slow.Next
		tmp.Next = rSlow
		rSlow = tmp
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != rSlow.Val {
			return false
		}
		slow = slow.Next
		rSlow = rSlow.Next
	}
	return true
}

/*
移除重复节点

编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:
输入：[1, 2, 3, 3, 2, 1]
输出：[1, 2, 3]

示例2:
输入：[1, 1, 1, 1, 2]
输出：[1, 2]

提示：
链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。

进阶：
如果不得使用临时缓冲区，该怎么解决？
*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	/*
		// 使用了缓冲区,用一个指针记录前一个节点
		// O(n)时间复杂度，O(n)空间复杂度
		tmp := make(map[int]bool)
		var pre *ListNode
		l := head
		for l != nil {
			if _, ok := tmp[l.Val]; !ok {
				tmp[l.Val] = true
				pre = l
				l = l.Next
			} else {
				l = l.Next
				pre.Next = l
			}
		}
		return head
	*/
	// 不用缓存区，O(n*n)时间复杂度，O(1)空间复杂度
	var pre *ListNode
	l := head
	for l != nil {
		if pre == nil {
			pre = l
			l = l.Next
		} else {
			tmp := head
			has := false
			for tmp != pre.Next {
				if tmp.Val == l.Val {
					has = true
					break
				}
				tmp = tmp.Next
			}
			if has {
				pre.Next = l.Next
				l = l.Next
			} else {
				pre = l
				l = l.Next
			}
		}
	}
	return head
}

/*
删除中间节点

实现一种算法，删除单向链表中间的某个节点（除了第一个和最后一个节点，不一定是中间节点），假定你只能访问该节点。

示例：
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
*/
func deleteNode(node *ListNode) {
	*node = *node.Next
}

/*
两个链表的第一个公共节点

输入两个链表，找出它们的第一个公共节点。
如下面的两个链表：
A: a1 - a2
			\
				c1 - c2 - c3

			/
B: b1 - b2
在节点 c1 开始相交。

示例 1：
A:	   4 - 1
			\
			  8 - 4 - 5
			/
B: 5 - 0 -1
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：
A: 0 - 9 - 1
			\
			  2 - 4
			/
B:		   3

输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例 3：
A: 2 - 6 - 4


B:     1 - 5
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。

注意：
如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
本题与主站 160 题相同：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
*/
func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}
