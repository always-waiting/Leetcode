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
