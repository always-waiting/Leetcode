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
