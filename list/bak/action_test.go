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

func TestRemoveElements(t *testing.T) {
	data := []int{6, 1, 2, 3, 6, 2, 6, 1, 6}
	val := 6
	list := newListNode(data)
	newlist := removeElements(list, val)
	i := 0
	for {
		if newlist == nil {
			break
		}
		if data[i] != val {
			if data[i] != newlist.Val {
				t.Errorf("删除元素失败")
				break
			}
			newlist = newlist.Next
		}
		i++
	}
	/*
		注意list的链表结果,第一个元素总数存在的
		因为那是调用函数的指针
	*/
	loop := list
	for {
		if loop == nil {
			break
		}
		t.Log(loop.Val)
		loop = loop.Next
	}

}

func TestIntersection(t *testing.T) {
	dataA := []int{4, 1, 8, 4, 5}
	dataB := []int{5, 0, 1, 8, 4, 5}
	interVal := 8
	skipA := 2
	skipB := 3
	listA, listB := newIntersectionList(dataA, dataB, interVal, skipA, skipB)
	interList := getIntersectionNode1(listA, listB)
	if interList == nil {
		t.Errorf("交集链表获取失败")
		return
	}
	loop := interList
	testA := listA
	testB := listB
	i := 0
	j := 0
	for {
		var tmpA, tmpB *ListNode
		if i < skipA {
			if dataA[i] != testA.Val {
				t.Errorf("修改了原链表A")
				break
			}
			i++
			testA = testA.Next
		} else {
			tmpA = testA
		}
		if j < skipB {
			if dataB[j] != testB.Val {
				t.Error("修改了原链表B")
				break
			}
			j++
			testB = testB.Next
		} else {
			tmpB = testB
		}
		if tmpA != nil && tmpB != nil {
			if loop != tmpA ||
				loop != tmpB {
				t.Error("交集获取失败")
				break
			}
			loop = loop.Next
			i++
			j++
			testA = testA.Next
			testB = testB.Next
		}
		if i > skipA && j > skipB &&
			tmpA == nil && tmpB == nil {
			break
		}
	}
}
