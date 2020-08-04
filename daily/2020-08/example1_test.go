package main

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	num1 := "92"
	num2 := "19"
	ret := addStrings(num1, num2)
	if ret != "111" {
		t.Errorf("字符串相加结果错误: got(%s), expect(%s)", ret, "111")
	}

}

func Test_canFinish(t *testing.T) {
	ret := canFinish(2, [][]int{[]int{1, 0}})
	t.Log(ret)
	if !ret {
		t.Errorf("课程表I结果错误: got(%v), expect(true)", ret)
	}
}
