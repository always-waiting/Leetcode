package list

import (
	"testing"
)

func Test_Exercise1(t *testing.T) {
	aInt := []int{4, 1, 8, 4, 5}
	bInt := []int{5, 0, 1, 8, 4, 5}
	{
		t.Log("链表相交......")
		a, b := newIntersectionList(aInt, bInt, 8, 2, 3)
		c := getIntersectionNode(a, b)
		if c.String() == "8->4->5" {
			t.Log("检查正确")
		} else {
			t.Errorf("检查错误: %s", c.String())
		}
	}
	{
		aInt := []int{1, 2, 2, 1}
		t.Log("回文链表......")
		a := newListNode(aInt)
		if isPalindrome(a) == true {
			t.Log("检查正确")
		} else {
			t.Errorf("检查错误")
		}
	}
	{
		t.Log("移除重复节点......")
		aInt := []int{1, 2, 3, 3, 2, 1}
		a := newListNode(aInt)
		a = removeDuplicateNodes(a)
		if a.String() != "1->2->3" {
			t.Errorf("检查错误")
		} else {
			t.Log("检查正确")
		}

	}
	{
		t.Log("删除中间节点......")
		aInt := []int{4, 5, 1, 9}
		a := newListNode(aInt)
		num5 := a.Next
		deleteNode(num5)
		if a.String() != "4->1->9" {
			t.Errorf("检查错误")
		} else {
			t.Log("检查正确")
		}
	}
}
