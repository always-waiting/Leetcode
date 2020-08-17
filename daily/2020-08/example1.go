package main

import "fmt"
import "math"

func main() {
	fmt.Println("vim-go")
	a := '3'
	b := '9'
	fmt.Println(a - '0' + b)
}

/*
1. 字符串相加	--	https://leetcode-cn.com/problems/add-strings/
2. 课程表	--	https://leetcode-cn.com/problems/course-schedule/
3. 回文对	--	https://leetcode-cn.com/problems/palindrome-pairs/
4. 相同的树	--	https://leetcode-cn.com/problems/same-tree/
5. 二叉树展开为链表	--	https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
6. 最小区间	--	https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/
7. 计数二进制子串	--	https://leetcode-cn.com/problems/count-binary-substrings/
8. 被围绕的区域	--	https://leetcode-cn.com/problems/surrounded-regions/
9. 克隆图	--	https://leetcode-cn.com/problems/clone-graph/
10. 滑动窗口的最大值I	--	https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
11. 132模式	--	https://leetcode-cn.com/problems/132-pattern/
12. 字符串相乘	--	https://leetcode-cn.com/problems/multiply-strings/
13. 除自身以外数组的乘积	--	https://leetcode-cn.com/problems/product-of-array-except-self/
14. 二叉树的右视图	--	https://leetcode-cn.com/problems/binary-tree-right-side-view/
15. 逃离大迷宫	--	https://leetcode-cn.com/problems/escape-a-large-maze/
16. 有效的括号	--	https://leetcode-cn.com/problems/valid-parentheses/
17. 组合总和 III	--	https://leetcode-cn.com/problems/combination-sum-iii/
18. 二叉搜索树中第K小的元素	--	https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
19. 验证二叉树	--	https://leetcode-cn.com/problems/validate-binary-tree-nodes/
20. 负二进制数相加	--	https://leetcode-cn.com/problems/adding-two-negabinary-numbers/
21. 移除盒子	--	https://leetcode-cn.com/problems/remove-boxes/
22. 平衡二叉树	--	https://leetcode-cn.com/problems/balanced-binary-tree/
23. 分割字符串的最大得分	--	https://leetcode-cn.com/problems/maximum-score-after-splitting-a-string/
*/

/*
1. 字符串相加
给定两个字符串形式的非负整数num1和num2 ，计算它们的和。

注意：
num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/

func addStrings(num1 string, num2 string) string {
	var length int
	idx1 := len(num1) - 1
	idx2 := len(num2) - 1
	if idx1 > idx2 {
		length = idx1 + 1
	} else {
		length = idx2 + 1
	}
	ret := make([]byte, length+1)
	carry := byte('0')
	for idx1 >= 0 || idx2 >= 0 {
		tmp := carry
		if idx1 >= 0 {
			tmp = tmp + num1[idx1] - '0'
		}
		if idx2 >= 0 {
			tmp = tmp + num2[idx2] - '0'
		}
		if tmp > byte('9') {
			ret[length] = tmp - (byte('9') - byte('0')) - 1
			carry = byte('1')
		} else {
			ret[length] = tmp
			carry = byte('0')
		}
		idx1--
		idx2--
		length--
	}
	if carry > byte('0') {
		ret[length] = carry
	} else {
		ret = ret[1:]
	}
	return string(ret)
}

/*
2. 课程表
你这个学期必须选修numCourse门课程，记为0到numCourse-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程0 ，你需要先完成课程1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。

示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有2门课程。学习课程1之前，你需要先完成课程0；并且学习课程0之前，你还应先完成课程1。这是不可能的。

提示：
输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	if len(prerequisites) == 0 {
		return true
	}
	for _, cond := range prerequisites {
		degree[cond[0]]++
	}
	study := make([]int, 0)
	for c, i := range degree {
		if i == 0 {
			study = append(study, c)
		}
	}
	for len(study) != 0 {
		c := study[0]
		for _, cond := range prerequisites {
			if cond[1] == c {
				degree[cond[0]]--
				if degree[cond[0]] == 0 {
					study = append(study, cond[0])
				}
			}
		}
		study = study[1:]
	}
	for _, i := range degree {
		if i != 0 {
			return false
		}
	}
	return true
}

/*
3. 回文对
给定一组互不相同的单词，找出所有不同的索引对(i, j)，使得列表中的两个单词，words[i] + words[j] ，可拼接成回文串。

示例 1：
输入：["abcd","dcba","lls","s","sssll"]
输出：[[0,1],[1,0],[3,2],[2,4]]
解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]

示例 2：
输入：["bat","tab","cat"]
输出：[[0,1],[1,0]]
解释：可拼接成的回文串为 ["battab","tabbat"]
*/
/*
// 超时!!
func palindromePairs(words []string) [][]int {
	ret := make([][]int, 0)
	for i, word1 := range words {
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if isPalindrome(word1 + word2) {
				ret = append(ret, []int{i, j})
			}
			if isPalindrome(word2 + word1) {
				ret = append(ret, []int{j, i})
			}
		}
	}
	return ret
}

func isPalindrome(in string) bool {
	total := len(in)
	mid := total / 2
	for i := 0; i < mid; i++ {
		if in[i] != in[total-1-i] {
			return false
		}
	}
	return true
}
*/
type Node struct {
	ch   [26]int
	flag int
}

