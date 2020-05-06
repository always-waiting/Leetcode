package list

import (
	"testing"
)

func Test_Exercise2(t *testing.T) {
	{
		t.Log("环路检查......")
		aInt := []int{1, 2, 3, 4, 5}
		pos := 3
		a := newCycleList(aInt, pos)
		ret := detectCycle(a)
		if ret.String() == "[4->5]->..." {
			t.Log("检查正确")
		} else {
			t.Errorf("检查错误: %s", ret.String())
		}
	}
	{
		t.Log("二叉树中的列表......")
		treeNode := []interface{}{1, 4, 4, nil, 2, 2, nil, 1, nil, 6, 8, nil, nil, nil, nil, 1, 3}
		tree := createTreeByLevelLoop(treeNode)

		aInt := []int{4, 2, 8}
		a := newListNode(aInt)
		if isSubPath(a, tree) {
			t.Log("检查正确")
		} else {
			t.Errorf("检查错误")
		}
	}
}
