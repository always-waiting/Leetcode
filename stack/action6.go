package stack

import (
	"strings"
)

/*
给定一个整数数组 asteroids，表示在同一行的行星。
对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。

示例 1:
输入:
asteroids = [5, 10, -5]
输出: [5, 10]
解释:
10 和 -5 碰撞后只剩下 10。 5 和 10 永远不会发生碰撞。

示例 2:
输入:
asteroids = [8, -8]
输出: []
解释:
8 和 -8 碰撞后，两者都发生爆炸。

示例 3:
输入:
asteroids = [10, 2, -5]
输出: [10]
解释:
2 和 -5 发生碰撞后剩下 -5。10 和 -5 发生碰撞后剩下 10。

示例 4:
输入:
asteroids = [-2, -1, 1, 2]
输出: [-2, -1, 1, 2]
解释:
-2 和 -1 向左移动，而 1 和 2 向右移动。
由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。
说明:

数组 asteroids 的长度不超过 10000。
每一颗行星的大小都是非零整数，范围是 [-1000, 1000] 。
*/
func asteroidCollision(asteroids []int) []int {
	collStack := make([]int, 0)
	for _, val := range asteroids {
		flag := true
		for len(collStack) != 0 && collStack[len(collStack)-1]*val < 0 && val < 0 {
			tmp1 := collStack[len(collStack)-1]
			tmp2 := -val
			if tmp1 < tmp2 {
				collStack = collStack[:len(collStack)-1]
			} else if tmp1 == tmp2 {
				collStack = collStack[:len(collStack)-1]
				flag = false
				break
			} else {
				flag = false
				break
			}
		}
		if flag {
			collStack = append(collStack, val)
		}
	}
	return collStack
}

/*
给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak < aj。
设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
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
	stack := make([]int, 0)
	for _, val := range nums {
		for len(stack) != 0 && stack[len(stack)-1] > val {
			if len(stack) > 1 {
				return true
			} else {
				stackb := make([]int, 0)
				for i := 0; i < len(stack)-1; i++ {
					stackb = append(stackb, stack[i])
				}
				stack = stackb
			}
		}
		stack = append(stack, val)
	}
	return false
}

/*
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
示例 1:
输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true

示例 2:
输入: "1,#"
输出: false

示例 3:
输入: "9,#,#,1"
输出: false
*/
func isValidSerialization(preorder string) bool {
	leaves := 0
	node := 0
	pres := strings.Split(preorder, ",")
	for i, s := range pres {
		if s == "#" {
			leaves++
		} else {
			node++
		}
		if leaves > node+1 {
			return false

		}
		if leaves == node+1 && i < len(pres)-1 {
			return false
		}

	}
	if leaves != node+1 {
		return false
	} else {
		return true
	}
}
