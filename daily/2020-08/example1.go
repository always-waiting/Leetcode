package main

import "fmt"

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
