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
		if idx == half {
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
			b = b.Next
			n++
		}
		a.Next = b
		if b != cur {
			b.Next = cur
		}
		idx++
	}
	cur.Next = nil
}

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	fast := head
	slow := head
	var preSlow *ListNode
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		preSlow, slow = slow, slow.Next
	}
	preSlow.Next = nil
	halfHead := slow.Next
	slow.Next = nil
	tree_node := &TreeNode{Val: slow.Val}
	tree_node.Left = sortedListToBST(head)
	tree_node.Right = sortedListToBST(halfHead)
	return tree_node
}

/*
给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
返回一个符合上述规则的链表的列表。
举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]

示例 1：
输入:
root = [1, 2, 3], k = 5
输出: [[1],[2],[3],[],[]]
解释:
输入输出各部分都应该是链表，而不是数组。
例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且 root.next.next.next = null。
第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。

示例 2：
输入:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
解释:
输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。


提示:
root 的长度范围： [0, 1000].
输入的每个节点的大小范围：[0, 999].
k 的取值范围： [1, 50].
*/
func splitListToParts(root *ListNode, k int) []*ListNode {
	var total int
	cur := root
	for cur != nil {
		total++
		cur = cur.Next
	}
	var a, b, c, length int
	if total >= k {
		c = 0
	} else {
		c = 1
	}
	for {
		m := (total + b) / (k - c)
		x := (total + b) % (k - c)
		if m == 0 {
			length = 1
			a = total
			b = 0
			break
		}
		if x == 0 {
			length = m
			a = k - b - c
			break
		}
		b++
		if c > 0 {
			c++
		}
	}
	ret := make([]*ListNode, 0)
	var num int
	for {
		if num == k {
			break
		}
		if a > 0 {
			cur := root
			step := 1
			for {
				if step == length {
					ret = append(ret, root)
					root, cur.Next = cur.Next, nil
					break
				}
				cur = cur.Next
				step++
			}
			a--
		} else if b > 0 {
			cur := root
			step := 1
			for {
				if step == length-1 {
					ret = append(ret, root)
					root, cur.Next = cur.Next, nil
					break
				}
				step++
				cur = cur.Next
			}
			b--
		} else {
			ret = append(ret, nil)
		}
		num++
	}
	return ret

}

/*
给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
删除完毕后，请你返回最终结果链表的头节点。

你可以返回任何满足题目要求的答案。
（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：
输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。

示例 2：
输入：head = [1,2,3,-3,4]
输出：[1,2,4]

示例 3：
输入：head = [1,2,3,-3,-2]
输出：[1]

提示：
给你的链表中可能有 1 到 1000 个节点。
对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000
*/
func removeZeroSumSublists(head *ListNode) *ListNode {
	var total int
	cur := head
	for cur == nil {
		total++
		cur = cur.Next
	}
	length := 1
	for length <= total {
		list := head
		var pre *ListNode
		reset := false
		for list != nil {
			sum := 0
			step := 0
			cur := list
			for step != length && cur != nil {
				sum = cur.Val + sum
				cur = cur.Next
				step++
			}
			if sum == 0 {
				if pre != nil {
					pre.Next = cur
				} else {
					head = cur
				}
				list = cur
				reset = true
			} else {
				pre, list = list, list.Next
			}
		}
		if reset {
			length = 1
		} else {
			length++
		}
	}
	return head
}

func removeZeroSumSublists1(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	sums := make([]int, 0)
	cur := dummy
	idx := -1
	for cur != nil {
		if idx == -1 {
			sums = append(sums, cur.Val)
		} else {
			sums = append(sums, sums[idx]+cur.Val)
		}
		idx++
		cur = cur.Next
	}
	idxList := make([]int, 0)
	skip := 0
	for idx1, val1 := range sums {
		for idx2, val2 := range sums {
			if val1 == val2 && idx1 < idx2 && idx1 >= skip && idx2 >= skip {
				idxList = append(idxList, idx1, idx2)
				skip = idx2 + 1
			}
		}
	}
	if len(idxList) == 0 {
		return head
	}
	cur = dummy
	var pre *ListNode
	idx = 0
	count := 0
	for cur != nil {
		if count == idxList[idx] {
			pre = cur
		}
		if count == idxList[idx+1] {
			pre.Next = cur.Next
		}
		cur = cur.Next
		count++
	}
	return dummy.Next
}
