package main

import (
	"fmt"
	"math"
	"sort"
)

func test() {
	fmt.Println("vim-go")
}

/*
1. 有效的山脉数组	--	https://leetcode-cn.com/problems/valid-mountain-array/
2. 插入区间			--	https://leetcode-cn.com/problems/insert-interval/
3. 单次接龙			--	https://leetcode-cn.com/problems/word-ladder/
4. 根据数字二进制下1的数目排序	--	https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/
5. 最接近原点的 K 个点	--	https://leetcode-cn.com/problems/k-closest-points-to-origin/
6. 下一个排列			--	https://leetcode-cn.com/problems/next-permutation/
7. 自由之路				--	https://leetcode-cn.com/problems/freedom-trail/
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

/*
1356. 根据数字二进制下 1 的数目排序
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。
如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。
请你返回排序后的数组。

示例 1：
输入：arr = [0,1,2,3,4,5,6,7,8]
输出：[0,1,2,4,8,3,5,6,7]
解释：[0] 是唯一一个有 0 个 1 的数。
[1,2,4,8] 都有 1 个 1 。
[3,5,6] 有 2 个 1 。
[7] 有 3 个 1 。
按照 1 的个数排序得到的结果数组为 [0,1,2,4,8,3,5,6,7]
示例 2：
输入：arr = [1024,512,256,128,64,32,16,8,4,2,1]
输出：[1,2,4,8,16,32,64,128,256,512,1024]
解释：数组中所有整数二进制下都只有 1 个 1 ，所以你需要按照数值大小将它们排序。
示例 3：
输入：arr = [10000,10000]
输出：[10000,10000]
示例 4：
输入：arr = [2,3,5,7,11,13,17,19]
输出：[2,3,5,17,7,11,13,19]
示例 5：
输入：arr = [10,100,1000,10000]
输出：[10,100,10000,1000]

提示：
1 <= arr.length <= 500
0 <= arr[i] <= 10^4
*/
func sortByBits(arr []int) []int {
	return nil
}

/*
973. 最接近原点的 K 个点
我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
（这里，平面上两点之间的距离是欧几里德距离。）
你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

示例 1：
输入：points = [[1,3],[-2,2]], K = 1
输出：[[-2,2]]
解释：
(1, 3) 和原点之间的距离为 sqrt(10)，
(-2, 2) 和原点之间的距离为 sqrt(8)，
由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
示例 2：
输入：points = [[3,3],[5,-1],[-2,4]], K = 2
输出：[[3,3],[-2,4]]
（答案 [[-2,4],[3,3]] 也会被接受。）

提示：
1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000
*/
type kNode struct {
	next *kNode
	pos  []int
	dis  int
}

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:K]
}

/*
31. 下一个排列
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

/*
514. 自由之路
视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。
给定一个字符串 ring，表示刻在外环上的编码；给定另一个字符串 key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
最初，ring 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
旋转 ring 拼出 key 字符 key[i] 的阶段中：
您可以将 ring 顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
示例：
输入: ring = "godding", key = "gd"
输出: 4
解释:
 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 当然, 我们还需要1步进行拼写。
 因此最终的输出是 4。
提示：
ring 和 key 的字符串长度取值范围均为 1 至 100；
两个字符串中都只有小写字符，并且均可能存在重复字符；
字符串 key 一定可以由字符串 ring 旋转拼出。
*/
func findRotateSteps(ring string, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return min(dp[m-1]...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
