package main

import (
	"fmt"
)

func example2() {
	fmt.Println("test")
}

/*
1. 安装栅栏	--	https://leetcode-cn.com/problems/erect-the-fence/	undo!!!
2. 回文子串	--	https://leetcode-cn.com/problems/palindromic-substrings/ undo!!
3. 扫雷游戏	--	https://leetcode-cn.com/problems/minesweeper/
4. 二叉树的最小深度	--	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
5. 重复的子字符串	--	https://leetcode-cn.com/problems/repeated-substring-pattern/
6. 递增子序列	--	https://leetcode-cn.com/problems/increasing-subsequences/
7. 电话号码的字母组合	--	https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
8. 重新安排行程	--	https://leetcode-cn.com/problems/reconstruct-itinerary/
9. 机器人能否返回原点	--	https://leetcode-cn.com/problems/robot-return-to-origin/
*/

/*
1. 安装栅栏
在一个二维的花园中，有一些用 (x, y) 坐标表示的树。由于安装费用十分昂贵，你的任务是先用最短的绳子围起所有的树。
只有当所有的树都被绳子包围时，花园才能围好栅栏。你需要找到正好位于栅栏边界上的树的坐标。

示例 1:
输入: [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
输出: [[1,1],[2,0],[4,2],[3,3],[2,4]]
解释:
见图片erect_the_fence_1.png
示例 2:
输入: [[1,2],[2,2],[4,2]]
输出: [[1,2],[2,2],[4,2]]
解释:
见图片erect_the_fence_2.png
即使树都在一条直线上，你也需要先用绳子包围它们。

注意:
所有的树应当被围在一起。你不能剪断绳子来包围树或者把树分成一组以上。
输入的整数在 0 到 100 之间。
花园至少有一棵树。
所有树的坐标都是不同的。
输入的点没有顺序。输出顺序也没有要求。
*/
func outerTrees(points [][]int) [][]int {
	ret := make([][]int, 0)
	if len(points) < 4 {
		return points
	}
	left_most := 0
	for i := 0; i < len(points); i++ {
		if points[i][0] < points[left_most][0] {
			left_most = i
		}
	}
	p := left_most
	for {
		q := (p + 1) % len(points)
		for i := 0; i < len(points); i++ {
			if orientation(points[p], points[i], points[q]) < 0 {
				q = i
			}
		}
		for i := 0; i < len(points); i++ {
			if i != p && i != q && orientation(points[p], points[i], points[q]) == 0 && inBetween(points[p], points[i], points[q]) {
				ret = append(ret, points[i])
			}
		}
		ret = append(ret, points[q])
		p = q
		if p == left_most {
			break
		}
	}
	return ret
}

