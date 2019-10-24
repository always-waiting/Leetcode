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

func TestReverse(t *testing.T) {
	data := []int{1, 2, 3, 4}
	list := newListNode(data)
	newList := reverseList(list)
	report := []int{4, 3, 2, 1}
	loop := list
	i := 0
	for {
		if loop == nil {
			break
		}
		if loop.Val != data[i] {
			t.Errorf("反转修改了原始列表")
			break
		}
		i++
		loop = loop.Next
	}
	loop = newList
	i = 0
	for {
		if loop == nil {
			break
		}
		if loop.Val != report[i] {
			t.Errorf("反转结果不对")
			break
		}
		i++
		loop = loop.Next
	}
}

func TestPalindrome(t *testing.T) {
	list := newListNode([]int{1, 2, 2, 1})
	if !isPalindrome(list) {
		t.Errorf("回文判断失败")
	}
}

func TestPalindrome1(t *testing.T) {
	data := []int{1, 2, 3, 2, 1}
	list := newListNode(data)
	if !isPalindrome1(list) {
		t.Errorf("回文判断失败")
	}
	i := 0
	for {
		if list == nil {
			break
		}

		if list.Val != data[i] {
			t.Errorf("判断回文后，原链表变更")
		}
		i++
		list = list.Next
	}
	if i != len(data) {
		t.Errorf("判断回文后，原链表变更(长度): %d -> %d", len(data), i)
	}

}
