package list

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("链表测试文本")
	m.Run()
}

func TestCreate(t *testing.T) {
	list := newListNode([]int{1, 2, 3})
	i := 1
	for {
		if list == nil {
			break
		}
		if list.Val != i {
			t.Error("链表生成失败")
			break
		}
		i++
		list = list.Next
	}
}

func TestCycleCreate(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	pos := 3
	list := newCycleList(data, pos)
	if !hasCycle(list) {
		t.Errorf("单环链表生成失败")
	}
}

func TestReverse1(t *testing.T) {
	t.Log("运行反转函数reverseList1")
	data := []int{1, 2, 3, 4, 5}
	list := newListNode(data)
	newList := list.reverseList1()
	report := []int{5, 4, 3, 2, 1}
	loop := newList
	i := 0
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
	if list.Next == nil {
		t.Logf("反转函数把原列表变量指到了反转对尾")
	}
}

func TestReverse2(t *testing.T) {
	t.Log("运行反转函数reverseList2")
	list := newListNode([]int{1, 2, 3, 4, 5})
	list.reverseList2()
	report := []int{5, 4, 3, 2, 1}
	loop := list
	i := 0
	for {
		if loop == nil {
			break
		}
		if loop.Val != report[i] {
			t.Errorf("反转结果不对: 第%d个 %d <-> %d", i+1, loop.Val, report[i])
			break

		}
		i++
		loop = loop.Next
	}
	t.Log("反转函数修改了原链表,但是循环过多")
}

func TestReverse3(t *testing.T) {
	t.Log("运行反转函数reverseList3")
	list := newListNode([]int{1, 2, 3, 4, 5})
	list.reverseList3()
	report := []int{5, 4, 3, 2, 1}
	loop := list
	i := 0
	for {
		if loop == nil {
			break
		}
		if loop.Val != report[i] {
			t.Errorf("反转结果不对: 第%d个 %d <-> %d", i+1, loop.Val, report[i])
			break

		}
		i++
		loop = loop.Next
	}
	t.Log("反转函数修改了原链表,但是内存占用过大")
}

func TestIntersectionCreate(t *testing.T) {
	dataA := []int{4, 1, 8, 4, 5}
	dataB := []int{5, 0, 1, 8, 4, 5}
	interVal := 8
	skipA := 2
	skipB := 3
	listA, listB := newIntersectionList(dataA, dataB, interVal, skipA, skipB)
	loopA := listA
	loopB := listB
	i := 0
	j := 0
	for {
		var tmpA, tmpB *ListNode
		if loopA != nil {
			if i < skipA {
				if dataA[i] != loopA.Val {
					t.Errorf("交集链表生成失败,第一个链表: 第%d个 %d <-> %d", i, dataA[i], loopA.Val)
					break
				}
				i++
				loopA = loopA.Next
			} else {
				tmpA = loopA
			}
		}
		if loopB != nil {
			if j < skipB {
				if dataB[j] != loopB.Val {
					t.Errorf("交集链表生成失败,第二个链表: 第%d个 %d <-> %d", i, dataB[j], loopB.Val)
					break
				}
				j++
				loopB = loopB.Next
			} else {
				tmpB = loopB
			}
		}
		if tmpA != nil && tmpB != nil {
			if tmpA != tmpB {
				t.Errorf("交集链表生成失败,交集部分不对")
				t.Log(tmpA.Val)
				t.Log(tmpB.Val)
				break
			}
			i++
			j++
			loopA = loopA.Next
			loopB = loopB.Next
		}
		if loopA == nil && loopB == nil {
			break
		}
	}
}
