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

func Test_isBipartite(t *testing.T) {
	graph := [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}
	ret := isBipartite(graph)
	if !ret {
		t.Errorf("判断二分图错误, got(false), expect(true): input %v", graph)
	}
	graph = [][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}}
	ret = isBipartite(graph)
	if ret {
		t.Errorf("判断二分图错误, got(true), expect(false): input %v", graph)
	}
	graph = [][]int{[]int{1}, []int{0}, []int{4}, []int{4}, []int{2, 3}}
	ret = isBipartite(graph)
	if !ret {
		t.Errorf("判断二分图错误, got(false), expect(true): input %v", graph)
	}

}

func Test_isBipartite1(t *testing.T) {
	graph := [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}
	ret := isBipartite1(graph)
	if !ret {
		t.Errorf("判断二分图错误, got(false), expect(true): input %v", graph)
	}
	graph = [][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}}
	ret = isBipartite1(graph)
	if ret {
		t.Errorf("判断二分图错误, got(true), expect(false): input %v", graph)
	}
	graph = [][]int{[]int{1}, []int{0}, []int{4}, []int{4}, []int{2, 3}}
	ret = isBipartite1(graph)
	if !ret {
		t.Errorf("判断二分图错误, got(false), expect(true): input %v", graph)
	}

}

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 40
	exp := 2
	ret := searchInsert(nums, target)
	if ret != exp {
		t.Errorf("搜索插入位置结果错误, got(%d), expect(%d)", ret, exp)
	}
}

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	ret := minimumTotal(triangle)
	exp := 11
	if ret != exp {
		t.Errorf("三角形最小路径和结果错误, got(%d), expect(%d)", ret, exp)
	}
	ret = minimumTotal1(triangle)
	if ret != exp {
		t.Errorf("三角形最小路径和结果错误, got(%d), expect(%d)", ret, exp)
	}
}

func Test_twoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(numbers, target)
	if !reflect.DeepEqual(ret, []int{1, 2}) {
		t.Errorf("两数之和 II - 输入有序数组结果错误, got(%v), expect(%v)", ret, []int{1, 2})
	}
}

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	ret := minPathSum(grid)
	exp := 7
	if ret != exp {
		t.Errorf("最小路径和结果错误: got(%d), expect(%d)", ret, exp)
	}
}

func Test_isInterleave(t *testing.T) {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	if !isInterleave(s1, s2, s3) {
		t.Errorf("交错字符串结果错误: got(false), expect(true)")
	}
}
