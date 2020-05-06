package list

/*
Contents:
1. 环路检测[detectCycle]	--	★
2. 二叉树中的列表[isSubPath]	--	★★★
3. 分割链表[partition]
*/

/*
环路检测

给定一个有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？
*/
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

/*
二叉树中的列表

给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。
如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。
一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

示例 1：
			1
		4		4
		  2   2
		1   6	8
			   1  3
输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。

示例 2：
			1
		4		4
		  2   2
		1   6	8
			   1  3
输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true

示例 3：
输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。


提示：
二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
链表包含的节点数目在 1 到 100 之间。
二叉树包含的节点数目在 1 到 2500 之间。
*/
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSamePath(head, root) ||
		isSubPath(head, root.Left) ||
		isSubPath(head, root.Right)
}

func isSamePath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		return isSamePath(head.Next, root.Left) || isSamePath(head.Next, root.Right)
	}
	return false
}

/*
分割链表

编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

示例:
输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
*/
func partition(head *ListNode, x int) *ListNode {
	var left, right *ListNode
	for head != nil {
		if head.Val >= x {
			tmp := head.Next
			head.Next = right
			right = head
			head = tmp
		} else {
			tmp := head.Next
			head.Next = left
			left = head
			head = tmp
		}
	}
	l := left
	for l != nil {
		if l.Next == nil {
			l.Next = right
			break
		}
		l = l.Next
	}
	if left == nil {
		return right
	}
	return left
}
