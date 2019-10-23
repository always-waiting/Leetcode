package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(data []int) *ListNode {
	var ret, loop *ListNode
	for _, val := range data {
		tmp := &ListNode{Val: val}
		if ret == nil {
			ret = tmp
			loop = tmp
			continue
		} else {
			loop.Next = tmp
			loop = tmp
		}
	}
	return ret
}
