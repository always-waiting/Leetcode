package main

import (
	"fmt"
)

func test() {
	fmt.Println("vim-go")
}

/*
1. 有效的山脉数组	--	https://leetcode-cn.com/problems/valid-mountain-array/
2. 插入区间			--	https://leetcode-cn.com/problems/insert-interval/
*/

/*
941. 有效的山脉数组
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]

示例 1：
输入：[2,1]
输出：false
示例 2：
输入：[3,5,5]
输出：false
示例 3：
输入：[0,3,2,1]
输出：true

提示：
0 <= A.length <= 10000
0 <= A[i] <= 10000
*/

func validMountainArray(A []int) bool {
	if len(A) < 2 {
		return false
	}
	if A[0] > A[1] {
		return false
	}
	up := true
	ret := true
	for i := 0; i < len(A)-1; i++ {
		if up {
			if A[i+1] < A[i] {
				up = false
			}
		} else {
			if A[i] <= A[i+1] {
				ret = false
			}
		}
	}
	if up == true {
		ret = false
	}
	return ret
}

/*
57. 插入区间
给出一个无重叠的，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：
输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	ret := [][]int{}
	cover := [][]int{newInterval}
	left := [][]int{}
	right := [][]int{}
	for _, inter := range intervals {
		flag := isCover(inter, newInterval)
		switch flag {
		case -1:
			left = append(left, inter)
		case 1:
			right = append(right, inter)
		case 0:
			cover = append(cover, inter)
		}
	}
	if len(left) != 0 {
		ret = append(ret, left...)
	}
	if len(cover) != 0 {
		mergeInter := merge(cover)
		ret = append(ret, mergeInter)
	}
	if len(right) != 0 {
		ret = append(ret, right...)
	}
	return ret
}

// 把重合的区间合并
func merge(cover [][]int) []int {
	min, max := cover[0][0], cover[0][1]
	for _, val := range cover {
		if min > val[0] {
			min = val[0]
		}
		if max < val[1] {
			max = val[1]
		}
	}
	return []int{min, max}
}

// -1: a在b左; 0: a,b重叠; 1: a在b右
func isCover(a []int, b []int) int {
	if a[1] < b[0] {
		return -1
	} else if a[0] > b[1] {
		return 1
	}
	return 0
}
