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
