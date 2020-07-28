package main

import (
	"fmt"
	"math"
)

/*
1. 跳水板
2. 恢复空格		--	https://leetcode-cn.com/problems/re-space-lcci/solution/hui-fu-kong-ge-by-leetcode-solution/
3. 最佳买卖股票时机含冷冻期		--	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/
4. 最长重复子数组		--		https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/
5. 两个数组的交集II
6. 不同的二叉搜索数		--		https://leetcode-cn.com/problems/unique-binary-search-trees/
7. 判断二分图
8. 搜索插入位置
9. 三角形最小路径和
10. 两数之和 II - 输入有序数组
11. 不同的二叉搜索树 II
12. 最小路径和
13. 交错字符串
14. 除数博弈
15. 二叉树的最大深度
*/

/*
1. 跳水板
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
2. 恢复空格		--	https://leetcode-cn.com/problems/re-space-lcci/solution/hui-fu-kong-ge-by-leetcode-solution/
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

/*
3. 最佳买卖股票时机含冷冻期		--	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/
给定一个整数数组，其中第i个元素代表了第i天的股票价格 。
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为1天)。

示例:
输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
4. 最长重复子数组		--		https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/
给两个整数数组A和B，返回两个数组中公共的、长度最长的子数组的长度。

示例：
输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。

提示：
1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
*/
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

/*
5. 两个数组的交集II
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]

示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]

说明：
输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。

进阶:
如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	for _, val := range nums1 {
		if _, ok := cache[val]; ok {
			cache[val]++
		} else {
			cache[val] = 1
		}
	}
	ret := make([]int, 0)
	for _, val := range nums2 {
		if _, ok := cache[val]; ok {
			cache[val]--
			if cache[val] >= 0 {
				ret = append(ret, val)
			}
		}
	}
	return ret
}

/*
6. 不同的二叉搜索数
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] = G[i] + G[j-1]*G[i-j]
		}
	}
	return G[n]
}

/*
7. 判断二分图
给定一个无向图graph，当这个图为二分图时返回true。
如果我们能将一个图的节点集合分割成两个独立的子集A和B，
并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。
每个节点都是一个在0到graph.length-1之间的整数。
这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。

示例 1:
输入: [[1,3], [0,2], [1,3], [0,2]]
输出: true
解释:
无向图如下:
0----1
|    |
|    |
3----2
我们可以将节点分成两组: {0, 2} 和 {1, 3}。

示例 2:
输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
输出: false
解释:
无向图如下:
0----1
| \  |
|  \ |
3----2
我们不能将节点分割成两个独立的子集。

注意:
graph 的长度范围为 [1, 100]。
graph[i] 中的元素的范围为 [0, graph.length - 1]。
graph[i] 不会包含 i 或者有重复的值。
图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。
*/
func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}
	n := len(graph)
	queue := createQueueBipartite(n)
	for queue.idx != len(graph)-1 {
		now := queue.get()
		if now == nil {
			queue.add(&bipartite{
				set: 0,
				val: queue.startIdx(),
			})
			continue
		}
		if !queue.set(now) {
			return false
		}
		links := graph[now.val]
		for _, val := range links {
			next := &bipartite{val: val, set: (now.set + 1) % 2}
			queue.add(next)
			if !queue.set(next) {
				return false
			}
		}
	}

	return true
}

type queueBipartite struct {
	data []*bipartite
	idx  int
	a0   map[int]bool
	a1   map[int]bool
	nums []int
}

func createQueueBipartite(n int) (ret queueBipartite) {
	ret = queueBipartite{}
	ret.data = make([]*bipartite, 0)
	ret.nums = make([]int, n)
	ret.a0 = make(map[int]bool)
	ret.a1 = make(map[int]bool)
	return
}

func (this *queueBipartite) startIdx() int {
	var ret int
	for i, val := range this.nums {
		if val == 0 {
			ret = i
			break
		}
	}
	return ret
}

func (this *queueBipartite) set(obj *bipartite) bool {
	if obj.set == 0 {
		if _, ok := this.a1[obj.val]; ok {
			return false
		}
		this.a0[obj.val] = true
	} else {
		if _, ok := this.a0[obj.val]; ok {
			return false
		}
		this.a1[obj.val] = true
	}
	return true
}

func (this *queueBipartite) get() *bipartite {
	if this.idx >= len(this.data) {
		return nil
	}
	ret := this.data[this.idx]
	this.idx++
	return ret
}

func (this *queueBipartite) add(obj *bipartite) {
	if _, ok := this.a0[obj.val]; ok {
		return
	}
	if _, ok := this.a1[obj.val]; ok {
		return
	}
	this.nums[obj.val] = 1
	this.data = append(this.data, obj)
}

