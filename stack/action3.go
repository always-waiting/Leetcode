package stack

import "strings"

/*
以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。

示例 1：
输入："/home/"
输出："/home"
解释：注意，最后一个目录名后面没有斜杠。

示例 2：
输入："/../"
输出："/"
解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。

示例 3：
输入："/home//foo/"
输出："/home/foo"
解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。

示例 4：
输入："/a/./b/../../c/"
输出："/c"

示例 5：
输入："/a/../../b/../c//.//"
输出："/c"

示例 6：
输入："/a//b////c/d//././/.."
输出："/a/b/c"
*/
func simplifyPath(path string) string {
	stack := Stack{}
	//stack.Push("/")
	paths := strings.Split(path, "/")
	for _, val := range paths {
		if val == ".." {
			stack.Pop()
		} else if val == "." {
			continue
		} else if val == "" {
		} else {
			stack.Push(val)
		}
	}
	ret := make([]string, 0)
	for !stack.IsEmpty() {
		ret = append(ret, stack.Pop().String())
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return "/" + strings.Join(ret, "/")
}

/*
给定一个二叉树，返回它的 前序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	if root.Left != nil {
		left := preorderTraversal(root.Left)
		ret = append(ret, left...)
	}
	if root.Right != nil {
		right := preorderTraversal(root.Right)
		ret = append(ret, right...)
	}
	return ret
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := Stack{}
	stack.Push(root)
	ret := make([]int, 0)
	for !stack.IsEmpty() {
		tmp := stack.Pop()
		node, ok := tmp.val.(*TreeNode)
		if !ok {
			panic("错误")
		}
		ret = append(ret, node.Val)
		if node.Left != nil && node.Right != nil {
			stack.Push(node.Right)
			stack.Push(node.Left)
		} else if node.Left != nil {
			stack.Push(node.Left)
		} else if node.Right != nil {
			stack.Push(node.Right)
		}
	}
	return ret
}

/*
根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/
func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	stackT := Stack{}
	stackI := Stack{}
	for i := len(T) - 1; i >= 0; i-- {
		for !stackT.IsEmpty() && stackT.Top().Int() <= T[i] {
			stackT.Pop()
			stackI.Pop()
		}
		if !stackT.IsEmpty() {
			ret[i] = stackI.Top().Int() - i
		}
		stackT.Push(T[i])
		stackI.Push(i)
	}
	return ret
}

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:
s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/
func decodeString(s string) string {
	stackNum := Stack{}
	stackStr := Stack{}
	var repeat string
	var ret string
	var level int
	var start bool
	var numstr string
	for _, val := range s {
		str := string(val)
		_, err := strconv.Atoi(str)
		if err == nil {
			numstr = numstr + str
			continue
		} else {
			if str == "[" {
				start = true
				if level > 0 {
					stackStr.Push(repeat)
				}
				level++
				repeat = ""
				num, _ := strconv.Atoi(numstr)
				numstr = ""
				stackNum.Push(num)
				continue
			} else if str == "]" {
				repeatNum := stackNum.Pop().Int()
				var tmp string
				for i := 0; i < repeatNum; i++ {
					tmp = tmp + repeat
				}
				level--
				if level > 0 {
					repeat = stackStr.Pop().String() + tmp
				} else {
					start = false
					ret = ret + tmp
				}
			} else {
				if start {
					repeat = repeat + str
				} else {
					ret = ret + str
				}
			}

		}
	}
	return ret
}

func decodeString1(s string) string {
	stackS := Stack{}
	stackI := Stack{}
	var multi int
	var repeat string
	for _, val := range s {
		if isDigit(val) {
			num, _ := strconv.Atoi(string(val))
			multi = multi*10 + num
		} else {
			if string(val) == "[" {
				stackI.Push(multi)
				stackS.Push(repeat)
				multi = 0
				repeat = ""
			} else if string(val) == "]" {
				num := stackI.Pop().Int()
				last_str := stackS.Pop().String()
				cur_repeat := ""
				for i := 0; i < num; i++ {
					cur_repeat = cur_repeat + repeat
				}
				repeat = last_str + cur_repeat
			} else {
				repeat = repeat + string(val)
			}
		}
	}
	return repeat
}

func isDigit(r rune) bool {
	arr := []rune("09")
	if r >= arr[0] && r <= arr[1] {
		return true
	}
	return false
}
