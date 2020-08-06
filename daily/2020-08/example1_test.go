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

func Test_isPalindrome(t *testing.T) {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	pairs := palindromePairs(words)
	t.Log(pairs)
	words = []string{"bat", "tab", "cat"}
	pairs = palindromePairs(words)
	t.Log(pairs)
	words = []string{"a", ""}
	pairs = palindromePairs(words)
	t.Log(pairs)
	words = []string{"a", "abc", "aba", ""}
	pairs = palindromePairs(words)
	t.Log(pairs)
}

func Test_f(t *testing.T) {
	a := "abcd"
	tree = []Node{Node{[26]int{}, -1}}
	insert(a, 0)
	insert("bcd", 1)
	insert("aab", 2)
	for _, node := range tree {
		t.Log(node)
	}
	t.Log(findWord("dcba", 0, 3))
}
