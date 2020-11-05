package main

import (
	"fmt"
	"math"
)

func test() {
	fmt.Println("vim-go")
}

/*
1. 有效的山脉数组	--	https://leetcode-cn.com/problems/valid-mountain-array/
2. 插入区间			--	https://leetcode-cn.com/problems/insert-interval/
3. 单次接龙			--	https://leetcode-cn.com/problems/word-ladder/
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

/*
127. 单词接龙
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:
如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]
输出: 5
解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
示例 2:
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]
输出: 0
解释: endWord "cog" 不在字典中，所以无法进行转换。
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}
