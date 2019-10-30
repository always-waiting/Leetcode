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
