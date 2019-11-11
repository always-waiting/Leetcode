package stack

/*
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

提示：
1 <= S.length <= 20000
S 仅由小写英文字母组成。
*/
func removeDuplicates(S string) string {
	stack := Stack{}
	for _, val := range S {
		if stack.IsEmpty() {
			stack.Push(val)
			continue
		}
		flag := true
		for !stack.IsEmpty() && stack.Top().Rune() == val {
			flag = false
			stack.Pop()
		}
		if flag {
			stack.Push(val)
		}
	}
	ret := make([]rune, 0)
	for !stack.IsEmpty() {
		ret = append(ret, stack.Top().Rune())
		stack.Pop()
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}

/*
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

示例 1：
输入：S = "ab#c", T = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。

示例 2：
输入：S = "ab##", T = "c#d#"
输出：true
解释：S 和 T 都会变成 “”。

示例 3：
输入：S = "a##c", T = "#a#c"
输出：true
解释：S 和 T 都会变成 “c”。

示例 4：
输入：S = "a#c", T = "b"
输出：false
解释：S 会变成 “c”，但 T 仍然是 “b”。

提示：
1 <= S.length <= 200
1 <= T.length <= 200
S 和 T 只含有小写字母以及字符 '#'。
*/
func backspaceCompare(S string, T string) bool {
	stack1 := Stack{}
	stack2 := Stack{}
	for _, val := range S {
		if string(val) == "#" {
			stack1.Pop()
		} else {
			stack1.Push(val)
		}
	}
	for _, val := range T {
		if string(val) == "#" {
			stack2.Pop()
		} else {
			stack2.Push(val)
		}
	}
	for !stack1.IsEmpty() && !stack2.IsEmpty() {
		if stack1.Pop().Rune() != stack2.Pop().Rune() {
			return false
		}
	}
	if !stack1.IsEmpty() || !stack2.IsEmpty() {
		return false
	}
	return true
}

/*
你现在是棒球比赛记录员。
给定一个字符串列表，每个字符串可以是以下四种类型之一：
1.整数（一轮的得分）：直接表示您在本轮中获得的积分数。
2. "+"（一轮的得分）：表示本轮获得的得分是前两轮有效 回合得分的总和。
3. "D"（一轮的得分）：表示本轮获得的得分是前一轮有效 回合得分的两倍。
4. "C"（一个操作，这不是一个回合的分数）：表示您获得的最后一个有效 回合的分数是无效的，应该被移除。

每一轮的操作都是永久性的，可能会对前一轮和后一轮产生影响。
你需要返回你在所有回合中得分的总和。

示例 1:
输入: ["5","2","C","D","+"]
输出: 30
解释:
第1轮：你可以得到5分。总和是：5。
第2轮：你可以得到2分。总和是：7。
操作1：第2轮的数据无效。总和是：5。
第3轮：你可以得到10分（第2轮的数据已被删除）。总数是：15。
第4轮：你可以得到5 + 10 = 15分。总数是：30。

示例 2:
输入: ["5","-2","4","C","D","9","+","+"]
输出: 27
解释:
第1轮：你可以得到5分。总和是：5。
第2轮：你可以得到-2分。总数是：3。
第3轮：你可以得到4分。总和是：7。
操作1：第3轮的数据无效。总数是：3。
第4轮：你可以得到-4分（第三轮的数据已被删除）。总和是：-1。
第5轮：你可以得到9分。总数是：8。
第6轮：你可以得到-4 + 9 = 5分。总数是13。
第7轮：你可以得到9 + 5 = 14分。总数是27。

注意：
输入列表的大小将介于1和1000之间。
列表中的每个整数都将介于-30000和30000之间。
*/
func calPoints(ops []string) int {
	score := 0
	stack := Stack{}
	for _, action := range ops {
		num := 0
		if action == "C" {
			tmp := stack.Pop()
			if tmp != nil {
				num = tmp.Int()
			}
			score = score - num
		} else if action == "D" {
			tmp := stack.Top()
			if tmp != nil {
				num = tmp.Int()
			}
			score = score + num*2
			stack.Push(num * 2)
		} else if action == "+" {
			tmp1 := stack.Pop()
			tmp2 := stack.Pop()
			if tmp2 != nil {
				num = num + tmp2.Int()
				stack.Push(tmp2.Int())
			}
			if tmp1 != nil {
				num = num + tmp1.Int()
				stack.Push(tmp1.Int())
			}
			score = score + num
			stack.Push(num)
		} else {
			num, _ := strconv.Atoi(action)
			score = score + num
			stack.Push(num)
		}
	}
	return score
}

/*
给定一个二叉树，返回它的中序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := inorderTraversal(root.Left)
	if ret == nil {
		ret = []int{root.Val}
	} else {
		ret = append(ret, root.Val)
	}
	LRet := inorderTraversal(root.Right)
	if LRet != nil {
		ret = append(ret, LRet...)
	}
	return ret
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := Stack{}
	stack.Push(root)
	for !stack.IsEmpty() {
		tmp := stack.Pop()
		node, ok := tmp.val.(*TreeNode)
		if !ok {
			panic("错误")
		}
		if node == nil {
			continue
		}
		right := node.Right
		left := node.Left
		node.Right = nil
		node.Left = nil
		if right != nil {
			stack.Push(right)
		}
		if left != nil {
			stack.Push(node)
			stack.Push(left)
		} else {
			ret = append(ret, node.Val)
		}
	}
	return ret
}
