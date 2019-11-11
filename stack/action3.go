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
		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
	}
	return ret
}
