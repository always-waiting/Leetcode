package list

/*
给出一个以头节点 head 作为第一个节点的链表。链表中的节点分别编号为：node_1, node_2, node_3, ... 。
每个节点都可能有下一个更大值（next larger value）：对于 node_i，如果其 next_larger(node_i) 是 node_j.val，那么就有 j > i 且  node_j.val > node_i.val，而 j 是可能的选项中最小的那个。如果不存在这样的 j，那么下一个更大值为 0 。
返回整数答案数组 answer，其中 answer[i] = next_larger(node_{i+1}) 。
注意：在下面的示例中，诸如 [2,1,5] 这样的输入（不是输出）是链表的序列化表示，其头节点的值为 2，第二个节点值为 1，第三个节点值为 5 。

示例 1：
输入：[2,1,5]
输出：[5,5,0]

示例 2：
输入：[2,7,4,3,5]
输出：[7,0,5,5,0]

示例 3：
输入：[1,7,5,1,9,2,5,1]
输出：[7,9,9,9,0,5,0,0]

提示：
对于链表中的每个节点，1 <= node.val <= 10^9
给定列表的长度在 [0, 10000] 范围内
*/
func nextLargerNodes(head *ListNode) []int {
	ret := make([]int, 0)
	cur := head
	for cur != nil {
		cmp := cur.Next
		find := false
		for cmp != nil {
			if cmp.Val > cur.Val {
				ret = append(ret, cmp.Val)
				find = true
				break
			}
			cmp = cmp.Next
		}
		if !find {
			ret = append(ret, 0)
		}
		cur = cur.Next
	}
	return ret
}

func nextLargerNodes1(head *ListNode) []int {
	stack := Stack{items: make([]*StackNode, 0)}
	ret := make([]int, 0)
	cur := head
	idx := 0
	for cur != nil {
		if stack.Len() == 0 {
			stack.Push(&StackNode{
				Val: cur.Val,
				Idx: idx,
			})
		} else {
			for stack.Len() != 0 {
				if stack.Top().Val < cur.Val {
					node := stack.Pop()
					ret[node.Idx] = cur.Val
				} else {
					stack.Push(&StackNode{
						Val: cur.Val,
						Idx: idx,
					})
					break
				}
			}
			if stack.Len() == 0 {
				stack.Push(&StackNode{
					Val: cur.Val,
					Idx: idx,
				})
			}
		}
		ret = append(ret, 0)
		idx++
		cur = cur.Next
	}
	return ret
}

type StackNode struct {
	Val int
	Idx int
}

type Stack struct {
	items []*StackNode
}

func (this *Stack) Push(node *StackNode) {
	this.items = append(this.items, node)
}

func (this *Stack) Pop() (node *StackNode) {
	len := len(this.items)
	if len == 0 {
		return
	}
	node = this.items[len-1]
	this.items = this.items[:len-1]
	return node
}

func (this Stack) Len() int {
	return len(this.items)
}

func (this Stack) Top() *StackNode {
	return this.items[len(this.items)-1]
}

/*
给定一个链表（链表结点包含一个整型值）的头结点 head。
同时给定列表 G，该列表是上述链表中整型值的一个子集。
返回列表 G 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。

示例 1：
输入:
head: 0->1->2->3
G = [0, 1, 3]
输出: 2
解释:
链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。

示例 2：
输入:
head: 0->1->2->3->4
G = [0, 3, 1, 4]
输出: 2
解释:
链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

注意:
如果 N 是给定链表 head 的长度，1 <= N <= 10000。
链表中每个结点的值所在范围为 [0, N - 1]。
1 <= G.length <= 10000
G 是链表中所有结点的值的一个子集.
*/
func numComponents(head *ListNode, G []int) int {
	GM := make(map[int]bool)
	for _, val := range G {
		GM[val] = true
	}
	cur := head
	startFlag := false
	count := 0
	for cur != nil {
		if _, ok := GM[cur.Val]; ok {
			if !startFlag {
				startFlag = true
			}
		} else {
			if startFlag {
				startFlag = false
				count++
			}
		}
		cur = cur.Next
	}
	if startFlag {
		count++
	}
	return count
}

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :
给定这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5

说明 :
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	cur := head
	var total int
	for cur != nil {
		cur = cur.Next
		total++
	}
	cur = head
	step := 1
	var pre, ret, linkHead, prelinkHead *ListNode
	for cur != nil {
		if total < k && step == 1 {
			if linkHead != nil {
				linkHead.Next = cur
			}
			if ret == nil {
				ret = cur
			}
			break
		}
		if step == k {
			if ret == nil {
				ret = cur
			}
			if prelinkHead != nil {
				prelinkHead.Next = cur
			}
			step = 1
		} else {
			if step == 1 {
				prelinkHead, linkHead = linkHead, cur
			}
			step++
		}
		total--
		cur.Next, pre, cur = pre, cur, cur.Next
		if step == 1 {
			linkHead.Next = nil
		}
	}
	return ret
}

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	curs := make([]*ListNode, 0)
	for _, list := range lists {
		curs = append(curs, list)
	}
	pre := dummy
	for len(curs) != 0 {
		var node *ListNode
		var idx int
		for id, list := range curs {
			if list == nil {
				continue
			}
			if node == nil || node.Val > list.Val {
				node = list
				idx = id
			}
		}
		if curs[idx] != nil {
			curs[idx] = curs[idx].Next
		}
		if node != nil {
			pre.Next = &ListNode{
				Val: node.Val,
			}
			pre = pre.Next
		}
		tmpCurs := make([]*ListNode, 0)
		for _, val := range curs {
			if val != nil {
				tmpCurs = append(tmpCurs, val)
			}
		}
		curs = tmpCurs
	}
	return dummy.Next
}
