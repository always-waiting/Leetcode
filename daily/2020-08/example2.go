package main

import (
	"fmt"
)

func example2() {
	fmt.Println("test")
}

/*
1. 安装栅栏	--	https://leetcode-cn.com/problems/erect-the-fence/	undo!!!
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
