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

func Test_countBinarySubstrings(t *testing.T) {
	s := "00110"
	t.Log(countBinarySubstrings(s))
}

func Test_solve(t *testing.T) {
	board := [][]byte{
		[]byte("OXXOX"),
		[]byte("XOOXO"),
		[]byte("XOXOX"),
		[]byte("OXOOO"),
		[]byte("XXOXO"),
	}
	solve(board)
	for _, v := range board {
		t.Log(string(v))
	}
}

func Test_find132pattern(t *testing.T) {
	nums := []int{-1, 3, 2, 0}
	t.Log(find132pattern(nums))
}

func Test_multiply(t *testing.T) {
	//t.Log(cal("123", '5'))
	//t.Log(cal("0", '0'))
	//t.Log(multiply("123", "456"))
	//t.Log(multiply("0", "0"))
	t.Log(multiply("498828660196", "840477629533"))
	//t.Log(cal("43", '7', 1))
	//t.Log(multiply("9133", "0"))
	t.Log(cal("9133", '0', 0))
}

func Test_productExceptSelf(t *testing.T) {
	t.Log(productExceptSelf1([]int{1, 2, 3, 4}))
}

func Test_isEscapePossible(t *testing.T) {
	t.Log(isEscapePossible(
		[][]int{[]int{0, 1}, []int{1, 0}},
		[]int{0, 0},
		[]int{0, 2},
	))
	t.Log(isEscapePossible(
		[][]int{},
		[]int{0, 0},
		[]int{999999, 999999},
	))
}

func Test_isValid(t *testing.T) {
	t.Log(isValid("{[]}"))
}

func Test_combinationSum3(t *testing.T) {
	// [[1,5,9],[1,6,8],[2,4,9],[2,5,8],[2,6,7],[3,4,8],[3,5,7],[4,5,6]]
	t.Log(combinationSum3(3, 15))
}

func Test_validateBinaryTreeNodes(t *testing.T) {
	t.Log(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))                //true
	t.Log(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}))                 //false
	t.Log(validateBinaryTreeNodes(2, []int{1, 0}, []int{-1, -1}))                               //false
	t.Log(validateBinaryTreeNodes(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1})) //false
	t.Log(validateBinaryTreeNodes(2, []int{-1, 0}, []int{-1, -1}))                              //true
	t.Log(validateBinaryTreeNodes(3, []int{1, -1, 0}, []int{-1, -1, -1}))                       //true
	t.Log(validateBinaryTreeNodes(4, []int{1, 2, 0, -1}, []int{-1, -1, -1, -1}))                //false
}

func Test_addNegabinary(t *testing.T) {
	t.Log([]int{1, 0, 0, 0, 0}, addNegabinary([]int{1, 1, 1, 1, 1}, []int{1, 0, 1}))
	t.Log([]int{1, 1, 0, 1, 0}, addNegabinary([]int{1}, []int{1, 0, 1}))
	t.Log([]int{1, 1, 0}, addNegabinary([]int{1}, []int{1}))
	t.Log([]int{0}, addNegabinary([]int{1, 1}, []int{1}))
	t.Log([]int{1, 1, 0}, addNegabinary([]int{1, 1, 1}, []int{1, 1}))
	t.Log([]int{1, 1, 0, 1, 0}, addNegabinary([]int{1, 1, 1}, []int{1, 1, 1}))
	t.Log([]int{1, 1, 0, 1}, addNegabinary([]int{1, 1, 1, 0}, []int{1, 1, 1}))
	t.Log([]int{1, 0}, addNegabinary([]int{1, 1, 1, 1}, []int{1, 1, 1}))
	t.Log([]int{1, 1, 0, 1, 0, 1}, addNegabinary([]int{1, 1, 1, 1}, []int{1, 1, 1, 0}))
	t.Log([]int{1, 1, 0, 1}, addNegabinary([]int{1, 1}, []int{1, 0}))
	t.Log([]int{1, 0, 1, 0}, addNegabinary([]int{1, 1, 1, 1}, []int{1, 1, 1, 1}))
}

func Test_maxScore(t *testing.T) {
	t.Log(maxScore("011101"))
	t.Log(maxScore("00111"))
	t.Log(maxScore("1111"))
	t.Log(maxScore("00"))

}

func Test_accountsMerge(t *testing.T) {
	t.Log(accountsMerge([][]string{
		[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
		[]string{"John", "johnnybravo@mail.com"},
		[]string{"Mary", "mary@mail.com"},
	}))
}
