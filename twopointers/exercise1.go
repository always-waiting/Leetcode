package twopoints

import (
	"fmt"
)

/*
Contents:
1. 反转字符串[reverseString]
2. K个不同整数的子数组[subarraysWithKDistinct]	--	★★★★★
*/

/*
反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]

示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	if end == start || end < 0 {
		return
	}
	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		start++
		end--
	}
}

/*
K个不同整数的子数组
给定一个正整数数组 A，如果A的某个子数组中不同整数的个数恰好为K，则称A的这个连续、不一定独立的子数组为好子数组。
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。

示例 1：
输入：A = [1,2,1,2,3], K = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].

示例 2：
输入：A = [1,2,1,3,4], K = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].

提示：
1 <= A.length <= 20000
1 <= A[i] <= A.length
1 <= K <= A.length
*/
/*
官方思路: https://leetcode-cn.com/problems/subarrays-with-k-different-integers/solution/k-ge-bu-tong-zheng-shu-de-zi-shu-zu-by-leetcode/
// 双循环时间复杂度在O(n*n), 导致了超时
*/
func subarraysWithKDistinct(A []int, K int) int {
	var ret int
	if len(A) < K {
		return ret
	}
	w1 := newWindow()
	w2 := newWindow()
	var l1, l2 int
	for _, val := range A {
		w1.add(val)
		w2.add(val)
		for w1.count > K {
			w1.remove(A[l1])
			l1++
		}
		for w2.count >= K {
			w2.remove(A[l2])
			l2++
		}
		ret += l2 - l1
	}
	return ret
}

type window struct {
	cache map[int]int
	count int
}

func newWindow() *window {
	ret := &window{}
	ret.cache = map[int]int{}
	return ret
}

func (this *window) add(i int) {
	if _, ok := this.cache[i]; !ok {
		this.cache[i] = 1
		this.count++
	} else {
		if this.cache[i] == 0 {
			this.count++
		}
		this.cache[i]++
	}
}

func (this *window) remove(i int) {
	if _, ok := this.cache[i]; ok {
		this.cache[i]--
		if this.cache[i] == 0 {
			this.count--
		}
	}
}