var tree []Node

func palindromePairs(words []string) [][]int {
	tree = []Node{Node{[26]int{}, -1}}
	n := len(words)
	for i := 0; i < n; i++ {
		insert(words[i], i)
	}
	ret := [][]int{}
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func insert(s string, id int) {
	add := 0
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if tree[add].ch[x] == 0 {
			tree = append(tree, Node{[26]int{}, -1})
			tree[add].ch[x] = len(tree) - 1
		}
		add = tree[add].ch[x]
	}
	tree[add].flag = id
}

// 找反向字符串是否在字典中!
func findWord(s string, left, right int) int {
	add := 0
	for i := right; i >= left; i-- {
		x := int(s[i] - 'a')
		if tree[add].ch[x] == 0 {
			return -1
		}
		add = tree[add].ch[x]
	}
	return tree[add].flag
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

/*
4. 相同的树
给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:
输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]
输出: true

示例 2:
输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false

示例 3:
输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]
输出: false
*/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

func preorderTraversal(p *TreeNode) []int {
	if p == nil {
		return nil
	}
	ret := []int{p.Val}
	left := preorderTraversal(p.Left)
	right := preorderTraversal(p.Right)
	if left != nil {
		ret = append(ret, left...)
	}
	if right != nil {
		ret = append(ret, right...)
	}
	return ret
}

func inorderTraversal(p *TreeNode) []int {
	if p == nil {
		return nil
	}
	left := inorderTraversal(p.Left)
	right := inorderTraversal(p.Right)
	ret := []int{}
	if left != nil {
		ret = append(ret, left...)
	}
	ret = append(ret, p.Val)
	if right != nil {
		ret = append(ret, right...)
	}
	return ret
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{val, nil, nil}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}
	leftInOrder := inorder[0:i]
	rightInOrder := inorder[i+1:]
	leftPreOrder := preorder[1 : len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]
	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)
	return root
}

/*
5. 二叉树展开为链表
给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	preorderflatten(root)
}

func preorderflatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := preorderflatten(root.Left)
	right := preorderflatten(root.Right)
	root.Left = nil
	root.Right = left
	if left == nil {
		root.Right = right
		return root
	}
	for left.Right != nil {
		left = left.Right
	}
	left.Right = right
	return root
}

/*
6. 最小区间
你有k个升序排列的整数列表。找到一个最小区间，使得k个列表中的每个列表至少有一个数包含在其中。
我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。

示例：
输入：[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
输出：[20,24]
解释：
列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。

提示：
给定的列表可能包含重复元素，所以在这里升序表示 >= 。
1 <= k <= 3500
-105 <= 元素的值 <= 105
对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。
*/
type heap struct {
	idx    []int
	minIdx int
	maxIdx int
}

func (this *heap) Init(nums [][]int) {
	this.idx = make([]int, len(nums))
	for i, idx := range this.idx {
		if idx >= len(nums[i]) ||
			idx >= len(nums[this.minIdx]) ||
			idx >= len(nums[this.maxIdx]) {
			break
		}
		if nums[i][idx] < nums[this.minIdx][idx] {
			this.minIdx = i
		}
		if nums[i][idx] > nums[this.maxIdx][idx] {
			this.maxIdx = i
		}
	}
}

