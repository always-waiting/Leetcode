package main

import (
	"math"
)

/*
1. 跳水板
2. 恢复空格		--	https://leetcode-cn.com/problems/re-space-lcci/solution/hui-fu-kong-ge-by-leetcode-solution/
*/

/*
你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，
长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。

返回的长度需要从小到大排列。

示例：

输入：
shorter = 1
longer = 2
k = 3
输出： {3,4,5,6}
提示：

0 < shorter <= longer
0 <= k <= 100000
*/
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	ret := make([]int, k+1)
	for i := k; i >= 0; i-- {
		l := i*shorter + (k-i)*longer
		ret[k-i] = l
	}
	return ret
}

/*
哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。
像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。
在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。
假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。
注意：本题相对原题稍作改动，只需返回未识别的字符数

示例：
输入：
dictionary = ["looked","just","like","her","brother"]
sentence = "jesslookedjustliketimherbrother"
输出： 7
解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。

提示：
0 <= len(sentence) <= 1000
dictionary中总字符数不超过 150000。
你可以认为dictionary和sentence中只包含小写字母。

思路:
方法一：Trie + 动态规划
方法二：字符串哈希 + 动态规划
*/

// 方法一：Trie + 动态规划
func respace(dictionary []string, sentence string) int {
	n, inf := len(sentence), 0x3f3f3f3f
	root := &Trie{next: [26]*Trie{}}
	for _, word := range dictionary {
		root.insert(word)
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		curPos := root
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1] - 'a')
			if curPos.next[t] == nil {
				break
			} else if curPos.next[t].isEnd {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			curPos = curPos.next[t]
		}
	}
	return dp[n]
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (this *Trie) insert(s string) {
	curPos := this
	for i := len(s) - 1; i >= 0; i-- {
		t := int(s[i] - 'a')
		if curPos.next[t] == nil {
			curPos.next[t] = &Trie{next: [26]*Trie{}}
		}
		curPos = curPos.next[t]
	}
	curPos.isEnd = true
}

const (
	P    = math.MaxInt32
	BASE = 41
)

//方法二：字符串哈希 + 动态规划
func respace2(dictionary []string, sentence string) int {
	hashValues := map[int]bool{}
	for _, word := range dictionary {
		hashValues[getHash(word)] = true
	}
	f := make([]int, len(sentence)+1)
	for i := 1; i < len(f); i++ {
		f[i] = len(sentence)
	}
	for i := 1; i <= len(sentence); i++ {
		f[i] = f[i-1] + 1
		hashValue := 0
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1]-'a') + 1
			hashValue = (hashValue*BASE + t) % P
			if hashValues[hashValue] {
				f[i] = min(f[i], f[j-1])
			}
		}
	}
	return f[len(sentence)]
}

func getHash(s string) int {
	hashValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		hashValue = (hashValue*BASE + int(s[i]-'a') + 1) % P
	}
	return hashValue
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
