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
