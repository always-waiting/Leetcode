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