func (this *heap) cal(nums [][]int) bool {
	for i, idx := range this.idx {
		if idx >= len(nums[i]) ||
			this.idx[this.minIdx] >= len(nums[this.minIdx]) ||
			this.idx[this.maxIdx] >= len(nums[this.maxIdx]) {
			return false
		}
		if nums[i][idx] < nums[this.minIdx][this.idx[this.minIdx]] {
			this.minIdx = i
		}
		if nums[i][idx] > nums[this.maxIdx][this.idx[this.maxIdx]] {
			this.maxIdx = i
		}
	}
	return true
}

func (this *heap) getRange(nums [][]int) []int {
	return []int{nums[this.minIdx][this.idx[this.minIdx]], nums[this.maxIdx][this.idx[this.maxIdx]]}
}

func (this *heap) step(nums [][]int) bool {
	this.idx[this.minIdx]++
	return this.cal(nums)
}

func smallestRange(nums [][]int) []int {
	h := &heap{}
	h.Init(nums)
	var ret []int
	for {
		if ret == nil {
			ret = h.getRange(nums)
			if !h.step(nums) {
				break
			}
			continue
		}
		tmp := h.getRange(nums)
		if ret[1]-ret[0] > tmp[1]-tmp[0] {
			ret = tmp
		}
		if !h.step(nums) {
			break
		}
	}
	return ret
}

/*
7. 计数二进制子串
给定一个字符串s，计算具有相同数量0和1的非空(连续)子字符串的数量，
并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。
示例 1 :
输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
请注意，一些重复出现的子串要计算它们出现的次数。
另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。

示例 2 :
输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
注意：

s.length 在1到50,000之间。
s 只包含“0”或“1”字符。
*/
func countBinarySubstrings(s string) int {
	counts := make([]int, 0)
	cs := s[0]
	c := 1
	for i := 1; i < len(s); i++ {
		if cs == s[i] {
			c++
		} else {
			cs = s[i]
			counts = append(counts, c)
			c = 1
		}
	}
	counts = append(counts, c)
	var ret int
	for i := 1; i < len(counts); i++ {
		ret = ret + min(counts[i-1], counts[i])
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
8. 被围绕的区域
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:
X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：
X X X X
X X X X
X X X X
X O X X
解释:
被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/
var n, m int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}

/*
给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）。
图中的每个节点都包含它的值 val（int）和其邻居的列表（list[Node]）。
class Node {
    public int val;
    public List<Node> neighbors;
}

测试用例格式：
简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。
该图在测试用例中使用邻接列表表示。邻接列表是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
给定节点将始终是图中的第一个节点（值为 1）。你必须将给定节点的拷贝作为对克隆图的引用返回。

示例 1：
输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
输出：[[2,4],[1,3],[2,4],[1,3]]
解释：
图中有 4 个节点。
节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
节点 4 的值是 4，它有两个邻居：节点 1 和 3 。

示例 2：
输入：adjList = [[]]
输出：[[]]
解释：输入包含一个空列表。该图仅仅只有一个值为 1 的节点，它没有任何邻居。

示例 3：
输入：adjList = []
输出：[]
解释：这个图是空的，它不含任何节点。

示例 4：
输入：adjList = [[2],[1]]
输出：[[2],[1]]

提示：
节点数不超过 100 。
每个节点值 Node.val 都是唯一的，1 <= Node.val <= 100。
无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p 的邻居。
图是连通图，你可以从给定节点访问到所有节点。
*/
type GNode struct {
	Val       int
	Neighbors []*GNode
}

// 增加一个已经存在的节点列表，避免死循环
func cloneGraph(node *GNode) *GNode {
	visited := map[*GNode]*GNode{}
	var cg func(node *GNode) *GNode
	cg = func(node *GNode) *GNode {
		if node == nil {
			return nil
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		ret := &GNode{node.Val, []*GNode{}}
		visited[node] = ret
		for _, innode := range node.Neighbors {
			ret.Neighbors = append(ret.Neighbors, cg(innode))
		}
		return ret
	}
	return cg(node)
}

/*
10. 滑动窗口的最大值I
给定一个数组nums和滑动窗口的大小k，请找出所有滑动窗口里的最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/
func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	getMax := func(in []int) int {
		ret := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] > ret {
				ret = in[i]
			}
		}
		return ret
	}
	for i := 0; i <= len(nums)-k; i++ {
		ret = append(ret, getMax(nums[i:i+k]))
	}
	return ret
}

/*
11. 132模式
给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：
当 i < j < k 时，ai < ak < aj。设计一个算法，当给定有n个数字的序列时，
验证这个序列中是否含有132模式的子序列。

注意：n 的值小于15000。
示例1:
输入: [1, 2, 3, 4]
输出: False
解释: 序列中不存在132模式的子序列。
示例 2:
输入: [3, 1, 4, 2]
输出: True
解释: 序列中有 1 个132模式的子序列： [1, 4, 2].
示例 3:
输入: [-1, 3, 2, 0]
输出: True
解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].
*/
func find132pattern(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	mins := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mins[i-1] {
			mins[i] = nums[i]
		} else {
			mins[i] = mins[i-1]
		}
	}
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= mins[i] {
			continue
		}
		if len(stack) == 0 {
			stack = append(stack, nums[i])
			continue
		}
		if nums[i] > stack[len(stack)-1] {
			for len(stack) != 0 {
				if stack[len(stack)-1] > mins[i] {
					return true
				}
				stack = stack[0 : len(stack)-1]
			}
		}
		stack = append(stack, nums[i])
	}
	return false
}

