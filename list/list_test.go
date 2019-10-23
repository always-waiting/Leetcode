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
