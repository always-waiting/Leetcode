package list

/*
Contents:
1. 环路检测[detectCycle]	--	★
2. 二叉树中的列表[isSubPath]	--	★★★
3. 分割链表[partition]
4. 扁平化多级双向链表[flatten]	--	★
5. 链表求和[addTwoNumbers]
6. 复杂链表的复制[copyRandomList]	--	★★
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

/*
扁平化多级双向链表

多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

示例 1：
输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]

示例 2：
输入：head = [1,2,null,3]
输出：[1,3,2]
输入的多级列表如下图所示：

  1---2---NULL
  |
  3---NULL

示例 3：
输入：head = []
输出：[]


如何表示测试用例中的多级链表？
以 示例 1 为例：
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
序列化其中的每一级之后：
[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。
[1,2,3,4,5,6,null]
[null,null,7,8,9,10,null]
[null,11,12,null]
合并所有序列化结果，并去除末尾的 null 。
[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]

提示：
节点数目不超过 1000
1 <= Node.val <= 10^5
*/
func flatten(root *Node) *Node {
	l := root
	for l != nil {
		if l.Child != nil {
			tmp := l.Next
			l.Next = flatten(l.Child)
			l.Child.Prev = l
			l.Child = nil
			for l.Next != nil {
				l = l.Next
			}
			if tmp == nil {
				continue
			} else {
				tmp.Prev = l
				l.Next = tmp
				l = l.Next
			}
		} else {
			l = l.Next
		}
	}
	return root
}

/*
链表求和

给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。

示例：
输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。

示例：
输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	var carry int
	var l *ListNode
	for l1 != nil || l2 != nil {
		var a1, a2 int
		if l1 != nil {
			a1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a2 = l2.Val
			l2 = l2.Next
		}
		val := (a1 + a2 + carry) % 10
		carry = (a1 + a2 + carry) / 10
		tmp := &ListNode{Val: val}
		if ret == nil {
			ret = tmp
			l = ret
		} else {
			l.Next = tmp
			l = tmp
		}
	}
	if carry != 0 {
		l.Next = &ListNode{Val: carry}
	}
	return ret
}

/*
复杂链表的复制

请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

提示：
-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000

注意：本题与主站 138 题相同：https://leetcode-cn.com/problems/copy-list-with-random-pointer/
*/

func copyRandomList(head *Node1) *Node1 {
	/*
		// hash法，空间复杂度高
		var pre, ret *Node1
		l := head
		for l != nil {
			if ret == nil {
				ret = &Node1{Val: l.Val}
				pre = ret
				l = l.Next
			} else {
				pre.Next = &Node1{Val: l.Val}
				l = l.Next
				pre = pre.Next
			}
		}
		l = head
		rl := ret
		for l != nil {
			if l.Random != nil {
				il := head
				count := 0
				for il != nil {
					if il == l.Random {
						break
					}
					count++
					il = il.Next
				}
				irl := ret
				for irl != nil {
					if count == 0 {
						rl.Random = irl
					}
					count--
					irl = irl.Next
				}
			}
			l = l.Next
			rl = rl.Next
		}
		return ret
	*/
	l := head
	// 2倍链表
	for l != nil {
		tmp := &Node1{Val: l.Val}
		nTmp := l.Next
		l.Next = tmp
		tmp.Next = nTmp
		l = nTmp
	}
	// l.Next.Random = l.Random.Next
	l = head
	for l != nil {
		if l.Random != nil {
			l.Next.Random = l.Random.Next
		}
		l = l.Next.Next
	}
	// 把链表拆分成2个
	var pre, ret *Node1
	l = head
	for l != nil {
		tmp := l
		if ret == nil {
			ret = l.Next
			pre = ret
			l = l.Next.Next
		} else {
			pre.Next = l.Next
			pre = pre.Next
			l = l.Next.Next
		}
		tmp.Next = l
	}
	return ret
}
