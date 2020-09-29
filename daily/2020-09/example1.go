package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
	a := '3'
	b := '9'
	fmt.Println(a - '0' + b)
}

/*
1. 预测赢家		--	https://leetcode-cn.com/problems/predict-the-winner/
2. 表示数值的字符串		--	https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
3. 前 K 个高频元素		--	https://leetcode-cn.com/problems/top-k-frequent-elements/
4. 组合总和 II		--	https://leetcode-cn.com/problems/combination-sum-ii/
5. 从中序与后序遍历序列构造二叉树	--	https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
6. 最长回文子串		--	https://leetcode-cn.com/problems/longest-palindromic-substring/
7. 二叉搜索树的最近公共祖先		--	https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
8. 填充每个节点的下一个右侧节点指针 II		--	https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
9. 二叉树的后序遍历		--	https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
*/

/*
1. 预测赢家
给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，
然后玩家 1 拿，…… 。每次一个玩家只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。
给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。

示例 1：
输入：[1, 5, 2]
输出：False
解释：一开始，玩家1可以从1和2中进行选择。
如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2 ）可选。
所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
因此，玩家 1 永远不会成为赢家，返回 False 。
示例 2：
输入：[1, 5, 233, 7]
输出：True
解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
     最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 True，表示玩家 1 可以成为赢家。
提示：
1 <= 给定的数组长度 <= 20.
数组里所有分数都为非负数且不会大于 10000000 。
如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。
*/
func PredictTheWinner(nums []int) bool {
	length := len(nums)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][length-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
2. 表示数值的字符串
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、
"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
*/
/*
确定有限状态自动机
预备知识
确定有限状态自动机（以下简称「自动机」）是一类计算模型。它包含一系列状态，这些状态中：
有一个特殊的状态，被称作「初始状态」。
还有一系列状态被称为「接受状态」，它们组成了一个特殊的集合。其中，一个状态可能既是「初始状态」，也是「接受状态」。
起初，这个自动机处于「初始状态」。随后，它顺序地读取字符串中的每一个字符，并根据当前状态和读入的字符，按照某个事先约定好的「转移规则」，从当前状态转移到下一个状态；当状态转移完成后，它就读取下一个字符。当字符串全部读取完毕后，如果自动机处于某个「接受状态」，则判定该字符串「被接受」；否则，判定该字符串「被拒绝」。

注意：如果输入的过程中某一步转移失败了，即不存在对应的「转移规则」，此时计算将提前中止。在这种情况下我们也判定该字符串「被拒绝」。
一个自动机，总能够回答某种形式的「对于给定的输入字符串 S，判断其是否满足条件 P」的问题。在本题中，条件 P 即为「构成合法的表示数值的字符串」。
自动机驱动的编程，可以被看做一种暴力枚举方法的延伸：它穷尽了在任何一种情况下，对应任何的输入，需要做的事情。
自动机在计算机科学领域有着广泛的应用。在算法领域，它与大名鼎鼎的字符串查找算法「KMP」算法有着密切的关联；在工程领域，它是实现「正则表达式」的基础。

*/
type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: map[CharType]State{
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: map[CharType]State{
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

/*
3. N 皇后
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
上图为 8 皇后问题的一种解法。
给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例：
输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
提示：
皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
*/
func solveNQueens(n int) [][]string {
	return nil
}

/*
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	ret := []string{}
	if root.Left != nil {
		tmp := binaryTreePaths(root.Left)
		for _, val := range tmp {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, val))
		}
	}
	if root.Right != nil {
		tmp := binaryTreePaths(root.Right)
		for _, val := range tmp {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, val))
		}
	}
	return ret
}

/*
3. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:
输入: nums = [1], k = 1
输出: [1]

提示：
你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
*/
func topKFrequent(nums []int, k int) []int {
	freq := make([]map[int]bool, len(nums)+1)
	freqM := map[int]int{}
	maxFreq := 0
	for _, val := range nums {
		if _, ok := freqM[val]; ok {
			freqM[val]++
		} else {
			freqM[val] = 1
		}
		if maxFreq < freqM[val] {
			maxFreq = freqM[val]
		}
		idx := freqM[val] - 1
		if _, ok := freq[idx][val]; ok {
			delete(freq[idx], val)
		}
		if freq[freqM[val]] == nil {
			freq[freqM[val]] = map[int]bool{val: true}
		} else {
			freq[freqM[val]][val] = true
		}
	}
	ret := []int{}
	for len(ret) != k {
		tmp := freq[maxFreq]
		for i, _ := range tmp {
			if len(ret) == k {
				break
			}
			ret = append(ret, i)
		}
		maxFreq--
	}
	return ret
}

// 最小堆
func topKFrequent1(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	/*
		// 最小堆，堆顶为最小值
		for h.Len() > 0 {
			fmt.Printf("%v\n", heap.Pop(h))
		}
	*/
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
4. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	cand := make([][]int, 0)
	ret := make([][]int, 0)
	seen := map[string]bool{}
	for i := 0; i < len(candidates); i++ {
		val := candidates[i]
		if len(cand) != 0 {
			nextCand := [][]int{}
			for _, vals := range cand {
				tmp := []int{}
				for _, ival := range vals {
					tmp = append(tmp, ival)
				}
				tmp = append(tmp, val)
				/*
					}
					for j := 0; j < len(cand); j++ {
					tmp := append(cand[j], val)
				*/
				res := sum(tmp)
				if res < target {
					key := fmt.Sprintf("%v", tmp)
					if _, ok := seen[key]; !ok {
						nextCand = append(nextCand, tmp)
						seen[key] = true
					}
				} else if res == target {
					key := fmt.Sprintf("%v", tmp)
					if _, ok := seen[key]; !ok {
						ret = append(ret, tmp)
						seen[key] = true
					}
				}
			}
			cand = append(cand, nextCand...)
		}
		if val < target {
			key := fmt.Sprintf("%v", []int{val})
			if _, ok := seen[key]; !ok {
				cand = append(cand, []int{val})
				seen[key] = true
			}
		} else if val == target {
			key := fmt.Sprintf("%v", []int{val})
			if _, ok := seen[key]; !ok {
				ret = append(ret, []int{val})
				seen[key] = true
			}
		}
	}
	return ret
}

func sum(in []int) int {
	ret := 0
	for _, val := range in {
		ret += val
	}
	return ret
}

/*
从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	leftInOrder := []int{}
	rightInOrder := []int{}
	rootPass := false
	for _, val := range inorder {
		if val != rootVal {
			if rootPass {
				rightInOrder = append(rightInOrder, val)
			} else {
				leftInOrder = append(leftInOrder, val)
			}
		} else {
			rootPass = true
		}
	}
	leftPostOrder := []int{}
	rightPostOrder := []int{}
	for i := 0; i < len(postorder)-1; i++ {
		if i <= len(leftInOrder)-1 {
			leftPostOrder = append(leftPostOrder, postorder[i])
		} else {
			rightPostOrder = append(rightPostOrder, postorder[i])
		}
	}
	left := buildTree(leftInOrder, leftPostOrder)
	right := buildTree(rightInOrder, rightPostOrder)
	root.Left = left
	root.Right = right
	return root
}

/*
最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"

还有一个Manacher 算法(没理解呢)
*/
// O(n2), O(n2)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	p := make([][]bool, len(s))
	pos := make([]int, 2)
	for i := 0; i < len(s); i++ {
		p[i] = make([]bool, len(s))
		p[i][i] = true
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				pos[0], pos[1] = i, i+1
				p[i][i+1] = true
			}
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			if !p[i+1][j-1] {
				p[i][j] = false
			} else if s[i] == s[j] {
				p[i][j] = true
				if pos[1]-pos[0] < j-i {
					pos[0], pos[1] = i, j
				}
			}
		}
	}
	return string(s[pos[0] : pos[1]+1])
}