type bipartite struct {
	val int
	set int
}

// 广度优先搜索
const (
	UNCOLOR, RED, GREEN = 0, 1, 2
)

func isBipartite1(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == UNCOLOR {
			queue := []int{i}
			color[i] = RED
			for i := 0; i < len(queue); i++ {
				node := queue[i]
				cNei := RED
				if color[node] == RED {
					cNei = GREEN
				}
				for _, neighbor := range graph[node] {
					if color[neighbor] == UNCOLOR {
						queue = append(queue, neighbor)
						color[neighbor] = cNei
					} else {
						if color[neighbor] != cNei {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/
// 二分查找法
func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	ans := n
	for l <= r {
		mid := (r-l)>>1 + l
		if target <= nums[mid] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

/*
120. 三角形最小路径和
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。*
*/
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	f[0] = []int{triangle[0][0]}
	for i := 1; i < n; i++ {
		f[i] = make([]int, i+1)
		f[i][0] = f[i-1][0] + triangle[i][0]
		f[i][i] = f[i-1][i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			f[i][j] = minArray([]int{f[i-1][j-1], f[i-1][j]}) + triangle[i][j]
		}
	}
	return minArray(f[n-1])
}

func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	f := [2][]int{}
	for i := 0; i < 2; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		curr := i % 2
		prev := 1 - curr
		f[curr][0] = f[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[curr][j] = min(f[prev][j-1], f[prev][j]) + triangle[i][j]
		}
		f[curr][i] = f[prev][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[(n-1)%2][i])
	}
	return ans
}

func minArray(a []int) int {
	ret := a[0]
	for i := 1; i < len(a); i++ {
		if ret > a[i] {
			ret = a[i]
		}
	}
	return ret
}

/*
10. 两数之和 II - 输入有序数组
给定一个已按照升序排列的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值index1和index2，其中index1必须小于index2。

说明:
返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

示例:
输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
*/
func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for start < end {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}

/*
95. 不同的二叉搜索树 II
给定一个整数n，生成所有由1 ... n为节点所组成的二叉搜索树。

示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
提示：
0 <= n <= 8
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return nil
	}
	allTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{Val: i}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)

			}
		}
	}
	return allTrees
}

/*
12. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0
// 二分查找法
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high - low) / 2
        if numbers[pivot] < numbers[high] {
            high = pivot
        } else if numbers[pivot] > numbers[high] {
            low = pivot + 1
        } else {
            high--
        }
	}
	return numbers[low]
}
*/
func minArray_question(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	ret := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if ret > numbers[i] {
			ret = numbers[i]
			break
		}
	}
	return ret
}

/*
12. 最小路径和
给定一个包含非负整数的 m * n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func minPathSum(grid [][]int) int {
	a := len(grid)
	if a == 0 {
		return 0
	}
	b := len(grid[0])
	if b == 0 {
		return 0
	}
	sum := make([][]int, a)
	sum[0] = make([]int, b)
	sum[0][0] = grid[0][0]
	for i := 1; i < a; i++ {
		sum[i] = make([]int, b)
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for j := 1; j < b; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j]
	}
	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			sa := sum[i][j-1] + grid[i][j]
			sb := sum[i-1][j] + grid[i][j]
			sum[i][j] = min(sa, sb)
		}
	}
	return sum[a-1][b-1]
}

/*
13. 交错字符串
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1：
输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true

示例 2：
输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	i, j, k := 0, 0, 0
	for k < len(s3) {
		if i < len(s1) && s1[i] == s3[k] {
			i++
			k++
			continue
		}
		if j < len(s2) && s2[j] == s3[k] {
			j++
			k++
			continue
		}
		fmt.Println(i, j, k)
		return false
	}
	return true
}

/*
14. 除数博弈
爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
最初，黑板上有一个数字N 。在每个玩家的回合，玩家需要执行以下操作：
选出任一x，满足 0 < x < N 且 N % x == 0 。
用N - x替换黑板上的数字N 。
如果玩家无法执行这些操作，就会输掉游戏。
只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。
假设两个玩家都以最佳状态参与游戏。

示例 1：
输入：2
输出：true
解释：爱丽丝选择 1，鲍勃无法进行操作。

示例 2：
输入：3
输出：false
解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。

提示：
1 <= N <= 1000
*/
func divisorGame(N int) bool {
	return N%2 == 0
}

/*
15. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	} else if root.Right == nil {
		return maxDepth(root.Left) + 1
	} else if root.Left == nil {
		return maxDepth(root.Right) + 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
