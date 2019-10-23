package list

import (
	"testing"
)

func TestMerge(t *testing.T) {
	l1 := newListNode([]int{1, 2, 5})
	l2 := newListNode([]int{1, 2, 4, 6})
	l3 := mergeTwoLists(l1, l2)
	report := []int{1, 1, 2, 2, 4, 5, 6}
	i := 0
	for {
		if l3 == nil {
			break
		}
		if report[i] != l3.Val {
			t.Errorf("合并在%d个元素出错%d<->%d", i, report[i], l3.Val)
			break
		}
		l3 = l3.Next
		i++
	}
	report = []int{1, 2, 5}
	loop := l1
	i = 0
	for {
		if loop == nil {
			break
		}
		if report[i] != loop.Val {
			t.Errorf("合并函数修改了输入链表")
			break
		}
		loop = loop.Next
		i++
	}
}