/*
12. 字符串相乘
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"

说明：
num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	sum := []string{}
	for i := len(num2) - 1; i >= 0; i-- {
		ele := cal(num1, num2[i], len(num2)-i-1)
		sum = append(sum, ele)
	}
	ret := "0"
	for _, val := range sum {
		ret = add(ret, val)
	}
	return ret
}

func cal(num1 string, i byte, j int) string {
	ret := []string{}
	for k := len(num1) - 1; k >= 0; k-- {
		loop := int(i) - '0'
		tmp := "0"
		for loop != 0 {
			tmp = add(tmp, string(num1[k]))
			loop--
		}
		loop = j + len(num1) - k - 1
		if tmp == "0" {
			continue
		}
		for loop != 0 {
			tmp = tmp + "0"
			loop--
		}
		ret = append(ret, tmp)
	}
	retStr := "0"
	for i := 0; i < len(ret); i++ {
		retStr = add(retStr, ret[i])
	}
	return retStr
}

func add(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := false
	ret := []byte{}
	for i >= 0 && j >= 0 {
		sum := num1[i] + num2[j] - '0'
		if carry {
			sum = sum + 1
		}
		if sum > '9' {
			carry = true
			ret = append(ret, sum-'9'+'0'-1)
		} else {
			carry = false
			ret = append(ret, sum)
		}
		i--
		j--
	}
	for i >= 0 {
		var sum byte
		if carry {
			sum = num1[i] + 1
		} else {
			sum = num1[i]
		}
		if sum > '9' {
			carry = true
			ret = append(ret, sum-'9'+'0'-1)
		} else {
			carry = false
			ret = append(ret, sum)
		}
		i--
	}
	for j >= 0 {
		var sum byte
		if carry {
			sum = num2[j] + 1
		} else {
			sum = num2[j]
		}
		if sum > '9' {
			carry = true
			ret = append(ret, sum-'9'+'0'-1)
		} else {
			carry = false
			ret = append(ret, sum)
		}
		j--
	}
	if carry {
		ret = append(ret, '1')
	}
	i = 0
	j = len(ret) - 1
	for i < j {
		tmp := ret[i]
		ret[i] = ret[j]
		ret[j] = tmp
		i++
		j--
	}
	return string(ret)
}

/*
// 解题结果，上面的代码测试超时，但是正确
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    ans := "0"
    m, n := len(num1), len(num2)
    for i := n - 1; i >= 0; i-- {
        curr := ""
        add := 0
        for j := n - 1; j > i; j-- {
            curr += "0"
        }
        y := int(num2[i] - '0')
        for j := m - 1; j >= 0; j-- {
            x := int(num1[j] - '0')
            product := x * y + add
            curr = strconv.Itoa(product % 10) + curr
            add = product / 10
        }
        for ; add != 0; add /= 10 {
            curr = strconv.Itoa(add % 10) + curr
        }
        ans = addStrings(ans, curr)
    }
    return ans
}

func addStrings(num1, num2 string) string {
    i, j := len(num1) - 1, len(num2) - 1
    add := 0
    ans := ""
    for ; i >= 0 || j >= 0 || add != 0; i, j = i - 1, j - 1 {
        x, y := 0, 0
        if i >= 0 {
            x = int(num1[i] - '0')
        }
        if j >= 0 {
            y = int(num2[j] - '0')
        }
        result := x + y + add
        ans = strconv.Itoa(result % 10) + ans
        add = result / 10
    }
    return ans
}
*/

