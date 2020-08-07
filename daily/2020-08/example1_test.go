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

func Test_TreeTraversal(t *testing.T) {
	p := &TreeNode{0, nil, nil}
	p.Left = &TreeNode{1, nil, nil}
	p.Right = &TreeNode{2, nil, nil}
	ret := preorderTraversal(p)
	t.Log(ret)

	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preOrder, inOrder)

	inOrder1 := inorderTraversal(tree)
	t.Log(inOrder)
	t.Log(inOrder1)
}

func Test_smallestRange(t *testing.T) {
	nums := [][]int{
		[]int{4, 10, 15, 24, 26},
		[]int{0, 9, 12, 20},
		[]int{5, 18, 22, 30},
	}
	ret := smallestRange(nums)
	t.Log(ret)
}
