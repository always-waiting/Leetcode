package main

import (
	"reflect"
	"testing"
)

func Test_divingBoard(t *testing.T) {
	s, l, k := 1, 2, 3
	ret := divingBoard(s, l, k)
	expect := []int{3, 4, 5, 6}
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("跳回板结果错误: got(%v), expect(%v)", ret, expect)
	}
	s, l, k = 1, 1, 0
	expect = []int{}
	ret = divingBoard(s, l, k)
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("跳回板结果错误: got(%v), expect(%v)", ret, expect)
	}
	s, l, k = 1, 1, 10000
	expect = []int{k}
	ret = divingBoard(s, l, k)
	if !reflect.DeepEqual(ret, expect) {
		t.Errorf("跳回板结果错误: got(%v), expect(%v)", ret, expect)
	}
}

func Test_respace(t *testing.T) {
	dict := []string{"app", "less", "apple"}
	sent := "appless"
	ret := respace(dict, sent)
	if ret != 0 {
		t.Errorf("恢复空格结构错误: got(%d), expect(%d)", ret, 0)
	}
	dict = []string{"looked", "just", "like", "her", "brother"}
	sent = "jesslookedjustliketimherbrother"
	ret = respace(dict, sent)
	if ret != 7 {
		t.Errorf("恢复空格结构错误: got(%d), expect(%d)", ret, 7)
	}
}

func Test_respace2(t *testing.T) {
	dict := []string{"app", "less", "apple"}
	sent := "appless"
	ret := respace2(dict, sent)
	if ret != 0 {
		t.Errorf("恢复空格结构错误: got(%d), expect(%d)", ret, 0)
	}
	dict = []string{"looked", "just", "like", "her", "brother"}
	sent = "jesslookedjustliketimherbrother"
	ret = respace2(dict, sent)
	if ret != 7 {
		t.Errorf("恢复空格结构错误: got(%d), expect(%d)", ret, 7)
	}
}

func Test_maxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	ret := maxProfit(prices)
	if ret != 3 {
		t.Errorf("最佳买卖股票时机含冷冻期结果错误: got(%d), expecet(%d)", ret, 3)
	}
}

func Test_findLength(t *testing.T) {
	a := []int{1, 2, 3, 2, 1}
	b := []int{3, 2, 1, 4, 7}
	ret := findLength(a, b)
	if ret != 3 {
		t.Errorf("最长重复子数组结果错误: got(%d), expect(%d)", ret, 3)
	}
}

func Test_intersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	ret := intersect(nums1, nums2)
	if !reflect.DeepEqual(ret, []int{2, 2}) {
		t.Errorf("两个数组的交集II结果错误: got(%v), expect(%v)", ret, []int{2, 2})
	}
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	ret = intersect(nums1, nums2)
	if !reflect.DeepEqual(ret, []int{9, 4}) {
		t.Errorf("两个数组的交集II结果错误: got(%v), expect(%v)", ret, []int{9, 4})
	}
}

func Test_numTrees(t *testing.T) {
	n := 3
	ret := numTrees(n)
	if ret != 5 {
		t.Errorf("不同的二叉搜索数结果错误: got(%d), expect(%d)", ret, 5)
	}
}