// O(n2), O(1)
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
				6
		/				\
		2				8
	/		\		/		\
	0		4		7		9
		/		\
		3		5
示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。


说明:
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val {
		return root
	}
	if root.Val == q.Val {
		return root
	}
	if find(root.Left, p) && find(root.Left, q) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if find(root.Right, p) && find(root.Right, q) {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func find(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == p.Val {
		return true
	}
	return find(root.Right, p) || find(root.Left, p)
}

/*
填充每个节点的下一个右侧节点指针 II
给定一个二叉树
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
进阶：
你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

示例：
			1								1	-> null
		/		\						/		\
		2		3						2	->	3	-> null
	/		\		\				/		\		\
	4		5		7				4	->	5	->	7	->	null
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。

提示：
树中的节点数小于 6000
-100 <= node.val <= 100
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelVals := depLoop(root)
	for _, vals := range levelVals {
		for i := 0; i < len(vals)-1; i++ {
			vals[i].Next = vals[i+1]
		}
	}
	return root
}

func depLoop(root *Node) [][]*Node {
	if root == nil {
		return nil
	}
	ret := [][]*Node{[]*Node{root}}
	left := depLoop(root.Left)
	right := depLoop(root.Right)
	if left != nil {
		for i := 0; i < len(left); i++ {
			tmp := append([]*Node{}, left[i]...)
			if i < len(right) {
				tmp = append(tmp, right[i]...)
			}
			ret = append(ret, tmp)
		}
	}
	if len(left) < len(right) {
		for i := len(left); i < len(right); i++ {
			ret = append(ret, right[i])
		}
	}
	return ret
}

func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	start := root
	var nStart, now *Node
	for start != nil {
		if start.Left != nil {
			if nStart == nil {
				nStart = start.Left
				now = start.Left
			} else {
				now.Next = start.Left
				now = start.Left
			}
		}
		if start.Right != nil {
			if nStart == nil {
				nStart = start.Right
				now = start.Right
			} else {
				now.Next = start.Right
				now = start.Right
			}
		}
		start = start.Next
	}
	connect(nStart)
	return root
}

/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	if root.Left != nil {
		ret = postorderTraversal(root.Left)
	}
	if root.Right != nil {
		tmp := postorderTraversal(root.Right)
		ret = append(ret, tmp...)
	}
	ret = append(ret, root.Val)
	return ret

}

// 迭代算法:
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	orders := []*TreeNode{root}
	ret := []int{}
	for len(orders) != 0 {
		now := orders[len(orders)-1]
		ret = append(ret, now.Val)
		orders = orders[0 : len(orders)-1]
		if now.Left != nil {
			orders = append(orders, now.Left)
		}
		if now.Right != nil {
			orders = append(orders, now.Right)
		}
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return ret
}
