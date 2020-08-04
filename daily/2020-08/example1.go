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
