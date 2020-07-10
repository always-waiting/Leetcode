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