func orientation(p, q, r []int) int {
	return (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
}

func inBetween(p, i, q []int) bool {
	a := i[0] >= p[0] && i[0] <= q[0] || i[0] <= p[0] && i[0] >= q[0]
	b := i[1] >= p[1] && i[1] <= q[1] || i[1] <= p[1] && i[1] >= q[1]
	return a && b
}

/*
2. 回文子串
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例 1：
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

提示：
输入的字符串长度不会超过 1000 。
*/
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

/*
3. 扫雷游戏
让我们一起来玩扫雷游戏！
给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，
'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，
数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。
现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。

示例 1：
输入:
[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]
Click : [3,0]
输出:
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]
示例 2：
输入:
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]
Click : [1,2]
输出:
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'X', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

注意：
输入矩阵的宽和高的范围为 [1,50]。
点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
输入面板不会是游戏结束的状态（即有地雷已被挖出）。
简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
*/
var steps [][]int = [][]int{
	[]int{1, 0}, []int{1, -1}, []int{0, -1}, []int{-1, -1},
	[]int{-1, 0}, []int{-1, 1}, []int{0, 1}, []int{1, 1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	stack := [][]int{click}
	seen := map[string]bool{}
	for len(stack) != 0 {
		now := stack[0]
		stack = stack[1:]
		mNum := 0
		candid := [][]int{}
		for _, step := range steps {
			x, y := now[0]+step[0], now[1]+step[1]
			if x >= len(board) || x < 0 ||
				y >= len(board[x]) || y < 0 {
				continue
			}
			if board[x][y] == 'M' {
				mNum++
			} else {
				candid = append(candid, []int{x, y})
			}
		}
		if mNum != 0 {
			board[now[0]][now[1]] = '0' + byte(mNum)
		} else {
			board[now[0]][now[1]] = 'B'
			for _, val := range candid {
				key := fmt.Sprintf("%d-%d", val[0], val[1])
				if _, ok := seen[key]; !ok {
					seen[key] = true
					stack = append(stack, val)
				}
			}
		}

	}
	return board
}

/*
4. 二叉树的最小深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。

示例:
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度2.
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	var ret int
	if root.Right != nil {
		ret = minDepth(root.Right) + 1
	}
	if root.Left != nil {
		tmp := minDepth(root.Left) + 1
		if ret == 0 || tmp < ret {
			ret = tmp
		}
	}
	return ret
}

/*
5. 重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
示例 2:
输入: "aba"
输出: False
示例 3:
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
*/
// KMP 算法为最优解!	--	https://zh.wikipedia.org/wiki/%E5%85%8B%E5%8A%AA%E6%96%AF-%E8%8E%AB%E9%87%8C%E6%96%AF-%E6%99%AE%E6%8B%89%E7%89%B9%E7%AE%97%E6%B3%95
func repeatedSubstringPattern(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		idx := 0
		tmp := true
		for j := i + 1; j < len(s); j++ {
			if s[j] != s[idx] {
				tmp = false
				break
			} else {
				idx++
				if idx%(i+1) == 0 {
					idx = 0
				}
			}
		}
		if tmp && idx == 0 {
			return tmp
		}
	}
	return false
}

/*
6. 递增子序列
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:
给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/
/*
遍历全部子序列，查看是否单调递增
遍历方式：计算出总个数，然后用二进制标识,1表示取，0表示弃
*/
var (
	temp []int
)

func findSubsequences(nums []int) [][]int {
	n = len(nums)
	ans := [][]int{}
	set := map[int]bool{}
	// 2的n-1次方 1<<n
	for i := 0; i < 1<<n; i++ {
		findSubsequences1(i, nums)
		hashValue := getHash(263, int(1e9+7))
		if check() && !set[hashValue] {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
			set[hashValue] = true
		}
	}
	return ans
}

func findSubsequences1(mask int, nums []int) {
	temp = []int{}
	for i := 0; i < n; i++ {
		if (mask & 1) != 0 { // 对整数的每一位和1求与，即能知道整数对应位的数值
			temp = append(temp, nums[i])
		}
		mask >>= 1 // 右移以便分析每一位
	}
}

func getHash(base, mod int) int {
	hashValue := 0
	for _, x := range temp {
		hashValue = hashValue*base%mod + (x + 101)
		hashValue %= mod
	}
	return hashValue
}

func check() bool {
	for i := 1; i < len(temp); i++ {
		if temp[i] < temp[i-1] {
			return false
		}
	}
	return len(temp) >= 2
}

/*
7. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意1不对应任何字母。

示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
var phone map[byte][]rune = map[byte][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

// 更简介的方式，递归(回溯)!!!
func letterCombinations(digits string) []string {
	i := 0
	retByte := [][]rune{[]rune(digits)}
	for i != len(digits) {
		transfer := phone[digits[i]]
		tmp := [][]rune{}
		for _, tpl := range retByte {
			for _, alpha := range transfer {
				a := make([]rune, len(tpl))
				for j, _ := range tpl {
					if j == i {
						a[j] = alpha
					} else {
						a[j] = tpl[j]
					}
				}
				tmp = append(tmp, a)
			}
		}
		retByte = tmp
		i++
	}
	ret := []string{}
	for _, tpl := range retByte {
		if string(tpl) != "" {
			ret = append(ret, string(tpl))
		}
	}
	return ret
}

/*
var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    combinations = []string{}
    backtrack(digits, 0, "")
    return combinations
}

func backtrack(digits string, index int, combination string) {
    if index == len(digits) {
        combinations = append(combinations, combination)
    } else {
        digit := string(digits[index])
        letters := phoneMap[digit]
        lettersCount := len(letters)
        for i := 0; i < lettersCount; i++ {
            backtrack(digits, index + 1, combination + string(letters[i]))
        }
    }
}
*/

/*
8. 重新安排行程
给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，
对该行程进行重新规划排序。所有这些机票都属于一个从JFK（肯尼迪国际机场）出发的先生，所以该行程必须从JFK 开始。

说明:
如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。
示例 1:
输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]
示例 2:
输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
*/
// 超时!
func findItinerary(tickets [][]string) []string {
	ret := findItineraryWithStart(tickets, "JFK")
	return ret
}

func findItineraryWithStart(tickets [][]string, start string) []string {
	ret := []string{start}
	if len(tickets) == 0 {
		return ret
	}
	var minNext []string
	for i, t := range tickets {
		if t[0] == start {
			nTickets := [][]string{}
			for j := 0; j < len(tickets); j++ {
				if j != i {
					nTickets = append(nTickets, tickets[j])
				}
			}
			next := findItineraryWithStart(nTickets, t[1])
			if next != nil && (minNext == nil || minNext[0] > next[0]) {
				minNext = next
			}
		}
	}
	if minNext == nil {
		return nil
	}
	ret = append(ret, minNext...)
	return ret
}

/*
func findItinerary(tickets [][]string) []string {
    var (
        m  = map[string][]string{}
        res []string
    )

    for _, ticket := range tickets {
        src, dst := ticket[0], ticket[1]
        m[src] = append(m[src], dst)
    }
    for key := range m {
        sort.Strings(m[key])
    }

    var dfs func(curr string)
    dfs = func(curr string) {
        for {
            if v, ok := m[curr]; !ok || len(v) == 0 {
                break
            }
            tmp := m[curr][0]
            m[curr] = m[curr][1:]
            dfs(tmp)
        }
        res = append(res, curr)
    }

    dfs("JFK")
    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
    }
    return res
}
*/

/*
9. 机器人能否返回原点
在二维平面上，有一个机器人从原点 (0, 0) 开始。给出它的移动顺序，判断这个机器人在完成移动后是否在 (0, 0) 处结束。
移动顺序由字符串表示。字符 move[i] 表示其第 i 次移动。机器人的有效动作有 R（右），L（左），U（上）和 D（下）。如果机器人在完成所有动作后返回原点，则返回 true。否则，返回 false。
注意：机器人“面朝”的方向无关紧要。 “R” 将始终使机器人向右移动一次，“L” 将始终向左移动等。此外，假设每次移动机器人的移动幅度相同。

示例 1:
输入: "UD"
输出: true
解释：机器人向上移动一次，然后向下移动一次。所有动作都具有相同的幅度，因此它最终回到它开始的原点。因此，我们返回 true。
示例 2:
输入: "LL"
输出: false
解释：机器人向左移动两次。它最终位于原点的左侧，距原点有两次 “移动” 的距离。我们返回 false，因为它在移动结束时没有返回原点。
*/
func judgeCircle(moves string) bool {
	var xStep, yStep int
	for _, act := range moves {
		if act == 'R' {
			xStep++
		} else if act == 'L' {
			xStep--
		} else if act == 'U' {
			yStep++
		} else if act == 'D' {
			yStep--
		}
	}
	if xStep == 0 && yStep == 0 {
		return true
	}
	return false
}