/*
13. 除自身以外数组的乘积
给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中output[i]等于nums中除nums[i]之外其余各元素的乘积。

示例:
输入: [1,2,3,4]
输出: [24,12,8,6]

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在O(n)时间复杂度内完成此题。
进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		ret = append(ret, left[i]*right[i])
	}
	return ret
}

func productExceptSelf1(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	tmp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * tmp
		tmp = tmp * nums[i]
	}
	return ret
}

/*
14. 二叉树的右视图
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{root.Val}
	rightRet := rightSideView(root.Right)
	leftRet := rightSideView(root.Left)
	i := 0
	j := 0
	for i < len(rightRet) && j < len(leftRet) {
		ret = append(ret, rightRet[i])
		i++
		j++
	}
	for i < len(rightRet) {
		ret = append(ret, rightRet[i])
		i++
	}
	for j < len(leftRet) {
		ret = append(ret, leftRet[j])
		j++
	}
	return ret
}

/*
15. 逃离大迷宫
在一个 10^6 x 10^6 的网格中，每个网格块的坐标为 (x, y)，其中 0 <= x, y < 10^6。
我们从源方格 source 开始出发，意图赶往目标方格 target。每次移动，我们都可以走到网格中在四个方向上相邻的方格，只要该方格不在给出的封锁列表blocked上。
只有在可以通过一系列的移动到达目标方格时才返回true。否则，返回false。

示例 1：
输入：blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
输出：false
解释：
从源方格无法到达目标方格，因为我们无法在网格中移动。

示例 2：
输入：blocked = [], source = [0,0], target = [999999,999999]
输出：true
解释：
因为没有方格被封锁，所以一定可以到达目标方格。

提示：
0 <= blocked.length <= 200
blocked[i].length == 2
0 <= blocked[i][j] < 10^6
source.length == target.length == 2
0 <= source[i][j], target[i][j] < 10^6
source != target
*/
type pos struct {
	x, y int
}

func makePos(a []int) pos {
	return pos{a[0], a[1]}
}

func (this pos) right() pos {
	return pos{this.x + 1, this.y}
}
func (this pos) left() pos {
	return pos{this.x - 1, this.y}
}
func (this pos) up() pos {
	return pos{this.x, this.y - 1}
}
func (this pos) down() pos {
	return pos{this.x, this.y + 1}
}
func (this pos) arrive(a []int) bool {
	return this.x == a[0] && this.y == a[1]
}
func (this pos) isBlocked(blocked [][]int) bool {
	for _, val := range blocked {
		if this.arrive(val) {
			return true
		}
	}
	return false
}
func (this pos) toString() string {
	return fmt.Sprintf("%d:%d", this.x, this.y)
}
func (this pos) toArray() []int {
	return []int{this.x, this.y}
}

var limit = math.Pow10(6)

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}
	return isEscapePossibleFun(blocked, source, target) && isEscapePossibleFun(blocked, target, source)
}

func isEscapePossibleFun(blocked [][]int, source []int, target []int) bool {
	stack := [][]int{source}
	seen := map[string]bool{}
	for len(stack) != 0 {
		if len(seen) == 20000 {
			return true
		}
		now := makePos(stack[0])
		seen[now.toString()] = true
		right := now.right()
		left := now.left()
		up := now.up()
		down := now.down()
		if right.arrive(target) || left.arrive(target) || up.arrive(target) || down.arrive(target) {
			return true
		}
		if _, ok := seen[right.toString()]; !ok && !right.isBlocked(blocked) && float64(right.x) < limit {
			seen[right.toString()] = true
			stack = append(stack, right.toArray())
		}
		if _, ok := seen[left.toString()]; !ok && !left.isBlocked(blocked) && left.x >= 0 {
			seen[left.toString()] = true
			stack = append(stack, left.toArray())
		}
		if _, ok := seen[up.toString()]; !ok && !up.isBlocked(blocked) && up.y >= 0 {
			seen[up.toString()] = true
			stack = append(stack, up.toArray())
		}
		if _, ok := seen[down.toString()]; !ok && !down.isBlocked(blocked) && float64(down.y) < limit {
			seen[down.toString()] = true
			stack = append(stack, down.toArray())
		}
		stack = stack[1:]
	}
	return false
}

