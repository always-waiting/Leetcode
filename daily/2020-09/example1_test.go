package main

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	//t.Log("+100", isNumber("+100"))
	//t.Log("5e2", isNumber("5e2"))
	//t.Log("-123", isNumber("-123"))
	//t.Log("3.1415", isNumber("3.1415"))
	//t.Log("-1E-16", isNumber("-1E-16"))
	//t.Log("0123", isNumber("0123"))
	t.Log("1", isNumber("1"))
	t.Log("3.", isNumber("3."))
	t.Log("12e", isNumber("12e"))
	t.Log("1a3.14", isNumber("1a3.14"))
	t.Log("1.2.3", isNumber("1.2.3"))
	t.Log("+-5", isNumber("+-5"))
	t.Log("12e5.4", isNumber("12e+5.4"))
}

func Test_topKFrequent(t *testing.T) {
	//t.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	//t.Log(topKFrequent([]int{1}, 1))
	t.Log(topKFrequent1([]int{1, 1, 1, 2, 2, 3}, 2))
}

func Test_combinationSum2(t *testing.T) {
	t.Log(combinationSum2([]int{1, 1, 1, 3, 3, 5}, 8))
	// [[1 7] [2 6] [7 1] [1 6 1] [1 2 5] [2 1 5]]
	//t.Log(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func Test_buildTree(t *testing.T) {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	t.Log(root.Right.Left)
}

func Test_longestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("abcba"))
}