/*
16. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"		输出: true
示例 2:
输入: "()[]{}"	输出: true
示例 3:
输入: "(]"		输出: false
示例 4:
输入: "([)]"	输出: false
示例 5:
输入: "{[]}"	输出: true
*/
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
			continue
		}
		if v == ')' || v == '}' || v == ']' {
			if len(stack) == 0 {
				return false
			}
			pv := stack[len(stack)-1]
			pair := string([]rune{pv, v})
			if pair != "()" && pair != "{}" && pair != "[]" {
				fmt.Println(pair)
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

/*
17. 组合总和 III
找出所有相加之和为n的k个数的组合。组合中只允许含有1 - 9的正整数，并且每种组合中不存在重复的数字。

说明：
所有数字都是正整数。
解集不能包含重复的组合。
示例 1:
输入: k = 3, n = 7		输出: [[1,2,4]]
示例 2:
输入: k = 3, n = 9		输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
func combinationSum3(k int, n int) [][]int {
	return comb(k, n, 0)
}

func comb(k, n, i int) [][]int {
	if k == 1 && n <= 9 {
		return [][]int{[]int{n}}
	}
	if n > 9 && k == 1 {
		return nil
	}
	ret := [][]int{}
	var limit int
	if n%k == 0 {
		limit = n / k
	} else {
		limit = n/k + 1
	}
	for j := i + 1; j < limit; j++ {
		inner := comb(k-1, n-j, j)
		if inner == nil {
			continue
		}
		for _, v := range inner {
			tmp := []int{j}
			tmp = append(tmp, v...)
			ret = append(ret, tmp)
		}
	}
	return ret
}

/*
18. 二叉搜索树中第K小的元素
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设k总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
*/
func kthSmallest(root *TreeNode, k int) int {
	var getMin func(*TreeNode, int) int
	getMin = func(r *TreeNode, i int) int {
		ret := r.Val
		if r.Left != nil {
			var leftMin int
			leftMin = getMin(r.Left, i)
			if leftMin > i && leftMin < ret {
				ret = leftMin
				return ret
			}
		}
		if r.Right != nil {
			var rightMin int
			rightMin = getMin(r.Right, i)
			if ret <= i && rightMin > i {
				ret = rightMin
			}
		}
		return ret
	}
	smallStack := make([]int, 0)
	for len(smallStack) < k {
		upNum := -1
		if len(smallStack) != 0 {
			upNum = smallStack[len(smallStack)-1]
		}
		get := getMin(root, upNum)
		smallStack = append(smallStack, get)
	}
	return smallStack[k-1]
}

/*
19. 验证二叉树
二叉树上有 n 个节点，按从 0 到 n - 1 编号，其中节点 i 的两个子节点分别是 leftChild[i] 和 rightChild[i]。
只有所有节点能够形成且只形成一颗有效的二叉树时，返回true；否则返回false。
如果节点i没有左子节点，那么leftChild[i]就等于-1。右子节点也符合该规则。
注意：节点没有值，本问题中仅仅使用节点编号。

示例 1：
输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
输出：true
示例 2：
输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
输出：false
示例 3：
输入：n = 2, leftChild = [1,0], rightChild = [-1,-1]
输出：false
示例 4：
输入：n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
输出：false

提示：
1 <= n <= 10^4
leftChild.length == rightChild.length == n
-1 <= leftChild[i], rightChild[i] <= n - 1
*/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(leftChild, rightChild)
	if root == -1 {
		return false
	}
	stack := []int{root}
	seen := map[int]bool{root: true}
	for len(stack) != 0 {
		now := stack[0]
		stack = stack[1:]
		lNode := leftChild[now]
		rNode := rightChild[now]
		if lNode != -1 {
			if _, ok := seen[lNode]; ok {
				return false
			} else {
				seen[lNode] = true
				stack = append(stack, lNode)
			}
		}
		if rNode != -1 {
			if _, ok := seen[rNode]; ok {
				return false
			} else {
				seen[rNode] = true
				stack = append(stack, rNode)
			}
		}
	}
	if len(seen) != n {
		return false
	}
	return true
}

func findRoot(leftChild, rightChild []int) int {
	indeg := make([]int, len(leftChild))
	for i := 0; i < len(leftChild); i++ {
		if leftChild[i] != -1 {
			indeg[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			indeg[rightChild[i]]++
		}
	}
	for i, val := range indeg {
		if val == 0 {
			return i
		}
	}
	return -1
}

/*
20. 负二进制数相加
给出基数为-2的两个数arr1和arr2，返回两数相加的结果。
数字以数组形式给出：数组由若干0和1组成，按最高有效位到最低有效位的顺序排列。
例如，arr = [1,1,0,1] 表示数字 (-2)^3 + (-2)^2 + (-2)^0 = -3。数组形式的数字也同样不含前导零：以arr为例，这意味着要么arr == [0]，要么arr[0] == 1。
返回相同表示形式的arr1和arr2相加的结果。两数的表示形式为：不含前导零、由若干0和1组成的数组。

示例：
输入：arr1 = [1,1,1,1,1], arr2 = [1,0,1]
输出：[1,0,0,0,0]
解释：arr1 表示11，arr2表示5，输出表示16 。
提示：
1 <= arr1.length <= 1000
1 <= arr2.length <= 1000
arr1 和 arr2 都不含前导零
arr1[i] 为 0 或 1
arr2[i] 为 0 或 1
*/
func addNegabinary(arr1, arr2 []int) []int {
	maxLen := len(arr1)
	if maxLen < len(arr2) {
		maxLen = len(arr2)
		tmp := arr1
		arr1 = arr2
		arr2 = tmp
	}
	ret := make([]int, maxLen+2)
	i := len(arr1) - 1
	j := len(arr2) - 1
	count := maxLen + 2 - 1
	for i >= 0 && j >= 0 {
		tmp := arr1[i] + arr2[j] + ret[count]
		ret[count] = tmp % 2
		if tmp >= 2 {
			ret[count-1] = ret[count-1] + tmp/2
			ret[count-2] = ret[count-2] + tmp/2
		}
		count--
		i--
		j--
	}
	for i >= 0 && count >= 0 {
		tmp := arr1[i] + ret[count]
		ret[count] = tmp % 2
		if tmp >= 2 {
			ret[count-1] = ret[count-1] + tmp/2
			ret[count-2] = ret[count-2] + tmp/2
		}
		i--
		count--
	}
	tmp := ret[1]
	ret[1] = ret[1] % 2
	ret[0] = (ret[0] + tmp/2) % 2
	for ret[0] == 0 && len(ret) != 1 {
		if ret[0] == 0 {
			ret = ret[1:]
		}
	}
	return ret
}

/*
21. 移除盒子
给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。

示例：
输入：boxes = [1,3,2,2,2,3,4,3,1]
输出：23
解释：
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)

提示：
1 <= boxes.length <= 100
1 <= boxes[i] <= 100
*/
func removeBoxes(boxes []int) int {
	return 0
}

/*
22. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	rDep := treeNodeDep(root.Right)
	lDep := treeNodeDep(root.Left)
	if rDep-lDep > 1 || lDep-rDep > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func treeNodeDep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dep := treeNodeDep(root.Right)
	lDep := treeNodeDep(root.Left)
	if lDep > dep {
		dep = lDep
	}
	return dep + 1
}

/*
23. 分割字符串的最大得分
给你一个由若干 0 和 1 组成的字符串 s ，请你计算并返回将该字符串分割成两个非空子字符串（即左子字符串和右子字符串）所能获得的最大得分。
「分割字符串的得分」为左子字符串中0 的数量加上右子字符串中1的数量。

示例 1：
输入：s = "011101"
输出：5
解释：
将字符串 s 划分为两个非空子字符串的可行方案有：
左子字符串 = "0" 且 右子字符串 = "11101"，得分 = 1 + 4 = 5
左子字符串 = "01" 且 右子字符串 = "1101"，得分 = 1 + 3 = 4
左子字符串 = "011" 且 右子字符串 = "101"，得分 = 1 + 2 = 3
左子字符串 = "0111" 且 右子字符串 = "01"，得分 = 1 + 1 = 2
左子字符串 = "01110" 且 右子字符串 = "1"，得分 = 2 + 1 = 3
示例 2：
输入：s = "00111"
输出：5
解释：当 左子字符串 = "00" 且 右子字符串 = "111" 时，我们得到最大得分 = 2 + 3 = 5
示例 3：
输入：s = "1111"
输出：3

提示：
2 <= s.length <= 500
字符串 s 仅由字符 '0' 和 '1' 组成。
*/
func maxScore(s string) int {
	ones := 0
	for _, v := range s {
		if v == '1' {
			ones++
		}
	}
	var max, zeros int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		tmp := zeros + ones
		if max < tmp {
			max = tmp
		}
	}
	return max
}

/*
721. 账户合并
给定一个列表accounts，每个元素accounts[i]是一个字符串列表，其中第一个元素accounts[i][0]是名称 (name)，其余元素是emails 表示该帐户的邮箱地址。
现在，我们想合并这些帐户。如果两个帐户都有一些共同的邮件地址，则两个帐户必定属于同一个人。
请注意，即使两个帐户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的帐户，但其所有帐户都具有相同的名称。
合并帐户后，按以下格式返回帐户：每个帐户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。accounts本身可以以任意顺序返回。

例子 1:
Input:
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
Explanation:
  第一个和第三个 John 是同一个人，因为他们有共同的电子邮件 "johnsmith@mail.com"。[["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
  第二个 John 和 Mary 是不同的人，因为他们的电子邮件地址没有被其他帐户使用。
  我们可以以任何顺序返回这些列表，例如答案[['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
  ['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']]仍然会被接受。

注意：
accounts的长度将在[1，1000]的范围内。
accounts[i]的长度将在[1，10]的范围内。
accounts[i][j]的长度将在[1，30]的范围内。
*/
func accountsMerge(accounts [][]string) [][]string {
	emailMap := map[string]int{}
	ret := make([][]string, 0)
	num := 0
	for _, infos := range accounts {
		name := infos[0]
		mails := infos[1:]
		mailList := []string{}
		var seen bool
		var idx int
		var ok bool
		for _, mail := range mails {
			if idx, ok = emailMap[mail]; ok {
				seen = true
			} else {
				mailList = append(mailList, mail)
				emailMap[mail] = num
			}
		}
		if seen {
			ret[idx] = append(ret[idx], mailList...)
		} else {
			tmp := append([]string{name}, mailList...)
			ret = append(ret, tmp)
			num++
		}
	}
	return ret
}

/*
type UnionFind struct {
	count int
	parent []int
}

func (u *UnionFind) Init(count int) {
	u.count = count
	u.parent = make([]int, count)
	for i, _ := range u.parent {
		u.parent[i] = -1
	}
}

func (u *UnionFind) Find(node int) int {
	if u.parent[node] < 0 {
		return node
	}
	result := u.Find(u.parent[node])
	u.parent[node] = result
	return result
}

func (u *UnionFind) Union(node1 int, node2 int) (int, bool) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)
	if root1 == root2 {
		return root2, false
	}
	u.parent[root1] = root2
	u.count--
	return root2, true
}

func accountsMerge(accounts [][]string) [][]string {
	var uf UnionFind
	uf.Init(len(accounts))
	emailToID := make(map[string]int, len(accounts))

	for i, account := range accounts {
		for _, val := range account[1:] {
			findIdx, ok := emailToID[val]
			if !ok {
				emailToID[val] = i
				continue
			}
			uf.Union(i, findIdx)
		}
	}

	tmpWork := make([][]string, len(accounts))
	for k, v := range emailToID {
		p := uf.Find(v)
		if tmpWork[p] == nil {
			tmpWork[p] = []string{k}
			continue
		}
		tmpWork[p] = append(tmpWork[p], k)
	}
	var result [][]string
	for i, r := range tmpWork {
		if r == nil {
			continue
		}
		sort.StringSlice(r).Sort()
		result = append(result, append([]string{accounts[i][0]}, r...))
	}

	return result
}
*/
